package iface

import gogo "github.com/gogo/protobuf/proto"

type Modeler interface {
	OnLoaded()
	Data() gogo.Message
	OnLogin(first bool)
	OnLogout()
}
