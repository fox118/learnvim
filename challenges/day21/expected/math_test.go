package day21

import "testing"

func TestDouble(t *testing.T) {
	if got := Double(4); got != 8 {
		t.Fatalf("Double(4) = %d", got)
	}
}
