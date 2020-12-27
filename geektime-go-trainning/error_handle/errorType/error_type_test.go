package errorType

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	err := test()
	switch err := err.(type) {
	case nil:
		fmt.Println("successed")
	case *MyError:
		fmt.Printf("*MyError error happened:%v\n", err)
	case MyError:
		fmt.Printf("MyError error happened:%v\n", err)
	default:
		fmt.Println("unsupported type")
	}
}
