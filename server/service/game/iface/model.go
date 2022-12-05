package iface

import gogo "github.com/gogo/protobuf/proto"

type Modeler interface {
	OnSave() gogo.Message
	OnLogin()
	OnLogout()
}
