package main

import "fmt"

//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
func greet(prefix string, who ...string) {
	fmt.Println(prefix, who)
}

//变长参数可以作为对应类型的slice进行二次传递
func F1(s ...string) {
	F2(s)
}

func F2(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

//使用空接口
func typecheck(values ...interface{}) {
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
	greet("hello:", "Joe", "Anna", "Eileen")
	F1("Joe", "Anna", "Eileen")

	fmt.Println("type check variable type")
	typecheck(true, "456", 10, 1.23)
}
