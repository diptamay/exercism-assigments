// Package grains has functionality for the number of grains of wheat on a chessboard
package grains

import (
	"errors"
	"math"
)

// Square calculates the number of grains of wheat on a specific chessboard square given that the number on each square doubles.
func Square(box int) (uint64, error) {
	if box < 1 || box > 64 {
		return 0, errors.New("Invalid box number")
	}
	wheat := uint64(math.Pow(2.0, float64(box)-1.0))
	return wheat, nil
}

// Total calculates the total number of grains of wheat on a chessboard given that the number on each square doubles.
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		total += val
	}
	return total
}
