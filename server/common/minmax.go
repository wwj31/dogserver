package common

type NUMBER interface {
	int | int8 | int16 | int32 | int64
}

func Max[Number NUMBER](a, b Number) Number {
	if a > b {
		return a
	}
	return b
}

func Min[Number NUMBER](a, b Number) Number {
	if a < b {
		return a
	}
	return b
}

func Abs[Number NUMBER](x Number) Number {
	if x >= 0 {
		return x
	}
	return -x
}
