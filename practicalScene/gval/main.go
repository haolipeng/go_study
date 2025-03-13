package main

import (
	"context"
	"fmt"

	"github.com/PaesslerAG/gval"
)

func testNestedFiled() {
	// 修改输入数据为三级嵌套结构
	data := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"abc": -1, // 确保 bar 也是一个嵌套的 map
			},
		},
	}

	// 修改表达式为三级字段访问
	value, err := gval.Evaluate("foo.bar.abc > 0", data)
	if err != nil {
		fmt.Println("评估错误:", err)
		return
	}

	fmt.Println("结果:", value) // 输出: 结果: true (因为 1 > 0)
}

func ParseAndEval() {
	eval, err := gval.Full(gval.Constant("maximum_time", 52)).
		NewEvaluable("response_time <= maximum_time")
	if err != nil {
		fmt.Println(err)
	}

	for i := 50; i < 55; i++ {
		value, err := eval(context.Background(), map[string]interface{}{
			"response_time": i,
		})
		if err != nil {
			fmt.Println(err)

		}

		fmt.Println(value)
	}
}

func main() {
	//ParseAndEval()
	testNestedFiled()
	return
}
