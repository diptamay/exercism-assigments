// Package reverse - reverses strings
package reverse

import (
	"strings"
)

// String reverses a string
func String(str string) string {
	if len(str) == 0 {
		return str
	}

	var buffer strings.Builder
	runed := []rune(str)
	for i := len(runed) - 1; i >= 0; i-- {
		buffer.WriteRune(runed[i])
	}
	return buffer.String()
}
