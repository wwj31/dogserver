package main

import (
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
)

// 说明：
//   改良Test2中多次测试Live缺乏表达力的问题，
//   通过Convey将每一个测试用例包裹起来，通过So执行断言判断，
//   Convey可以嵌套使用，以表达测试不同状态的关系，
//   注意：最外层Convey二参数必须传入t

// 缺点：
//   测试数据本质还是由stub技术构建，整个测试环境，
//   没有很好的把输入连接起来。

func Test_Live3(t *testing.T) {
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
	convey.Convey("Live", t, func() {
		t.Log("LOG: Live")
		convey.Convey("GoodGoodStudy error", func() {
			t.Log("LOG: GoodGoodStudy error")
			convey.So(Live(0, 0, 0), convey.ShouldBeError)
		})
		convey.Convey("GoodGoodStudy ok", func() {
			t.Log("LOG: GoodGoodStudy ok")
			convey.Convey("BuyHouse error", func() {
				t.Log("LOG: BuyHouse error")
				convey.So(Live(100, 100, 100), convey.ShouldBeError)
			})
			convey.Convey("BuyHouse ok", func() {
				t.Log("LOG: BuyHouse ok")
				convey.Convey("Marry error", func() {
					t.Log("LOG: Marry error")
					convey.So(Live(100, 100, 100), convey.ShouldBeError)
				})
				convey.Convey("Marry ok", func() {
					t.Log("LOG: Marry ok")
					convey.So(Live(100, 100, 100), convey.ShouldBeNil)
				})
			})
		})
	})
}
