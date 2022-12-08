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
	"inner.JoinChannelResp":    func() interface{} { return &JoinChannelResp{} },
	"inner.LeaveChannelReq":    func() interface{} { return &LeaveChannelReq{} },
	"inner.JoinChannelReq":     func() interface{} { return &JoinChannelReq{} },
	"inner.MessageToChannel":   func() interface{} { return &MessageToChannel{} },
	"inner.Error":              func() interface{} { return &Error{} },
	"inner.BindSessionWithRID": func() interface{} { return &BindSessionWithRID{} },
	"inner.PullPlayer":         func() interface{} { return &PullPlayer{} },
	"inner.GateMsgWrapper":     func() interface{} { return &GateMsgWrapper{} },
	"inner.GSessionClosed":     func() interface{} { return &GSessionClosed{} },
	"inner.ItemInfo":           func() interface{} { return &ItemInfo{} },
	"inner.MailInfo":           func() interface{} { return &MailInfo{} },
	"inner.Mail":               func() interface{} { return &Mail{} },
	"inner.RoleInfo":           func() interface{} { return &RoleInfo{} },
}
