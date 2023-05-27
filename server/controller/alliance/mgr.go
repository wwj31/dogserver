package alliance

import (
	"server/proto/innermsg/inner"
	"server/service/alliance"

	"server/common/router"
)

// 玩家登录，同步并请求数据
var _ = router.Reg(func(alliance *alliance.Mgr, msg *inner.CreateAllianceReq) any {
	// todo
	return &inner.CreateAllianceRsp{}
})
