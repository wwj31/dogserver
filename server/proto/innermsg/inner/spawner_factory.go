// Code generated by "spawner "; DO NOT EDIT.

package inner

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
	"inner.Account":               func() interface{} { return &Account{} },
	"inner.AllianceInfoNtf":       func() interface{} { return &AllianceInfoNtf{} },
	"inner.DisbandAllianceRsp":    func() interface{} { return &DisbandAllianceRsp{} },
	"inner.SetMemberPositionRsp":  func() interface{} { return &SetMemberPositionRsp{} },
	"inner.KickOutMembersRsp":     func() interface{} { return &KickOutMembersRsp{} },
	"inner.AddMemberReq":          func() interface{} { return &AddMemberReq{} },
	"inner.MemberInfoOnLogoutReq": func() interface{} { return &MemberInfoOnLogoutReq{} },
	"inner.DisbandAllianceReq":    func() interface{} { return &DisbandAllianceReq{} },
	"inner.MemberInfoOnLoginReq":  func() interface{} { return &MemberInfoOnLoginReq{} },
	"inner.AllianceInfoRsp":       func() interface{} { return &AllianceInfoRsp{} },
	"inner.KickOutMembersReq":     func() interface{} { return &KickOutMembersReq{} },
	"inner.CreateAllianceReq":     func() interface{} { return &CreateAllianceReq{} },
	"inner.CreateAllianceRsp":     func() interface{} { return &CreateAllianceRsp{} },
	"inner.AddMemberRsp":          func() interface{} { return &AddMemberRsp{} },
	"inner.AllianceInfoReq":       func() interface{} { return &AllianceInfoReq{} },
	"inner.SetMemberPositionReq":  func() interface{} { return &SetMemberPositionReq{} },
	"inner.MemberInfoOnLoginRsp":  func() interface{} { return &MemberInfoOnLoginRsp{} },
	"inner.MemberInfoOnLogoutRsp": func() interface{} { return &MemberInfoOnLogoutRsp{} },
	"inner.PlayerInfo":            func() interface{} { return &PlayerInfo{} },
	"inner.Ok":                    func() interface{} { return &Ok{} },
	"inner.GateMsgWrapper":        func() interface{} { return &GateMsgWrapper{} },
	"inner.PullPlayer":            func() interface{} { return &PullPlayer{} },
	"inner.Error":                 func() interface{} { return &Error{} },
	"inner.NewPlayerInfo":         func() interface{} { return &NewPlayerInfo{} },
	"inner.BindSessionWithRID":    func() interface{} { return &BindSessionWithRID{} },
	"inner.GSessionClosed":        func() interface{} { return &GSessionClosed{} },
	"inner.KickOutReq":            func() interface{} { return &KickOutReq{} },
	"inner.KickOutRsp":            func() interface{} { return &KickOutRsp{} },
	"inner.MailInfo":              func() interface{} { return &MailInfo{} },
	"inner.RoleInfo":              func() interface{} { return &RoleInfo{} },
	"inner.AllianceInfo":          func() interface{} { return &AllianceInfo{} },
	"inner.Mail":                  func() interface{} { return &Mail{} },
	"inner.RoomInfo":              func() interface{} { return &RoomInfo{} },
	"inner.CreateRoomRsp":         func() interface{} { return &CreateRoomRsp{} },
	"inner.CreateRoomReq":         func() interface{} { return &CreateRoomReq{} },
}
