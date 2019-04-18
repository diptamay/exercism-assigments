// Package isbn is used to validate book identification numbers
package isbn

// IsValidISBN validates book identification numbers
func IsValidISBN(isbn string) bool {
	const asciiStart = 48
	lenI := len(isbn)
	if lenI < 10 || lenI > 13 {
		return false
	}
	sum := 0
	factor := 10
	for i := 0; i < lenI; i++ {
		if (isbn[i]-asciiStart > 9 || isbn[i]-asciiStart < 0) && (i == lenI-1 && isbn[i] != 'X') && isbn[i] != '-' {
			return false
		}
		if i == lenI-1 && isbn[i] == 'X' {
			sum += 10
		} else if isbn[i] != '-' {
			sum += factor * int(isbn[i]-asciiStart)
			factor--
		}
	}
	//means there were either more or less than 10 digits without dashes et. al.
	if factor > 1 || factor < 0 || sum%11 != 0 {
		return false
	}
	return true
}
