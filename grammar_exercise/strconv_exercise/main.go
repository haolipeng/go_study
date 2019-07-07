package strconv_exercise

import (
	"fmt"
	"strconv"
)

/*
strconv是用来做数据转换的库
使用最多的是Atoi和Itoa
*/
func string2int() {
	i, _ := strconv.Atoi("-42")
	fmt.Println("after convert,int value is ", i)

	s := strconv.Itoa(-42)
	fmt.Println("after convert,string value is ", s)
}

func useParseBool() {
	fmt.Println(strconv.ParseBool("1"))    // true
	fmt.Println(strconv.ParseBool("t"))    // true
	fmt.Println(strconv.ParseBool("T"))    // true
	fmt.Println(strconv.ParseBool("true")) // true
	fmt.Println(strconv.ParseBool("True")) // true
	fmt.Println(strconv.ParseBool("TRUE")) // true
	fmt.Println(strconv.ParseBool("TRue"))
	// false strconv.ParseBool: parsing "TRue": invalid syntax
	fmt.Println(strconv.ParseBool("0"))     // false
	fmt.Println(strconv.ParseBool("f"))     // false
	fmt.Println(strconv.ParseBool("F"))     // false
	fmt.Println(strconv.ParseBool("false")) // false
	fmt.Println(strconv.ParseBool("False")) // false
	fmt.Println(strconv.ParseBool("FALSE")) // false
	fmt.Println(strconv.ParseBool("FALse"))
	// false strconv.ParseBool: parsing "FAlse": invalid syntax
}

// Format xxx 转成string字符串
func useFormatBool() {
	fmt.Println(strconv.FormatBool(false))
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatFloat(3.1415, 'E', -1, 64))
}

func main() {
	fmt.Println("----------------ParseBool function-------------------------")
	useParseBool()

	fmt.Println("----------------FormatBool function-------------------------")
	useFormatBool()

	fmt.Println("----------------Atoi and Itoa function-------------------------")
	string2int()
}
