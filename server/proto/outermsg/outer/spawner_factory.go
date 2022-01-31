// Code generated by "spawner -pool=true"; DO NOT EDIT.

package outer

type factory func() interface{}

func Spawner(name string, newPool ...bool) (interface{}, bool) {
	if len(newPool) > 0 && newPool[0] {
		p, ok := spawnerPools[name]
		if !ok {
			return nil, false
		}
		return p.Get(), true
	}
	f, ok := spawner[name]
	if !ok {
		return nil, ok
	}
	return f(), true
}

var spawner = map[string]factory{
	"outer.Mail":                func() interface{} { return &Mail{} },
	"outer.MailInfo":            func() interface{} { return &MailInfo{} },
	"outer.Ok":                  func() interface{} { return &Ok{} },
	"outer.Fail":                func() interface{} { return &Fail{} },
	"outer.Unknown":             func() interface{} { return &Unknown{} },
	"outer.UseItemResp":         func() interface{} { return &UseItemResp{} },
	"outer.UseItemReq":          func() interface{} { return &UseItemReq{} },
	"outer.ItemChangeNotify":    func() interface{} { return &ItemChangeNotify{} },
	"outer.ItemInfoPush":        func() interface{} { return &ItemInfoPush{} },
	"outer.LoginReq":            func() interface{} { return &LoginReq{} },
	"outer.EnterGameReq":        func() interface{} { return &EnterGameReq{} },
	"outer.Ping":                func() interface{} { return &Ping{} },
	"outer.EnterGameResp":       func() interface{} { return &EnterGameResp{} },
	"outer.LoginResp":           func() interface{} { return &LoginResp{} },
	"outer.RoleInfoPush":        func() interface{} { return &RoleInfoPush{} },
	"outer.Pong":                func() interface{} { return &Pong{} },
	"outer.MailListResp":        func() interface{} { return &MailListResp{} },
	"outer.AddMailNotify":       func() interface{} { return &AddMailNotify{} },
	"outer.ReadMailResp":        func() interface{} { return &ReadMailResp{} },
	"outer.ReceiveMailItemReq":  func() interface{} { return &ReceiveMailItemReq{} },
	"outer.ReceiveMailItemResp": func() interface{} { return &ReceiveMailItemResp{} },
	"outer.MailListReq":         func() interface{} { return &MailListReq{} },
	"outer.ReadMailReq":         func() interface{} { return &ReadMailReq{} },
	"outer.DeleteMailReq":       func() interface{} { return &DeleteMailReq{} },
}

func Put(name string, x interface{}) {
	pool, ok := spawnerPools[name]
	if !ok {
		return
	}
	pool.Put(x)
}
