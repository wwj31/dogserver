//go install golang.org/x/tools/cmd/stringer

package typ

//go:generate stringer -type Attribute -linecomment
type Attribute int64

const (
	_     Attribute = iota
	Level           // 等级
	Exp             // 经验
	Glod            // 金币

	Attribute_max
)

func (s Attribute) Int64() int64 {
	return int64(s)
}
