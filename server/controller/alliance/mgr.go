package alliance

import (
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/service/alliance"

	"server/common/router"
)

// 创建联盟
var _ = router.Reg(func(mgr *alliance.Mgr, msg *inner.CreateAllianceReq) any {
	allianceId, err := mgr.CreateAlliance(msg.MasterShortId)
	if err != nil {
		log.Errorw("create alliance failed", "master", msg.MasterShortId, "err", err)
		return &inner.Error{}
	}

	return &inner.CreateAllianceRsp{AllianceId: allianceId}
})

// 联盟被解散
var _ = router.Reg(func(mgr *alliance.Mgr, msg *inner.DisbandAllianceReq) any {
	mgr.Disband(msg.GetAllianceId())
	return &inner.DisbandAllianceReq{}
})
