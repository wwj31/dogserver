// Code generated by "spawner -pool=false"; DO NOT EDIT.

package outer

type factory func() interface{}

func Spawner(name string, newPool ...bool) (interface{}, bool) {
	f, ok := spawner[name]
	if !ok {
		return nil, ok
	}
	return f(), true
}

func Put(name string, x interface{}) {}

var spawner = map[string]factory{
	"outer.AgentMembersReq":    func() interface{} { return &AgentMembersReq{} },
	"outer.AgentMembersRsp":    func() interface{} { return &AgentMembersRsp{} },
	"outer.JoinAllianceNtf":    func() interface{} { return &JoinAllianceNtf{} },
	"outer.Base":               func() interface{} { return &Base{} },
	"outer.PlayerInfo":         func() interface{} { return &PlayerInfo{} },
	"outer.Unknown":            func() interface{} { return &Unknown{} },
	"outer.FailRsp":            func() interface{} { return &FailRsp{} },
	"outer.SetRoleInfoRsp":     func() interface{} { return &SetRoleInfoRsp{} },
	"outer.BindPhoneRsp":       func() interface{} { return &BindPhoneRsp{} },
	"outer.ModifyPasswordRsp":  func() interface{} { return &ModifyPasswordRsp{} },
	"outer.SetRoleInfoReq":     func() interface{} { return &SetRoleInfoReq{} },
	"outer.BindPhoneReq":       func() interface{} { return &BindPhoneReq{} },
	"outer.ModifyPasswordReq":  func() interface{} { return &ModifyPasswordReq{} },
	"outer.RoleInfo":           func() interface{} { return &RoleInfo{} },
	"outer.LoginRsp":           func() interface{} { return &LoginRsp{} },
	"outer.LoginReq":           func() interface{} { return &LoginReq{} },
	"outer.HeartRsp":           func() interface{} { return &HeartRsp{} },
	"outer.EnterGameReq":       func() interface{} { return &EnterGameReq{} },
	"outer.EnterGameRsp":       func() interface{} { return &EnterGameRsp{} },
	"outer.HeartReq":           func() interface{} { return &HeartReq{} },
	"outer.MailInfo":           func() interface{} { return &MailInfo{} },
	"outer.ReadMailReq":        func() interface{} { return &ReadMailReq{} },
	"outer.ReceiveMailItemReq": func() interface{} { return &ReceiveMailItemReq{} },
	"outer.Mail":               func() interface{} { return &Mail{} },
	"outer.MailListReq":        func() interface{} { return &MailListReq{} },
	"outer.MailListRsp":        func() interface{} { return &MailListRsp{} },
	"outer.ReadMailRsp":        func() interface{} { return &ReadMailRsp{} },
	"outer.ReceiveMailItemRsp": func() interface{} { return &ReceiveMailItemRsp{} },
	"outer.AddMailNotify":      func() interface{} { return &AddMailNotify{} },
	"outer.DeleteMailReq":      func() interface{} { return &DeleteMailReq{} },
}
