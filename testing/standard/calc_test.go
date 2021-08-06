package standard

import "testing"

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, 5) != 10 {
			t.Fatal("fail")
		}
	})
}
