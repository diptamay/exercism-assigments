// Package raindrops contains function which convert a number to a string, the contents of which depend on the number's factors.
//
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
package raindrops

import "strconv"

// Convert converts a number to a string, the contents of which depend on the number's factors.
func Convert(num int) string {
	var ret = ""
	if num%3 == 0 {
		ret += "Pling"
	}
	if num%5 == 0 {
		ret += "Plang"
	}
	if num%7 == 0 {
		ret += "Plong"
	}

	//if nothing was converted then return itself
	if len(ret) == 0 {
		ret += strconv.Itoa(num)
	}

	return ret
}
