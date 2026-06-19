package day08

import "testing"

func TestWelcome(t *testing.T) {
	got := Welcome("Vim")
	want := "Hello, Vim!"
	if got != want {
		t.Fatalf("Welcome() = %q, want %q", got, want)
	}
}
