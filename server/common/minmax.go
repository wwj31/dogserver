package common

func Max[Number int | int8 | int16 | int32 | int64](a, b Number) Number {
	if a > b {
		return a
	}
	return b
}

func Min[Number int | int8 | int16 | int32 | int64](a, b Number) Number {
	if a < b {
		return a
	}
	return b
}

func Abs[Number int | int8 | int16 | int32 | int64](x Number) Number {
	if x >= 0 {
		return x
	}
	return -x
}
