package test

import (
	"errors"
	"server/cmd/unittest/livemock/foo"
	"server/cmd/unittest/livemock/mock"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/smartystreets/goconvey/convey"
)

// 说明：
//   mock和stub本质是两种不同的构建测试数据的技术，
//   使用mock必须保证功能实现了接口，
//   模块和模块依赖的是接口，而不是实现。
//   这样能很好的适配整个调用过程，无需打桩修改调用函数。
func Test_Live(t *testing.T) {
	ctrl := gomock.NewController(t)
	life := mock.NewMockLife(ctrl)
	handler := func(money int64) error {
		if money <= 0 {
			return errors.New("error")
		}
		return nil
	}
	life.EXPECT().GoodGoodStudy(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().BuyHouse(gomock.Any()).AnyTimes().DoAndReturn(handler)
	life.EXPECT().Marry(gomock.Any()).AnyTimes().DoAndReturn(handler)
	convey.Convey("Live", t, func() {
		person := foo.New(life)
		convey.Convey("GoodGoodStudy error", func() {
			convey.So(person.Live(0, 100, 100), convey.ShouldBeError)
		})
		convey.Convey("GoodGoodStudy ok", func() {
			convey.Convey("BuyHouse error", func() {
				convey.So(person.Live(100, 0, 100), convey.ShouldBeError)
			})
			convey.Convey("BuyHouse ok", func() {
				convey.Convey("Marry error", func() {
					convey.So(person.Live(100, 100, 0), convey.ShouldBeError)
				})
				convey.Convey("Marry ok", func() {
					convey.So(person.Live(100, 100, 100), convey.ShouldBeNil)
				})
			})
		})
	})
}
