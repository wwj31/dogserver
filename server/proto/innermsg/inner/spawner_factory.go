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
	"inner.SetMemberPositionRsp":  func() interface{} { return &SetMemberPositionRsp{} },
	"inner.SetMemberPositionReq":  func() interface{} { return &SetMemberPositionReq{} },
	"inner.RoomLogoutRsp":         func() interface{} { return &RoomLogoutRsp{} },
	"inner.RoomLogoutReq":         func() interface{} { return &RoomLogoutReq{} },
	"inner.RoomLoginRsp":          func() interface{} { return &RoomLoginRsp{} },
	"inner.RoomLoginReq":          func() interface{} { return &RoomLoginReq{} },
	"inner.RoomInfoRsp":           func() interface{} { return &RoomInfoRsp{} },
	"inner.RoomInfoReq":           func() interface{} { return &RoomInfoReq{} },
	"inner.RoomInfo":              func() interface{} { return &RoomInfo{} },
	"inner.RoleInfo":              func() interface{} { return &RoleInfo{} },
	"inner.PullPlayer":            func() interface{} { return &PullPlayer{} },
	"inner.PlayerInfo":            func() interface{} { return &PlayerInfo{} },
	"inner.Ok":                    func() interface{} { return &Ok{} },
	"inner.NewPlayerInfo":         func() interface{} { return &NewPlayerInfo{} },
	"inner.MemberInfoOnLogoutRsp": func() interface{} { return &MemberInfoOnLogoutRsp{} },
	"inner.MemberInfoOnLogoutReq": func() interface{} { return &MemberInfoOnLogoutReq{} },
	"inner.MemberInfoOnLoginRsp":  func() interface{} { return &MemberInfoOnLoginRsp{} },
	"inner.MemberInfoOnLoginReq":  func() interface{} { return &MemberInfoOnLoginReq{} },
	"inner.MailInfo":              func() interface{} { return &MailInfo{} },
	"inner.Mail":                  func() interface{} { return &Mail{} },
	"inner.LeaveRoomRsp":          func() interface{} { return &LeaveRoomRsp{} },
	"inner.LeaveRoomReq":          func() interface{} { return &LeaveRoomReq{} },
	"inner.KickOutRsp":            func() interface{} { return &KickOutRsp{} },
	"inner.KickOutReq":            func() interface{} { return &KickOutReq{} },
	"inner.KickOutMembersRsp":     func() interface{} { return &KickOutMembersRsp{} },
	"inner.KickOutMembersReq":     func() interface{} { return &KickOutMembersReq{} },
	"inner.JoinRoomRsp":           func() interface{} { return &JoinRoomRsp{} },
	"inner.JoinRoomReq":           func() interface{} { return &JoinRoomReq{} },
	"inner.GateMsgWrapper":        func() interface{} { return &GateMsgWrapper{} },
	"inner.GSessionClosed":        func() interface{} { return &GSessionClosed{} },
	"inner.Error":                 func() interface{} { return &Error{} },
	"inner.DisbandRoomRsp":        func() interface{} { return &DisbandRoomRsp{} },
	"inner.DisbandRoomReq":        func() interface{} { return &DisbandRoomReq{} },
	"inner.DisbandAllianceRsp":    func() interface{} { return &DisbandAllianceRsp{} },
	"inner.DisbandAllianceReq":    func() interface{} { return &DisbandAllianceReq{} },
	"inner.CreateRoomRsp":         func() interface{} { return &CreateRoomRsp{} },
	"inner.CreateRoomReq":         func() interface{} { return &CreateRoomReq{} },
	"inner.CreateAllianceRsp":     func() interface{} { return &CreateAllianceRsp{} },
	"inner.CreateAllianceReq":     func() interface{} { return &CreateAllianceReq{} },
	"inner.BindSessionWithRID":    func() interface{} { return &BindSessionWithRID{} },
	"inner.AllianceInfoRsp":       func() interface{} { return &AllianceInfoRsp{} },
	"inner.AllianceInfoReq":       func() interface{} { return &AllianceInfoReq{} },
	"inner.AllianceInfoNtf":       func() interface{} { return &AllianceInfoNtf{} },
	"inner.AllianceInfo":          func() interface{} { return &AllianceInfo{} },
	"inner.AddMemberRsp":          func() interface{} { return &AddMemberRsp{} },
	"inner.AddMemberReq":          func() interface{} { return &AddMemberReq{} },
	"inner.Account":               func() interface{} { return &Account{} },
}
