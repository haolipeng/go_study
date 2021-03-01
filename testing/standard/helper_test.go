package standard

import "testing"

type calcCases struct {
	A, B, Expected int
}

func testMulHelper(t *testing.T, cases *calcCases) {
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

	testMulHelper(t, &calcCases{
		A:        2,
		B:        0,
		Expected: 1,
	})
}
