package login

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/cast"

	"github.com/golang-jwt/jwt/v4"

	"server/common/rdskey"

	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/rds"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/login/account"
)

type Claims struct {
	jwt.RegisteredClaims
	UID string
	RID string
}

const (
	GuestLogin  = 1
	PhoneLogin  = 2
	WeiXinLogin = 3
	TokenLogin  = 4
)

// 登录规则
// 1.游客登录使用DeviceID，游客账号可以绑定未使用过的微信和电话
// 2.微信登录创建的角色，不能在绑定DeviceID，可以绑定为使用过的电话
// 3.电话登录创建的角色，不能在绑定DeviceID，可以绑定为使用过的微信

func (s *Login) Login(gSession common.GSession, req *outer.LoginReq) {
	go tools.Try(func() {
		rds.LockDo(rdskey.LockLoginKey(req.DeviceID), func() {
			var (
				acc        *account.Account
				newPlayer  bool
				newShortID int64
				err        error
				errCode    outer.ERROR
			)

			defer func() {
				if acc == nil {
					gSession.SendToClient(s, &outer.FailRsp{
						Error: errCode,
						Info:  err.Error(),
					})
				}
				s.responseLoginToClient(acc, newPlayer, gSession)
			}()

			var result *mongo.SingleResult
			acc = account.New()
			switch req.LoginType {
			case GuestLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"device_id": req.DeviceID})
				if result.Err() == mongo.ErrNoDocuments {
					acc.DeviceID = req.DeviceID
				}
			case TokenLogin:
				// 解析并验证 JWT
				var claims *Claims
				claims, err = common.JWTParseToken(req.Token, &Claims{})
				if err != nil {
					log.Errorw("token login failed ", "err", err, "req", req.String())
					errCode = outer.ERROR_LOGIN_TOKEN_INVALID
					return
				}

				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"_id": claims.UID})
				if result.Err() == mongo.ErrNoDocuments {
					log.Errorw("token login can not find account", "err", err, "req", req.String())
					errCode = outer.ERROR_LOGIN_TOKEN_INVALID
					return
				}
			case WeiXinLogin:
				if req.WeiXinOpenID == "" {
					err = fmt.Errorf("weixin login failed, openID is nil")
					return
				}
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"wei_xin_open_id": req.WeiXinOpenID})
				if result.Err() == mongo.ErrNoDocuments {
					acc.WeiXinOpenID = req.WeiXinOpenID
				}
			case PhoneLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"phone": req.Phone})
				if result.Err() == mongo.ErrNoDocuments {
					err = fmt.Errorf("未找到账号:%v", req.Phone)
					return
				}
			}

			dispatchGameID := actortype.GameName(1)
			if result.Err() == mongo.ErrNoDocuments {
				var shortIdVal interface{}
				shortIdVal, err = rds.Ins.EvalSha(context.Background(), s.sha1, []string{rdskey.ShortIDKey()}).Result()
				if err == redis.Nil {
					err = fmt.Errorf("short id pool was empty")
					log.Errorw(err.Error())
					return
				}

				if err != nil {
					err = fmt.Errorf("shor id get failed :%v", err)
					log.Errorw(err.Error())
					return
				}

				arr, ok := shortIdVal.([]interface{})
				if !ok || len(arr) != 1 {
					err = fmt.Errorf("shortIdVal  failed:%v len:%v", shortIdVal, len(arr))
					log.Errorw(err.Error())
					return
				}

				acc.UUID = tools.XUID()
				acc.Roles = make(map[string]account.Role)
				rid := tools.XUID()
				newShortID = cast.ToInt64(arr[0])
				acc.Roles[rid] = account.Role{RID: rid, ShorID: newShortID, CreateAt: time.Now()}
				acc.LastShortID = acc.Roles[rid].ShorID
				acc.LastLoginRID = rid
				log.Infof("acc device %v", acc.DeviceID)
				if _, err = mongodb.Ins.Collection(account.Collection).InsertOne(context.Background(), acc); err != nil {
					log.Errorw("login insert new account failed ", "UUID", acc.UUID, "err", err)
				}
				newPlayer = true
			} else {
				if result.Err() != nil {
					log.Errorw("login mongo find failed", "err", result.Err())
					return
				}

				if err = result.Decode(acc); err != nil {
					log.Errorw("login find account decode failed", "err", err)
					return
				}
			}

			// 获得最后一次登录的gSession,踢掉旧链接
			val := rds.Ins.Get(context.Background(), rdskey.SessionKey(acc.LastLoginRID)).Val()
			oldGateSession := common.GSession(val)
			if oldGateSession.Valid() {
				gate, _ := oldGateSession.Split()
				s.RequestWait(gate, &inner.KickOutReq{
					GateSession: oldGateSession.String(),
					RID:         acc.LastLoginRID,
				}, 3*time.Second)
			}
			rds.Ins.Set(context.Background(), rdskey.SessionKey(acc.LastLoginRID), gSession.String(), 3*24*time.Hour)

			_, err = s.RequestWait(dispatchGameID, &inner.PullPlayer{
				Account: acc.ToPb(),
				RoleInfo: &inner.LoginRoleInfo{
					RID:     acc.LastLoginRID,
					ShortID: acc.Roles[acc.LastLoginRID].ShorID,
				},
			})

			log.Infow("login success dispatch the player to game",
				"new", newPlayer, "role", acc.Roles[acc.LastLoginRID], "req", req.String(), "to game", dispatchGameID)

			if err != nil {
				log.Errorw("send to game failed ", "err", err, "game", dispatchGameID)
				return
			}
		})
	})
}

func (s *Login) responseLoginToClient(acc *account.Account, newPlayer bool, gSession common.GSession) {
	gateId, _ := gSession.Split()
	_, err := s.RequestWait(gateId, &inner.BindSessionWithRID{
		GateSession: gSession.String(),
		RID:         acc.LastLoginRID,
	})
	if err != nil {
		log.Errorw("bind session with rid failed",
			"gSession", gSession.String(),
			"RID", acc.LastLoginRID,
			"err", err)
		return
	}

	// 创建 JWT 声明
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
		UID: acc.UUID,
		RID: acc.LastLoginRID,
	}

	signedToken, signErr := common.JWTSignedToken(claims)
	if signErr != nil {
		log.Errorw("jwt signed token failed",
			"RID", acc.LastLoginRID,
			"err", err)
	}

	md5 := common.EnterGameToken(acc.LastLoginRID, newPlayer)
	gSession.SendToClient(s, &outer.LoginRsp{
		RID:       acc.LastLoginRID,
		NewPlayer: newPlayer,
		Token:     signedToken,
		Checksum:  md5,
	})
}
