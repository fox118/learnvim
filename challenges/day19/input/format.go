package day19

import (
	"math"
	"strings"
)

func Label(name string, count int) string {
	return fmt.Sprintf("%s:%d", strings.TrimSpace(name), count)
}
