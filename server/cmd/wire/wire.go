//go:build wireinject

package main

import (
	"github.com/google/wire"
	"server/b"
	"server/c"
)

func InitModelC() *c.ModelC {
	wire.Build(MySet)
	return nil
}

func InitModelB(str string) *b.ModelB {
	wire.Build(MySet)
	return nil
}
