// Package collatzconjecture returns the steps required to reach 1
// Take any positive integer n. If n is even, divide n by 2 to get n / 2.
// If n is odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
// The conjecture states that no matter which number you start with, you will always reach 1 eventually.
package collatzconjecture

import "errors"

// CollatzConjecture given a number n, return the number of steps required to reach 1.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("invalid number")
	}

	steps := 0
	stepper(&steps, n)
	return steps, nil
}

func stepper(steps *int, n int) {
	switch {
	case n == 1:
		return
	case n%2 == 0:
		*steps++
		stepper(steps, n/2)
	default:
		*steps++
		stepper(steps, 3*n+1)
	}

}
