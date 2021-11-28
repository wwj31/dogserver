package main

import (
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// 说明：
//   最基础的单元测试，将依赖的函数进行打桩处理，
//   然后依次断言执行要测试的函数Live()。
// 缺点：
//   对于每一次测试，都需要重新构建桩函数，以下示例中
//   总共构建了4次GoodGoodStudy桩函数。
func Test_Live1(t *testing.T) {
	patches := gomonkey.NewPatches()
	// GoodGoodStudy error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// BuyHouse error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// Marry error
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return nil
	})
	patches.ApplyFunc(Marry, func(int64) error {
		return errors.New("error")
	})
	assert.Error(t, Live(100, 100, 100))
	patches.Reset()
	// ok
	patches.ApplyFunc(GoodGoodStudy, func(int64) error {
		return nil
	})
	patches.ApplyFunc(BuyHouse, func(int64) error {
		return nil
	})
	patches.ApplyFunc(Marry, func(int64) error {
		return nil
	})
	assert.NoError(t, Live(100, 100, 100))
	patches.Reset()
}
