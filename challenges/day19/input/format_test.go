package day19

import "testing"

func TestLabel(t *testing.T) {
	got := Label(" vim ", 3)
	want := "vim:3"
	if got != want {
		t.Fatalf("Label() = %q, want %q", got, want)
	}
}
