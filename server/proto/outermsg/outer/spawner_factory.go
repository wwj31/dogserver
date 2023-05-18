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
	"outer.BindPhoneReq":       func() interface{} { return &BindPhoneReq{} },
	"outer.BindPhoneRsp":       func() interface{} { return &BindPhoneRsp{} },
	"outer.LoginRsp":           func() interface{} { return &LoginRsp{} },
	"outer.RoleInfo":           func() interface{} { return &RoleInfo{} },
	"outer.HeartReq":           func() interface{} { return &HeartReq{} },
	"outer.EnterGameRsp":       func() interface{} { return &EnterGameRsp{} },
	"outer.HeartRsp":           func() interface{} { return &HeartRsp{} },
	"outer.SendSMSRsp":         func() interface{} { return &SendSMSRsp{} },
	"outer.EnterGameReq":       func() interface{} { return &EnterGameReq{} },
	"outer.SendSMSReq":         func() interface{} { return &SendSMSReq{} },
	"outer.LoginReq":           func() interface{} { return &LoginReq{} },
	"outer.MailListRsp":        func() interface{} { return &MailListRsp{} },
	"outer.DeleteMailReq":      func() interface{} { return &DeleteMailReq{} },
	"outer.Mail":               func() interface{} { return &Mail{} },
	"outer.MailListReq":        func() interface{} { return &MailListReq{} },
	"outer.ReadMailRsp":        func() interface{} { return &ReadMailRsp{} },
	"outer.MailInfo":           func() interface{} { return &MailInfo{} },
	"outer.ReceiveMailItemReq": func() interface{} { return &ReceiveMailItemReq{} },
	"outer.AddMailNotify":      func() interface{} { return &AddMailNotify{} },
	"outer.ReadMailReq":        func() interface{} { return &ReadMailReq{} },
	"outer.ReceiveMailItemRsp": func() interface{} { return &ReceiveMailItemRsp{} },
}
