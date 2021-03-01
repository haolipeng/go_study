//https://mp.weixin.qq.com/s?src=11&timestamp=1614216678&ver=2911&signature=jlknThXda2EEigDB-YExrs0JDF1-YuA1T8cs*Zd10Vd5FWA-tTDNQn1rTPUHi-HWSrAb35znfRG7eoVPiK4Rguz1nWl7dQpdfgTzRcT7ivRMDOTK2cNthA0eLrGxYosL&new=1
package standard

import (
	"fmt"
	"testing"
)

func Fib(n int) int {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)

	actual := Fib(in)

	fmt.Println("value is:", actual)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}

//table-driven tests
func TestCaseTable(t *testing.T) {
	var fibTests = []struct {
		in       int //输出参数
		expected int //期望的结果
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 4},
		{6, 8},
		{7, 13},
	}

	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("Fib(%d) = %d; expected %d", tt.in, actual, tt.expected)
		}
	}
}
