package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("total num:%d\n", total)
}

//使用空接口
func typeCheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println("int type", v)
		case string:
			fmt.Println("string type")
		case bool:
			fmt.Println("bool type")
		case float32:
			fmt.Println("float32 type")
		}
	}
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	//通过slice作为函数参数来传递
	s := []int{1, 2, 3, 4}
	sum(s...)

	fmt.Println("type check variable type")
	typeCheck(true, "456", 10, 1.23)
}
