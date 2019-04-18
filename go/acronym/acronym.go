//Package acronym consists of functions for returning string acronyms
package acronym

import (
	"strings"
)

// Abbreviate returns the acronym of a string
func Abbreviate(s string) string {
	var sb strings.Builder
	length := len(s)
	for i, char := range s {
		if i == 0 && char != ' ' && char != '-' {
			ch := s[i]
			sb.WriteString(string(ch))
		}
		if (i > 0 && i < length-1) && (char == ' ' || char == '-') {
			ch := s[i+1]
			if ch != ' ' && ch != '-' {
				sb.WriteString(string(ch))
			}
		}
	}
	return strings.ToUpper(sb.String())
}
