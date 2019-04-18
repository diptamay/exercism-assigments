//Package hamming of functions for calculating Hamming Distance
package hamming

import "errors"

//Distance returns the Hamming Distance or throws an error for unequal sequences
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("hamming distance is only defined for sequences of equal length")
	}
	mismatch := 0
	for i := range a {
		if a[i] != b[i] {
			mismatch++
		}
	}
	return mismatch, nil
}
