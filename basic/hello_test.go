package basic

import "testing"

func TestHellow(t *testing.T) {
	res := Hello()
	exp := "Hello World"

	if res != exp {
		t.Errorf("got %q want %q", res, exp)
	}
}
