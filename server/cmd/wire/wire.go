// +build wireinject

package main

import (
	"github.com/google/wire"
	"server/c"
)

func InitModelC() *c.ModelC {
	wire.Build(MySet)
	return nil
}
