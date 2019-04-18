// Package dna contains functionality given a single stranded DNA string, compute how many times each nucleotide occurs in the string.
package dna

import "github.com/pkg/errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// ValidParts - maintains valid nucleotides
var ValidParts = map[rune]rune{
	'A': 'A',
	'C': 'C',
	'G': 'G',
	'T': 'T',
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, x := range d {
		_, ok := ValidParts[x]
		if ok {
			h[x] ++
		} else {
			return h, errors.New("invalid sequence in strand")
		}
	}
	return h, nil
}
