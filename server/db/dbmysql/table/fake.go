package table

import "reflect"

type Fake struct {
	Id   uint64 `gorm:"primary_key"`
	Data string
}

func (f Fake) ModelName() string {
	return reflect.TypeOf(f).Name()
}

func (f Fake) SplitNum() int {
	return 0
}

func (f Fake) Key() uint64 {
	return f.Id
}
