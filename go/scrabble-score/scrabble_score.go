/*
Package scrabble has functionality that given a word, compute the scrabble score for that word.

Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10

*/
package scrabble

import "strings"

// Score given a word, computes the scrabble score for that word
func Score(input string) int {
	score := 0
	for _, char := range strings.ToUpper(input) {
		switch {
		case 'D' == char || 'G' == char:
			score += 2
		case 'B' == char || 'C' == char || 'M' == char || 'P' == char:
			score += 3
		case 'F' == char || 'H' == char || 'V' == char || 'W' == char || 'Y' == char:
			score += 4
		case 'K' == char:
			score += 5
		case 'J' == char || 'X' == char:
			score += 8
		case 'Q' == char || 'Z' == char:
			score += 10
		default:
			score++
		}
	}
	return score
}
