package day18

import "testing"

func TestSameName(t *testing.T) {
	if !SameName(" Vim ", "vim") {
		t.Fatal("expected names to match")
	}
}
