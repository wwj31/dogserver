package common

import (
	"math/rand"
)

// 按权重随机取一个数
// rands map[数]权重值
// 返回取出来的数
func RandomWeight32(rands map[int64]int64) int64 {
	total := int64(0)
	arrangement := [][]int64{}
	for val, weight := range rands {
		total += weight
		arrangement = append(arrangement, []int64{val, weight}) // {{4:1000},{1:2100},{6:3000},{5:500}}
	}
	// 在 0~权重总和 中随机取一个值，然后判断该值落在哪个数的权重中
	retValue := rand.Int63n(total)
	addValue := int64(0)
	for _, v := range arrangement {
		if addValue <= retValue && retValue < (addValue+v[1]) {
			return v[0]
		}
		addValue += v[1]
	}

	return -1
}
