package c

import (
	"server/iface"
)

type ModelC struct {
	a iface.A
	//b *b.ModelB
}

func New(a iface.A) *ModelC {
	return &ModelC{
		a: a,
		//b:b,
	}
}

func (s *ModelC) Print() {
	s.a.FA()
	//s.b.FB()
}
