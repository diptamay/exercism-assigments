// Package armstrong determines whether a number is an Armstrong number
package armstrong

import "math"

// noOfDigits returns the number of digits in the number
func noOfDigits(num int) int {
	count := 0;
	for ; num > 0; {
		num = num / 10
		count++
	}
	return count
}

// IsNumber determines whether a number is an Armstrong number
func IsNumber(num int) bool {
	numDigits := float64(noOfDigits(num))
	n := num
	armstrong := 0.0;
	for ; n > 0; {
		remainder := float64(n % 10)
		armstrong = armstrong + math.Pow(remainder, numDigits)
		n = n / 10
	}
	return num == int(armstrong)
}
