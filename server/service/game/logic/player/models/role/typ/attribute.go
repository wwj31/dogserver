//go install golang.org/x/tools/cmd/stringer@latest

package typ

//go:generate stringer -type Attribute -linecomment
type Attribute int64

const (
	_     Attribute = iota
	Level           // 等级
	Exp             // 经验
	Gold            // 金币

	AttributeMax
)

func (s Attribute) Int64() int64 {
	return int64(s)
}

func (s Attribute) Int32() int32 {
	return int32(s)
}
