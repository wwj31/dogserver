package iface

import (
	gogo "github.com/gogo/protobuf/proto"
	"server/proto/outermsg/outer"
)

type Modeler interface {
	OnLoaded()
	Data() gogo.Message
	OnLogin(first bool, enterGameRsp *outer.EnterGameRsp)
	OnLogout()
}
