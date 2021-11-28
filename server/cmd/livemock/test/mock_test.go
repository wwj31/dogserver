package test

import (
	"errors"
	"server/cmd/livemock/foo"
	"server/cmd/livemock/mock"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/smartystreets/goconvey/convey"
)

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
