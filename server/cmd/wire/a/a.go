package a

import "fmt"

type ModelA struct {
	Name string
}

func NewA() *ModelA {
	return &ModelA{
		Name: "A",
	}
}

func (s *ModelA) FA() {
	fmt.Println(s.Name)
}

type ModelA2 struct {
	name string
}

func NewA2() *ModelA2 {
	return &ModelA2{
		name: "A2",
	}
}

func (s *ModelA2) FA() {
	fmt.Println(s.name)
}
