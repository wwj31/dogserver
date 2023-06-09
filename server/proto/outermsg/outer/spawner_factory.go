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
	"outer.AgentMembersReq":      func() interface{} { return &AgentMembersReq{} },
	"outer.AgentMembersRsp":      func() interface{} { return &AgentMembersRsp{} },
	"outer.InviteAllianceReq":    func() interface{} { return &InviteAllianceReq{} },
	"outer.SetMemberPositionReq": func() interface{} { return &SetMemberPositionReq{} },
	"outer.DisbandAllianceRsp":   func() interface{} { return &DisbandAllianceRsp{} },
	"outer.KickOutMemberRsp":     func() interface{} { return &KickOutMemberRsp{} },
	"outer.DisbandAllianceReq":   func() interface{} { return &DisbandAllianceReq{} },
	"outer.KickOutMemberReq":     func() interface{} { return &KickOutMemberReq{} },
	"outer.InviteAllianceRsp":    func() interface{} { return &InviteAllianceRsp{} },
	"outer.SetMemberPositionRsp": func() interface{} { return &SetMemberPositionRsp{} },
	"outer.AllianceInfoNtf":      func() interface{} { return &AllianceInfoNtf{} },
	"outer.PlayerInfo":           func() interface{} { return &PlayerInfo{} },
	"outer.Base":                 func() interface{} { return &Base{} },
	"outer.FailRsp":              func() interface{} { return &FailRsp{} },
	"outer.Unknown":              func() interface{} { return &Unknown{} },
	"outer.SetRoleInfoRsp":       func() interface{} { return &SetRoleInfoRsp{} },
	"outer.BindPhoneRsp":         func() interface{} { return &BindPhoneRsp{} },
	"outer.ModifyPasswordRsp":    func() interface{} { return &ModifyPasswordRsp{} },
	"outer.SetRoleInfoReq":       func() interface{} { return &SetRoleInfoReq{} },
	"outer.BindPhoneReq":         func() interface{} { return &BindPhoneReq{} },
	"outer.ModifyPasswordReq":    func() interface{} { return &ModifyPasswordReq{} },
	"outer.EnterGameRsp":         func() interface{} { return &EnterGameRsp{} },
	"outer.HeartRsp":             func() interface{} { return &HeartRsp{} },
	"outer.LoginReq":             func() interface{} { return &LoginReq{} },
	"outer.LoginRsp":             func() interface{} { return &LoginRsp{} },
	"outer.EnterGameReq":         func() interface{} { return &EnterGameReq{} },
	"outer.HeartReq":             func() interface{} { return &HeartReq{} },
	"outer.RoleInfo":             func() interface{} { return &RoleInfo{} },
	"outer.AddMailNotify":        func() interface{} { return &AddMailNotify{} },
	"outer.MailListRsp":          func() interface{} { return &MailListRsp{} },
	"outer.DeleteMailReq":        func() interface{} { return &DeleteMailReq{} },
	"outer.MailInfo":             func() interface{} { return &MailInfo{} },
	"outer.ReadMailRsp":          func() interface{} { return &ReadMailRsp{} },
	"outer.ReceiveMailItemReq":   func() interface{} { return &ReceiveMailItemReq{} },
	"outer.Mail":                 func() interface{} { return &Mail{} },
	"outer.ReadMailReq":          func() interface{} { return &ReadMailReq{} },
	"outer.ReceiveMailItemRsp":   func() interface{} { return &ReceiveMailItemRsp{} },
	"outer.MailListReq":          func() interface{} { return &MailListReq{} },
}
