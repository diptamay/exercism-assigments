// Package proverb contains function given a list of inputs, generate the relevant proverb
package proverb

import "strings"

// proverb statements
const (
	Single = "And all for the want of a %."
	Multi  = "For want of a %1 the %2 was lost."
)

// Proverb given a list of inputs, generate the relevant proverb
func Proverb(rhyme []string) []string {
	var ret []string
	lenR := len(rhyme)
	switch {
	case lenR == 1:
		ret = append(ret, strings.ReplaceAll(Single, "%", rhyme[0]))
	case lenR > 1:
		for i := range rhyme {
			if i != lenR-1 {
				temp := strings.ReplaceAll(Multi, "%1", rhyme[i])
				temp = strings.ReplaceAll(temp, "%2", rhyme[i+1])
				ret = append(ret, temp)
			} else {
				ret = append(ret, strings.ReplaceAll(Single, "%", rhyme[0]))
			}
		}
	}
	return ret
}
