package day10

import "strings"

func Slug(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}
