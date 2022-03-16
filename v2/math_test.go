package math

import "testing"

func TestSum(t *testing.T) {
	t.Run("1 + 2", func(t *testing.T) {
		res := Sum(1, 2)
		if res != 3 {
			t.Fail()
		}
	})
}
