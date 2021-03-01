package standard

import (
	"fmt"
	"testing"
)

func setup() {
	fmt.Println("setup")
}

func teardown() {
	fmt.Println("teardown")
}

func Test1(t *testing.T) {
	fmt.Println("I am Test1")
}
