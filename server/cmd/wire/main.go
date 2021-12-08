package main

import (
	"github.com/google/wire"
	"server/a"
	"server/b"
	"server/c"
	"server/iface"
)



var MySet = wire.NewSet(
	wire.Bind(new(iface.A), new(*a.ModelA)),
	//wire.Bind(new(iface.B), new(*b.ModelB)),
	//wire.Struct(a.ModelA{}),
	//wire.Struct(new(b.ModelB)),
	a.NewA,
	a.NewA2,
	b.New,
	c.New,
	)

func main()  {
	modelC := InitModelC()
	modelC.Print()
}