package alliance

import (
	"github.com/wwj31/dogactor/actor"

	"server/common"
	"server/common/actortype"
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/rdsop"
)

func (a *Alliance) manifestMaintenance() {
	if len(a.manifests) == 0 {
		return
	}

	// 先统计各个清单空房间数量
	emptyRoomStat := make(map[string]map[actor.Id]struct{}) // map[清单id][房间actorId]
	for id, _ := range a.manifests {
		emptyRoomStat[id] = make(map[actor.Id]struct{})
	}
	roomList := rdsop.RoomList(a.allianceId)

	for _, roomId := range roomList {
		roomActorId := actortype.RoomName(roomId)
		v, err := a.RequestWait(roomActorId, &inner.RoomIsEmptyReq{})
		if err != nil {
			log.Errorw("RoomIsEmptyReq error", "err", err)
			continue
		}

		rsp := v.(*inner.RoomIsEmptyRsp)
		if rsp.Empty {
			if emptyRoomStat[rsp.ManifestId] == nil {
				log.Errorw("room cannot find manifest", "id", rsp.ManifestId)
				continue
			}
			emptyRoomStat[rsp.ManifestId][roomActorId] = struct{}{}
		}
	}
	log.Infow("regular maintain", "emptyRoomStat", emptyRoomStat)

	// 检查是否满足清单所需数量，不满足的执行补充，多于所需数量尝试销毁
	for id, manifest := range a.manifests {
		// 补充房间
		for i := len(emptyRoomStat[id]); i < int(manifest.GameParams.MaintainEmptyRoom); i++ {
			roomMgrId := rdsop.GetRoomMgrId()
			if roomMgrId == -1 {
				log.Errorw("load rooms failed", "roomMgrId", roomMgrId)
				return
			}

			v, err := a.RequestWait(actortype.RoomMgrName(roomMgrId), &inner.CreateRoomReq{
				RoomId:         0,
				GameType:       manifest.GameType,
				CreatorShortId: a.master.ShortId,
				AllianceId:     a.AllianceId(),
				GameParams:     common.ProtoMarshal(manifest.GameParams),
				ManifestId:     id,
			})
			if yes, _ := common.IsErr(v, err); yes {
				log.Errorw("create room failed", "err", err, "v", v)
				continue
			}
			log.Infow("create alliance room success", "alliance", a.allianceId, "manifestId", id, "manifest", manifest.GameParams.String())
		}

		// 尝试销毁多余的空房间
		for i := int(manifest.GameParams.MaintainEmptyRoom); i < len(emptyRoomStat[id]); i++ {
			// 随便选个空房间，尝试销毁
			var roomActorId actor.Id
			for v := range emptyRoomStat[id] {
				roomActorId = v
				break
			}

			v, err := a.RequestWait(roomActorId, &inner.DisbandRoomReq{})
			if yes, _ := common.IsErr(v, err); yes {
				log.Warnw("disband room failed", "err", err, "v", v)
				continue
			}
			log.Infow("disband alliance room success", "alliance", a.allianceId, "room", roomActorId, "manifestId", id, "manifest", manifest.GameParams.String())
		}
	}

}
