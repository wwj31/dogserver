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
	"inner.JoinChannelReq":     func() interface{} { return &JoinChannelReq{} },
	"inner.LeaveChannelReq":    func() interface{} { return &LeaveChannelReq{} },
	"inner.MessageToChannel":   func() interface{} { return &MessageToChannel{} },
	"inner.BindSessionWithRID": func() interface{} { return &BindSessionWithRID{} },
	"inner.KickOutReq":         func() interface{} { return &KickOutReq{} },
	"inner.GateMsgWrapper":     func() interface{} { return &GateMsgWrapper{} },
	"inner.KickOutRsp":         func() interface{} { return &KickOutRsp{} },
	"inner.Error":              func() interface{} { return &Error{} },
	"inner.GSessionClosed":     func() interface{} { return &GSessionClosed{} },
	"inner.PullPlayer":         func() interface{} { return &PullPlayer{} },
	"inner.MailInfo":           func() interface{} { return &MailInfo{} },
	"inner.RoleInfo":           func() interface{} { return &RoleInfo{} },
	"inner.Mail":               func() interface{} { return &Mail{} },
	"inner.ItemInfo":           func() interface{} { return &ItemInfo{} },
}
