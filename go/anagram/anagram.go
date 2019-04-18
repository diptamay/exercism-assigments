// Package anagram given a word and a list of possible anagrams, returns the correct sublist.
package anagram

import (
	"sort"
	"strings"
)

// Sorted sorts and returns a new string
func Sorted(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Detect given a word and a list of possible anagrams, returns the correct sublist.
func Detect(subject string, candidates []string) []string {
	ret := []string{}
	loweredSubject := strings.ToLower(subject)
	sortedSubject := Sorted(loweredSubject)
	lenSubject := len(sortedSubject)
	for _, candidate := range candidates {
		lenC := len(candidate)
		if lenC > 0 && lenSubject == lenC {
			loweredCandidate := strings.ToLower(candidate)
			if Sorted(loweredCandidate) == sortedSubject && loweredSubject != loweredCandidate {
				ret = append(ret, candidate)
			}
		}
	}
	return ret
}
