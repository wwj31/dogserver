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
	"outer.Base":               func() interface{} { return &Base{} },
	"outer.Unknown":            func() interface{} { return &Unknown{} },
	"outer.FailRsp":            func() interface{} { return &FailRsp{} },
	"outer.SetRoleInfoRsp":     func() interface{} { return &SetRoleInfoRsp{} },
	"outer.BindPhoneRsp":       func() interface{} { return &BindPhoneRsp{} },
	"outer.SetRoleInfoReq":     func() interface{} { return &SetRoleInfoReq{} },
	"outer.ModifyPasswordReq":  func() interface{} { return &ModifyPasswordReq{} },
	"outer.ModifyPasswordRsp":  func() interface{} { return &ModifyPasswordRsp{} },
	"outer.BindPhoneReq":       func() interface{} { return &BindPhoneReq{} },
	"outer.LoginReq":           func() interface{} { return &LoginReq{} },
	"outer.HeartReq":           func() interface{} { return &HeartReq{} },
	"outer.HeartRsp":           func() interface{} { return &HeartRsp{} },
	"outer.LoginRsp":           func() interface{} { return &LoginRsp{} },
	"outer.EnterGameRsp":       func() interface{} { return &EnterGameRsp{} },
	"outer.RoleInfo":           func() interface{} { return &RoleInfo{} },
	"outer.EnterGameReq":       func() interface{} { return &EnterGameReq{} },
	"outer.ReadMailRsp":        func() interface{} { return &ReadMailRsp{} },
	"outer.ReceiveMailItemRsp": func() interface{} { return &ReceiveMailItemRsp{} },
	"outer.Mail":               func() interface{} { return &Mail{} },
	"outer.MailInfo":           func() interface{} { return &MailInfo{} },
	"outer.MailListRsp":        func() interface{} { return &MailListRsp{} },
	"outer.ReadMailReq":        func() interface{} { return &ReadMailReq{} },
	"outer.MailListReq":        func() interface{} { return &MailListReq{} },
	"outer.AddMailNotify":      func() interface{} { return &AddMailNotify{} },
	"outer.ReceiveMailItemReq": func() interface{} { return &ReceiveMailItemReq{} },
	"outer.DeleteMailReq":      func() interface{} { return &DeleteMailReq{} },
}
