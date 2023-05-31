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
	"inner.DisbandAllianceReq":    func() interface{} { return &DisbandAllianceReq{} },
	"inner.AllianceInfoReq":       func() interface{} { return &AllianceInfoReq{} },
	"inner.AllianceInfoRsp":       func() interface{} { return &AllianceInfoRsp{} },
	"inner.DisbandAllianceRsp":    func() interface{} { return &DisbandAllianceRsp{} },
	"inner.MemberInfoOnLoginRsp":  func() interface{} { return &MemberInfoOnLoginRsp{} },
	"inner.AddMemberReq":          func() interface{} { return &AddMemberReq{} },
	"inner.CreateAllianceReq":     func() interface{} { return &CreateAllianceReq{} },
	"inner.MemberInfoOnLoginReq":  func() interface{} { return &MemberInfoOnLoginReq{} },
	"inner.AddMemberRsp":          func() interface{} { return &AddMemberRsp{} },
	"inner.SetMemberPositionReq":  func() interface{} { return &SetMemberPositionReq{} },
	"inner.AllianceDisbandedNtf":  func() interface{} { return &AllianceDisbandedNtf{} },
	"inner.CreateAllianceRsp":     func() interface{} { return &CreateAllianceRsp{} },
	"inner.SetMemberPositionRsp":  func() interface{} { return &SetMemberPositionRsp{} },
	"inner.MemberInfoOnLogoutReq": func() interface{} { return &MemberInfoOnLogoutReq{} },
	"inner.PlayerInfo":            func() interface{} { return &PlayerInfo{} },
	"inner.PullPlayer":            func() interface{} { return &PullPlayer{} },
	"inner.NewPlayerInfo":         func() interface{} { return &NewPlayerInfo{} },
	"inner.Error":                 func() interface{} { return &Error{} },
	"inner.Ok":                    func() interface{} { return &Ok{} },
	"inner.GateMsgWrapper":        func() interface{} { return &GateMsgWrapper{} },
	"inner.KickOutReq":            func() interface{} { return &KickOutReq{} },
	"inner.KickOutRsp":            func() interface{} { return &KickOutRsp{} },
	"inner.GSessionClosed":        func() interface{} { return &GSessionClosed{} },
	"inner.BindSessionWithRID":    func() interface{} { return &BindSessionWithRID{} },
	"inner.Mail":                  func() interface{} { return &Mail{} },
	"inner.MailInfo":              func() interface{} { return &MailInfo{} },
	"inner.RoleInfo":              func() interface{} { return &RoleInfo{} },
	"inner.AllianceInfo":          func() interface{} { return &AllianceInfo{} },
}
