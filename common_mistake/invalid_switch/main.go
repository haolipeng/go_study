package main

import "fmt"

//你可以通过在每个“case”块的结尾使用“fallthrough”，来强制“case”代码块进入。
//你也可以重写switch语句，来使用“case”块中的表达式列表。
func main() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ': //这个case不会return true，除非使用fallthrough
		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //true
	fmt.Println(isSpace(' '))  //false

	//正确写法
	isSpaceGood := func(ch byte) bool {
		switch ch {
		case ' ', '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpaceGood('\t')) //true
	fmt.Println(isSpaceGood(' '))  //true
}
