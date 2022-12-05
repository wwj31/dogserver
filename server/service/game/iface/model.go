package iface

type Modeler interface {
	OnSave()
	OnLogin()
	OnLogout()
}
