package day11

import "testing"

func TestDisplayName(t *testing.T) {
	p := Profile{Username: "fox", Email: "fox@example.com"}
	if got := DisplayName(p); got != "fox" {
		t.Fatalf("DisplayName() = %q", got)
	}
}
