package day17

import "fmt"

func FirstRune(s string) (rune, error) {
	for _, r := range s {
		return r, nil
	}
	return 0, nil
}
