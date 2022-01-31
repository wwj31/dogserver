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

var spawner = map[string]factory{
	"inner.ItemMap":               func() interface{} { return &ItemMap{} },
	"inner.Error":                 func() interface{} { return &Error{} },
	"inner.GateMsgWrapper":        func() interface{} { return &GateMsgWrapper{} },
	"inner.GameMsgWrapper":        func() interface{} { return &GameMsgWrapper{} },
	"inner.G2LRoleOffline":        func() interface{} { return &G2LRoleOffline{} },
	"inner.G2DGameStop":           func() interface{} { return &G2DGameStop{} },
	"inner.L2GTSessionDisabled":   func() interface{} { return &L2GTSessionDisabled{} },
	"inner.GT2GSessionClosed":     func() interface{} { return &GT2GSessionClosed{} },
	"inner.L2GTSessionAssignGame": func() interface{} { return &L2GTSessionAssignGame{} },
}

func Put(name string, x interface{}) {}
