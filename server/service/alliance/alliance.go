package alliance

import (
	"context"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/event"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/tools"
	"go.mongodb.org/mongo-driver/bson"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/common/mongodb"
	"server/common/rds"
	"server/common/router"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
	"server/rdsop"
)

type RID = string

func New(id int32) *Alliance {
	return &Alliance{
		allianceId:       id,
		members:          make(map[RID]*Member),
		membersByShortId: make(map[int64]*Member),
		manifests:        make(map[string]*Manifest),
	}
}

type (
	Alliance struct {
		actor.Base
		roomLoad         bool
		allianceId       int32
		disband          bool
		members          map[RID]*Member
		membersByShortId map[int64]*Member
		master           *Member

		// 自动创建房间清单
		manifests map[string]*Manifest

		currentMsg      actor.Message
		currentGSession common.GSession
	}
)

func (a *Alliance) MemberColl() string {
	return fmt.Sprintf("alliance_%v", a.allianceId)
}

func (a *Alliance) ManifestColl() string {
	return fmt.Sprintf("alliance_manifest_%v", a.allianceId)
}

func (a *Alliance) OnInit() {
	a.loadAllMember()
	a.loadAllManifest()

	// 统一返回结果
	router.Result(a, a.responseHandle)

	a.System().OnEvent(a.ID(), func(ev event.EvNewActor) {
		if actortype.IsActorOf(ev.ActorId, actortype.RoomMgrActor) {
			a.loadRooms()
		}
	})

	log.Debugf("Alliance OnInit %v members:%v", a.ID(), len(a.members))
}

func (a *Alliance) loadAllMember() {
	cur, err := mongodb.Ins.Collection(a.MemberColl()).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance member failed", "err", err)
		return
	}

	var members []*Member
	err = cur.All(context.Background(), &members)
	if err != nil {
		log.Errorw("decode all member failed", "err", err)
		return
	}

	for _, member := range members {
		member.Alliance = a
		a.members[member.RID] = member
		a.membersByShortId[member.ShortId] = member
		if member.Position == Master {
			a.master = member
		}
		log.Debugf("alliance:%v load member %+v", a.allianceId, *member)
	}
}

func (a *Alliance) loadAllManifest() {
	cur, err := mongodb.Ins.Collection(a.ManifestColl()).Find(context.Background(), bson.M{})
	if err != nil {
		log.Errorw("load all alliance manifest failed", "err", err)
		return
	}

	var manifests []*Manifest
	err = cur.All(context.Background(), &manifests)
	if err != nil {
		log.Errorw("decode all manifest failed", "err", err)
		return
	}

	for _, manifest := range manifests {
		a.manifests[manifest.UID] = manifest
		log.Debugf("alliance:%v load manifest %+v", a.allianceId, *manifest)
	}
}

// 所有消息,处理完统一返回流程
func (a *Alliance) responseHandle(resultMsg any) {
	msg, ok := resultMsg.(proto.Message)
	if !ok {
		return
	}

	// 网关消息，直接将消息转发给session, 其他服务消息，走内部通讯接口
	if actortype.IsActorOf(a.currentMsg.GetSourceId(), actortype.GatewayActor) {
		a.Send2Client(a.currentGSession, msg)
	} else {
		var err error
		if a.currentMsg.GetRequestId() != "" {
			err = a.Response(a.currentMsg.GetRequestId(), msg)
		} else {
			err = a.Send(a.currentMsg.GetSourceId(), msg)
		}

		if err != nil {
			log.Warnw("response to actor failed",
				"source", a.currentMsg.GetSourceId(), "msg name", a.currentMsg.GetMsgName())
		}
	}
}

