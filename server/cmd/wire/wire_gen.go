// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"server/a"
	"server/c"
)

// Injectors from wire.go:

func InitModelC() *c.ModelC {
	modelA := &a.ModelA{}
	modelC := c.New(modelA)
	return modelC
}