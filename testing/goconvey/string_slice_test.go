package goconvey

import (
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

//Convey用于定义测试，Convey可以嵌套使用
//So用来做判断
func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual的描述", t, func() {
		//一次测试了多种情况
		Convey("true when a!=nil && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("TestStringSliceEqual should return true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("TestStringSliceEqual should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("TestStringSliceEqual should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}

func TestStringSliceEqualWithNested(t *testing.T) {
	Convey("TestStringSliceEqual", t, func() {
		Convey("should return true when a != nil  && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("should return true when a ＝= nil  && b ＝= nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})

		Convey("should return false when a ＝= nil  && b != nil", func() {
			a := []string(nil)
			b := []string{}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})

		Convey("should return false when a != nil  && b != nil", func() {
			a := []string{"hello", "world"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeFalse)
		})
	})
}
