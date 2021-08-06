package standard

import "testing"

type calcCases struct {
	A, B, Expected int
}

func testMulHelper(t *testing.T, cases *calcCases) {
	//有些帮助函数还可能在不同的函数中被调用，报错信息都在同一处，不方便问题定位。
	t.Helper() //不加此行会导致错误定位不准
	if val := Mul(cases.A, cases.B); val != cases.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			cases.A, cases.B, cases.Expected, val)
	}
}

func TestMulWithHelp(t *testing.T) {
	testMulHelper(t, &calcCases{
		A:        2,
		B:        3,
		Expected: 6,
	})

	testMulHelper(t, &calcCases{
		A:        2,
		B:        -3,
		Expected: -6,
	})

	testMulHelper(t, &calcCases{ //加上t.Helper()后定位到这里
		A:        2,
		B:        0,
		Expected: 1,
	})
}
