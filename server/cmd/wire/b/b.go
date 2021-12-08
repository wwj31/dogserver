package b

import "fmt"

type ModelB struct {
	name string
}

func New() *ModelB {
	return &ModelB{
		name: "B",
	}
}

func (s *ModelB)FB()  {
	fmt.Println(s.name)
}