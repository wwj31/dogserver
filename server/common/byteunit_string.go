// Code generated by "stringer -type ByteUnit"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[B-1]
}

const _ByteUnit_name = "B"

var _ByteUnit_index = [...]uint8{0, 1}

func (i ByteUnit) String() string {
	i -= 1
	if i < 0 || i >= ByteUnit(len(_ByteUnit_index)-1) {
		return "ByteUnit(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ByteUnit_name[_ByteUnit_index[i]:_ByteUnit_index[i+1]]
}
