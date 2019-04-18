// Package isogram provide function to determine if a word or phrase is an isogram.
package isogram

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(str string) bool {
	track := map[rune]bool{}
	for _, ch := range str {
		if isAplha(ch) {
			//if uppercase - covert to lower case for tracking
			if ch <= 90 {
				ch += 32
			}
			exist := track[ch]
			if !exist {
				track[ch] = true
			} else {
				return false
			}
		}
	}
	return true
}

func isAplha(ch rune) bool {
	return (ch >= 97 && ch <= 122) || (ch >= 65 && ch <= 90)
}
