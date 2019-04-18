// Package leap returns if a year is a leap year
package leap

// IsLeapYear returns true/false if a year is a leap year
func IsLeapYear(year int) bool {
	if year < 0 {
		return false
	}
	switch {
	case year%100 == 0 && year%400 == 0:
		return true
	case year%4 == 0 && year%100 != 0 && year%400 != 0:
		return true
	default:
		return false

	}

}
