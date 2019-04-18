// Package bob returns bob's canned responses depending on whats said/asked bob
package bob

import (
	"strings"
)

// Hey returns bob's responses depending of whats asked/said
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	allUpper := IsAllUpper(remark)
	suffix := remark[len(remark)-1]

	switch {
	case suffix == '?' && allUpper:
		return "Calm down, I know what I'm doing!"
	case suffix == '?':
		return "Sure."
	case allUpper:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

// IsAllUpper checks if the whole strings is capitalized and contains atleast valid letters
func IsAllUpper(s string) bool {
	return strings.ToUpper(s) == s && strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
