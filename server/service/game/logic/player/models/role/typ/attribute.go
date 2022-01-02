// Package typ go install golang.org/x/tools/cmd/stringer
//go:generate stringer -type Attribute -linecomment
package typ

type Attribute int64

const (
	_     Attribute = iota
	Level           // 等级
	Exp             // 经验
	Glod            // 金币
)

func (s Attribute) Int64() int64 {
	return int64(s)
}
