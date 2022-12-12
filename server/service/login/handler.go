package login

import (
	"context"
	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/redis"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/service/login/account"
)

func redisKey(key string) string {
	return "lock:login:{" + key + "}"
}

func (s *Login) Login(gSession common.GSession, msg *outer.LoginReq) {
	go tools.Try(func() {
		redis.LockDo(redisKey(msg.PlatformUUID), func() {
			var (
				acc       *account.Account
				newPlayer bool
				err       error
			)

			defer func() {
				if acc == nil {
					_ = s.Send2Client(gSession, &outer.Fail{
						Error: outer.ERROR_FAILED,
						Info:  err.Error(),
					})
				}

				md5 := common.LoginMD5(acc.UUID, acc.LastLoginRID, newPlayer)
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

				_ = s.Send2Client(gSession, &outer.LoginResp{
					UID:       acc.UUID,
					RID:       acc.LastLoginRID,
					NewPlayer: newPlayer,
					Token:     md5,
				})
				log.Debugw("player login success", "RID", acc.LastLoginRID, "UID", acc.UUID)
			}()

			uuid := msg.PlatformName + ":" + msg.PlatformUUID
			result := mongodb.Ins.Collection(account.Collection).FindOne(context.Background(), bson.M{"_id": uuid})
			if result.Err() == mongo.ErrNoDocuments {
				acc = account.New()
				acc.PlatformID = msg.PlatformUUID
				acc.PlatformName = msg.PlatformName
				acc.UUID = uuid
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
