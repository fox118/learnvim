package day09

import "testing"

func TestPassing(t *testing.T) {
	cases := []struct {
		score int
		want  bool
	}{
		{90, true},
		{70, true},
		{69, false},
	}

	for _, tc := range cases {
		if got := Passing(tc.score); got != tc.want {
			t.Fatalf("Passing(%d) = %v, want %v", tc.score, got, tc.want)
		}
	}
}
