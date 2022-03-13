package iface

import gogo "github.com/gogo/protobuf/proto"

type Chat interface {
	Modeler
	SendToChannel(channel string, msg gogo.Message)
}
