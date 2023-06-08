package login

import (
	"context"
	"fmt"
	"time"

	"server/rdsop"

	"github.com/golang-jwt/jwt/v4"

	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"server/common"
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

func (s *Login) Login(gSession common.GSession, req *outer.LoginReq) {
	go tools.Try(func() {
		rds.LockDo(rdsop.LockLoginKey(req.DeviceID), func() {
			var (
				acc       *account.Account
				newPlayer bool
				err       error
				errCode   outer.ERROR
			)

			defer func() {
				if err != nil {
					gSession.SendToClient(s, &outer.FailRsp{Error: errCode, Info: err.Error()})
					return
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

			case PhoneLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"phone": req.Phone})
				if err = result.Err(); err == mongo.ErrNoDocuments {
					errCode = outer.ERROR_PHONE_NOT_FOUND
					return
				}

			case TokenLogin:
				// 解析并验证 JWT
				var claims *Claims
				claims, err = common.JWTParseToken(req.Token, &Claims{})
				if err != nil {
					log.Warnw("token login failed ", "err", err, "req", req.String())
					errCode = outer.ERROR_LOGIN_TOKEN_INVALID
					return
				}

				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"_id": claims.UID})
				if err = result.Err(); err == mongo.ErrNoDocuments {
					log.Warnw("token login can not find account", "err", err, "req", req.String())
					errCode = outer.ERROR_LOGIN_TOKEN_INVALID
					return
				}
			case WeiXinLogin:
				if req.WeiXinOpenID == "" {
					err = fmt.Errorf("weixin login failed, openID is nil")
					return
				}
				return

				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"wei_xin_open_id": req.WeiXinOpenID})
				if result.Err() == mongo.ErrNoDocuments {
					acc.WeiXinOpenID = req.WeiXinOpenID
				}
			}

			if result.Err() == mongo.ErrNoDocuments {
				//err = s.initAccount(acc, req.OS, req.ClientVersion, req.UpShortId)
				err = s.initAccount(acc, req.OS, req.ClientVersion, 0)
				if err != nil {
					errCode = outer.ERROR_NEW_ACCOUNT_FAILED
					return
				}
				newPlayer = true
			} else {
				if err = result.Err(); err != nil {
					log.Errorw("login mongo find failed", "err", result.Err())
					return
				}

				if err = result.Decode(acc); err != nil {
					log.Errorw("login find account decode failed", "err", err)
					return
				}

				// 手机账号登录，需要单独校验密码是否正确
				if req.LoginType == PhoneLogin && req.PhonePassword != acc.PhonePassword {
					log.Warnw("password error", "pwd", acc.PhonePassword)
					errCode = outer.ERROR_PHONE_PASSWORD_ERROR
					return
				}
			}

			// 获得最后一次登录的gSession,踢掉旧链接
			val := rds.Ins.Get(context.Background(), rdsop.SessionKey(acc.LastLoginRID)).Val()
			oldGateSession := common.GSession(val)
			if oldGateSession.Valid() {
				gate, _ := oldGateSession.Split()
				s.RequestWait(gate, &inner.KickOutReq{
					GateSession: oldGateSession.String(),
					RID:         acc.LastLoginRID,
				}, 3*time.Second)
			}
			rds.Ins.Set(context.Background(), rdsop.SessionKey(acc.LastLoginRID), gSession.String(), 7*24*time.Hour)

			// 找玩家最近登录过的game节点，如果没找到就重新分配节点
			var dispatchGameId string
			gameNodeId, _ := rds.Ins.Get(context.Background(), rdsop.GameNodeKey(acc.LastShortID)).Result()
			if gameNodeId != "" {
				dispatchGameId = gameNodeId
			} else {
				dispatchGameId = s.getGameNode()
			}

			_, err = s.RequestWait(dispatchGameId, &inner.PullPlayer{
				RID: acc.LastLoginRID,
				NewData: &inner.NewPlayerInfo{
					AccountInfo: acc.ToPb(),
					ShortId:     acc.Roles[acc.LastLoginRID].ShorID,
				},
			})

			rds.Ins.Set(context.Background(), rdsop.GameNodeKey(acc.LastShortID), dispatchGameId, 7*24*time.Hour)
			log.Infow("login success dispatch the player to game",
				"new", newPlayer, "role", acc.Roles[acc.LastLoginRID], "req", req.String(), "to game", dispatchGameId)

			if err != nil {
				log.Errorw("send to game failed ", "err", err, "game", dispatchGameId)
				return
			}
		})
	})
}

func (s *Login) responseLoginToClient(acc *account.Account, newPlayer bool, gSession common.GSession) {
	// 走到这里，说明已经登录成功，
	// 通知gateway，绑定关联Session的用户信息
	gateId, _ := gSession.Split()
	_, err := s.RequestWait(gateId, &inner.BindSessionWithRID{
		GateSession: gSession.String(),
		RID:         acc.LastLoginRID,
	})
	if err != nil {
		log.Errorw("bind session with rid failed",
			"gSession", gSession.String(),
			"rid", acc.LastLoginRID,
			"err", err)
		return
	}

	// 重新给前端创建新的Token
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)),
		},
		UID: acc.UUID,
		RID: acc.LastLoginRID,
	}

	signedToken, signErr := common.JWTSignedToken(claims)
	if signErr != nil {
		log.Errorw("jwt signed token failed",
			"rid", acc.LastLoginRID,
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
