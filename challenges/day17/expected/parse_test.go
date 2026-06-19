package day17

import "testing"

func TestFirstRune(t *testing.T) {
	got, err := FirstRune("vim")
	if err != nil || got != 'v' {
		t.Fatalf("FirstRune(vim) = %q, %v", got, err)
	}

	if _, err := FirstRune(""); err == nil {
		t.Fatal("FirstRune(empty) expected error")
	}
}
