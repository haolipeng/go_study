package errorType

import "fmt"

var err error

type MyError struct {
	Msg  string
	Line string
	File string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%s:%s\n", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{
		File: "error_type.go",
		Line: "100",
		Msg:  "hello world",
	}
}