func (a *Alliance) loadRooms() {
	if a.roomLoad {
		return
	}
	a.roomLoad = true
	roomList := rdsop.RoomList(a.allianceId)
	log.Debugw("load rooms", "room list", roomList)
	for _, roomId := range roomList {
		roomInfo := rdsop.NewRoomInfo{RoomId: roomId}.GetInfoFromRedis()
		gameParamsBytes, _ := proto.Marshal(roomInfo.Params)
		roomMgrId := rdsop.GetRoomMgrId()
		if roomMgrId == -1 {
			log.Errorw("load rooms failed", "roomMgrId", roomMgrId)
			return
		}

		v, err := a.RequestWait(actortype.RoomMgrName(roomMgrId), &inner.CreateRoomReq{
			RoomId:         roomId,
			GameType:       roomInfo.GameType,
			CreatorShortId: roomInfo.CreatorShortId,
			AllianceId:     roomInfo.AllianceId,
			GameParams:     gameParamsBytes,
		})
		if yes, _ := common.IsErr(v, err); yes {
			log.Errorw("load request create room failed", "err", err, "v", v)
			continue
		}
		log.Infow("create alliance room success", "alliance", a.allianceId, "room", roomInfo)
	}
}

func (a *Alliance) Send2Client(gSession common.GSession, msg proto.Message) {
	log.Infow("output", "alliance", a.ID(), "msg", reflect.TypeOf(msg), "data", msg.String())
	gSession.SendToClient(a, msg)
}

func (a *Alliance) OnStop() bool {
	log.Infof("stop Alliance %v", a.ID())
	if a.disband {
		err := mongodb.Ins.Collection(a.MemberColl()).Drop(context.Background())
		if err != nil {
			log.Warnw("disband alliance drop table failed", "coll", a.MemberColl())
		}
		rds.Ins.Del(context.Background(), rdsop.JoinAllianceKey(a.master.ShortId))
	}
	return true
}

func (a *Alliance) OnHandle(msg actor.Message) {
	a.currentMsg = msg

	message, _, gSession, err := common.UnwrappedGateMsg(msg.Payload())
	expect.Nil(err)
	a.currentGSession = gSession

	pt, ok := message.(proto.Message)
	if !ok {
		log.Warnw("unknown msg", "msg", reflect.TypeOf(message).String())
		return
	}

	log.Infow("input", "alliance", a.ID(), "source", msg.GetSourceId(), "msg", reflect.TypeOf(pt), "data", pt.String())
	if routerErr := router.Dispatch(a, pt); routerErr != nil {
		log.Warnw("alliance dispatch the message failed", "err", routerErr)
	}
}

func (a *Alliance) MemberInfo(rid RID) *Member                { return a.members[rid] }
func (a *Alliance) MemberInfoByShortId(shortId int64) *Member { return a.membersByShortId[shortId] }
func (a *Alliance) Master() *Member                           { return a.master }
func (a *Alliance) AllianceId() int32                         { return a.allianceId }
func (a *Alliance) Disband()                                  { a.disband = true }

func (a *Alliance) KickOutMember(shortId int64) {
	member := a.membersByShortId[shortId]
	if member != nil {
		delete(a.members, member.RID)
		delete(a.membersByShortId, member.ShortId)
	}
}

func (a *Alliance) Members() (arr []*Member) {
	for _, member := range a.members {
		arr = append(arr, member)
	}
	return
}

func (a *Alliance) ManifestListPB(typ outer.GameType) (list []*outer.Manifest) {
	for _, manifest := range a.manifests {
		if manifest.GameType == typ.Int32() {
			list = append(list, manifest.ToPB())
		}
	}
	return
}
func (a *Alliance) SetManifest(uid string, typ outer.GameType, params *outer.GameParams) *Manifest {
	obj := &Manifest{
		Alliance:   a,
		UID:        uid,
		GameType:   typ.Int32(),
		GameParams: params,
	}

	if uid != "" {
		manifest, ok := a.manifests[uid]
		if ok {
			manifest.GameParams = params
			manifest.Save()
			return manifest
		}
	} else {
		for _, manifest := range a.manifests {
			if manifest.ArgEqual(obj) {
				manifest.GameParams.MaintainEmptyRoom = params.MaintainEmptyRoom
				manifest.Save()
				return manifest
			}
		}
		obj.UID = tools.UUID()
	}

	a.manifests[obj.UID] = obj
	obj.Save()
	return obj
}

func (a *Alliance) PlayerOnline(gSession common.GSession, rid RID) {

	log.Infow("player online ", "gSession", gSession, "rid", rid)
}

func (a *Alliance) PlayerOffline(gSession common.GSession, rid RID) {

	log.Infow("player offline ", "gSession", gSession, "rid", rid)
}
