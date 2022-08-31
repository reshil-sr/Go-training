package basic

import "testing"

func TestProd(t *testing.T) {
	intMap := map[int]int{
		1: 2,
		2: 3,
		3: 5,
	}

	exp := 30

	if res := Prod(intMap); res != exp {
		t.Errorf("got %d want %d", res, exp)
	}
}
