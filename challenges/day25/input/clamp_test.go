package day25

import "testing"

func TestClamp(t *testing.T) {
	cases := map[int]int{-2: 0, 4: 4, 12: 10}
	for in, want := range cases {
		if got := Clamp(in); got != want {
			t.Fatalf("Clamp(%d) = %d, want %d", in, got, want)
		}
	}
}
