package alg

import "testing"

func TestPow2(t *testing.T) {
	p := Pow2(2)
	if p!=4 {
		t.Fatalf("expected 4, but was %d.", p)
	}
}
