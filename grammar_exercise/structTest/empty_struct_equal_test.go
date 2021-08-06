package structTest

import (
	"fmt"
	"testing"
)

//Go 团队官方说了句
//Pointers to distinct zero-size variables may or may not be equal;
func TestEmptyStructEqual(t *testing.T) {
	a := new(struct{})
	b := new(struct{})

	fmt.Println(a, b, a == b)

	c := new(struct{})
	d := new(struct{})
	fmt.Println(c, d)
	fmt.Println(c, d, c == d)
}
