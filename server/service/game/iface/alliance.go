package iface

type Alliance interface {
	Modeler
	AllianceId() int32
	SetAllianceId(id int32)
}
