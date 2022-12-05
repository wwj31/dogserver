package b

import "fmt"

type ModelB struct {
	name string
}

func New(n string) *ModelB {
	return &ModelB{
		name: n,
	}
}

func (s *ModelB) FB() {
	fmt.Println(s.name)
}
