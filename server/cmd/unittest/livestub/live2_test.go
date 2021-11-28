package main

import (
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

// 说明：
//   一次性、依次确定桩函数的多个输出output，
//   避免了Test_Live1 中重复的坏味道。
// 缺点：
//   Live集中执行的次数没有层次，缺乏表达能力，
//   必须结合上下文分析，才能知道每一次测试用例
//   测试的目标是什么。
func Test_Live2(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()
	output := []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 3},
	}
	patches.ApplyFuncSeq(GoodGoodStudy, output)
	output = []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 2},
	}
	patches.ApplyFuncSeq(BuyHouse, output)
	output = []gomonkey.OutputCell{
		{Values: gomonkey.Params{errors.New("error")}, Times: 1},
		{Values: gomonkey.Params{nil}, Times: 1},
	}
	patches.ApplyFuncSeq(Marry, output)
	// GoodGoodStudy error
	assert.Error(t, Live(100, 100, 100))
	// BuyHouse error
	assert.Error(t, Live(100, 100, 100))
	// Marry error
	assert.Error(t, Live(100, 100, 100))
	// ok
	assert.NoError(t, Live(100, 100, 100))
}
