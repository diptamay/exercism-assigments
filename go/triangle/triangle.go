// Package triangle contains functions to determine if a triangle is equilateral, isosceles, or scalene.
// An equilateral triangle has all three sides the same length.
// An isosceles triangle has at least two sides the same length.
// (It is sometimes specified as having exactly two sides the same length, but for the purposes of this exercise we'll say at least two.)
// A scalene triangle has all sides of different lengths.
package triangle

import "math"

// Kind is type of triangle
type Kind string

const (
	//NaT - not a triangle
	NaT = "NaT"

	//Equ - equilateral
	Equ = "Equ"

	//Iso -  isosceles
	Iso = "Iso"

	//Sca - scalene
	Sca = "Sca"
)

// KindFromSides determines if a triangle is equilateral, isosceles, or scalene.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		a <= 0 || b <= 0 || c <= 0 ||
		math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		(a+b) < c || (a+c) < b || (b+c) < a:
		k = NaT
	case a == b && a == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
