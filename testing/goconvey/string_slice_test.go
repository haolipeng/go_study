package goconvey

import (
	"testing"
)
import . "github.com/smartystreets/goconvey/convey"

//Convey用于定义测试
//So用来做判断
//Convey可以嵌套使用
func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual的描述", t, func() {
		Convey("true when a!=nil && b != nil", func() {
			a := []string{"hello", "goconvey"}
			b := []string{"hello", "goconvey"}
			So(StringSliceEqual(a, b), ShouldBeTrue)
		})

		Convey("true when a == nil && b == nil", func() {
			So(StringSliceEqual(nil, nil), ShouldBeTrue)
		})
	})
}
