// Package lsproduct contains LargestSeriesProduct
package lsproduct

import (
	"errors"
)

// LargestSeriesProduct given a string of digits, calculate the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(str string, span int) (int, error) {
	length := len(str)
	if length < span {
		return -1, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	product := -1
	for i := 0; i <= length-span; i++ {
		temp := 1
		for j := 0; j < span; j++ {
			v := int(str[i+j] - '0')
			if v > 9 {
				return -1, errors.New("digits input must only contain digits")
			}
			temp = temp * v
		}
		if temp > product {
			product = temp
		}
	}
	return product, nil
}
