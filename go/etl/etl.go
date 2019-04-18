// Package etl extract some scrabble scores from a legacy system and transform
package etl

import "strings"

//DataIn Input type
type DataIn map[int][]string

//DataOut Output type
type DataOut map[string]int

// Transform some scrabble scores from a legacy system
func Transform(input DataIn) DataOut {
	out := make(map[string]int)
	for score, arr := range input {
		for _, char := range arr {
			out[strings.ToLower(char)] = score
		}
	}
	return out
}
