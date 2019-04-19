// Package luhn functionality given a number determine whether or not it is valid per the Luhn formula.
package luhn

import (
	"strconv"
	"strings"
)

// Valid given a number determine whether or not it is valid per the Luhn formula.
func Valid(str string) bool {
	trimmed := strings.TrimSpace(str)
	length := len(trimmed)
	if length <= 1 {
		return false
	}

	//double every 2nd letter or subtract by 9 after doubling
	doubled := ""
	flip := false
	for i := length - 1; i >= 0; i-- {
		if trimmed[i] != 32 { // not space
			v := int(trimmed[i] - '0')
			if v < 0 || v > 9 {
				return false
			} else if flip {
				dbl := 2 * v
				if dbl > 9 {
					doubled = strconv.Itoa(dbl-9) + doubled
				} else {
					doubled = strconv.Itoa(dbl) + doubled
				}
				flip = false
			} else {
				doubled = strconv.Itoa(v) + doubled
				flip = true
			}
		}
	}
	//fmt.Printf("Testing=[%s], doubled =[%s]", str, doubled)

	//find the sum
	sum := 0
	for i := 0; i < len(doubled); i++ {
		v := int(doubled[i] - '0')
		sum = sum + v
	}
	//fmt.Printf(", sum=[%d]\n", sum)

	// is it finally a Luhn num
	if sum%10 != 0 {
		return false
	}
	return true
}
