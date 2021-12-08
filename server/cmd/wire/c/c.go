package c

import (
	"server/a"
)

type ModelC struct {
	a *a.ModelA
	//b *b.ModelB
}

func New(a *a.ModelA) *ModelC {
	return &ModelC{
		a:a,
		//b:b,
	}
}

func (s *ModelC)Print()  {
	s.a.FA()
	//s.b.FB()
}