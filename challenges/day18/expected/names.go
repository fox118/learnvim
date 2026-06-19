package day18

import "strings"

func SameName(a, b string) bool {
	return cleanName(a) == cleanName(b)
}

func cleanName(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
