package day27

import "testing"

func TestCountCommands(t *testing.T) {
	if got := CountCommands([]string{"w", "dd", "p"}); got != 3 {
		t.Fatalf("CountCommands() = %d", got)
	}
}
