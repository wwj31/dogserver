package common

func Min(x, y int64) int64 {
	if x <= y {
		return x
	}
	return y
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}
