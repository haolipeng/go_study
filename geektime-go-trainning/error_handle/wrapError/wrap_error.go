package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

//要验证的内容
//%v 使用 %v 作为格式化参数，那么错误信息会保持一行， 其中依次包含调用栈的上下文文本。
//%+v 会输出完整的调用栈详情
//打印堆栈信息 无需多次Wrap，会导致调用堆栈重复，不利于后续定位问题
//Cause方法用于判断底层错误

/////////////////////////Wrap方法使用///////////////////////////
func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	err := foo()
	return errors.WithMessage(err, "bar failed")
}

func main() {
	//err := ReadFirst()
	//fmt.Printf("%+v", err)

	err := bar()

	//判断错误根因
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}
}
