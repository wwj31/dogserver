package login

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"server/common/rdskey"
	"time"

	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	rds "server/common/redis"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/login/account"
)

var secretKey = []byte("fuck you!!!!")

type Claims struct {
	jwt.RegisteredClaims
	UID string
	RID string
}

const (
	GuestLogin  = 1
	PhoneLogin  = 2
	TokenLogin  = 3
	WeiXinLogin = 4
)

func (s *Login) Login(gSession common.GSession, req *outer.LoginReq) {
	go tools.Try(func() {
		rds.LockDo(rdskey.LoginKey(req.DeviceID), func() {
			var (
				acc       *account.Account
				newPlayer bool
				err       error
			)

			defer func() {
				if acc == nil {
					gSession.SendToClient(s, &outer.Fail{
						Error: outer.ERROR_FAILED,
						Info:  err.Error(),
					})
				}

				gateId, _ := gSession.Split()
				_, err = s.RequestWait(gateId, &inner.BindSessionWithRID{
					GateSession: gSession.String(),
					RID:         acc.LastLoginRID,
				})
				if err != nil {
					log.Errorw("bind session with rid failed",
						"gsession", gSession.String(),
						"RID", acc.LastLoginRID,
						"err", err)
					return
				}

				// 创建 JWT 声明
				claims := &Claims{
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * 24 * time.Hour)),
					},
					UID: acc.UUID,
					RID: acc.LastLoginRID,
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				signedToken, err := token.SignedString(secretKey)
				if err != nil {
					log.Errorw("jwt signedString failed",
						"gsession", gSession.String(),
						"RID", acc.LastLoginRID,
						"err", err)
					return
				}

				md5 := common.EnterGameToken(acc.LastLoginRID, newPlayer)
				gSession.SendToClient(s, &outer.LoginRsp{
					RID:       acc.LastLoginRID,
					NewPlayer: newPlayer,
					Token:     signedToken,
					Checksum:  md5,
				})
				log.Debugw("player login success", "RID", acc.LastLoginRID, "UID", acc.UUID)
			}()

			var result *mongo.SingleResult
			switch req.LoginType {
			case GuestLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"device_id": req.DeviceID})
			case TokenLogin:
				// 解析并验证 JWT
				parsedToken, err := jwt.ParseWithClaims(req.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
					return secretKey, nil
				})
				if err != nil {
					log.Errorw("token login failed ", "err", err, "req", req.String())
					return
				}
				claims, ok := parsedToken.Claims.(*Claims)
				if !ok {
					log.Warnw("asset token failed ", "parsedToken", parsedToken, "req", req.String())
					return
				}
				if !parsedToken.Valid {
					log.Warnw("invalid token ", "expire_at", claims.ExpiresAt.Time, "req", req.String())
					return
				}
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"_id": claims.UID})
			case WeiXinLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"wei_xin_open_id": req.WeiXinOpenID})
			case PhoneLogin:
				result = mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"phone": req.Phone})
				// TODO smsCode存redis里，这里需要校验code是否有效
			}

			if result.Err() == mongo.ErrNoDocuments {
				acc = account.New()
				acc.UUID = tools.XUID()
				acc.WeiXinOpenID = req.WeiXinOpenID
				acc.DeviceID = req.DeviceID
				acc.Phone = req.Phone
				acc.SID = actortype.GameName(1)
				acc.Roles = make(map[string]account.Role)
				rid := tools.XUID()
				acc.Roles[rid] = account.Role{RID: rid}
				acc.LastLoginRID = rid
				if _, err = mongodb.Ins.Collection(account.Collection).InsertOne(context.Background(), acc); err != nil {
					log.Errorw("login insert new account failed ", "UUID", acc.UUID, "err", err)
				}
				newPlayer = true
			} else {
				if result.Err() != nil {
					log.Errorw("login mongo find failed", "err", result.Err())
					return
				}

				acc = &account.Account{}
				if err = result.Decode(acc); err != nil {
					log.Errorw("login find account decode failed", "err", err)
					return
				}

			}
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

			_, err = s.RequestWait(acc.SID, &inner.PullPlayer{
				RID: acc.LastLoginRID,
			})

			if err != nil {
				log.Errorw("send to game failed ", "err", err, "sid", acc.SID)
				return
			}
		})
	})
}
