// Package letter calculate frequencies
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// ConcurrentFrequency counts the frequency in parallel of each rune in a list of
// text and returns this as a FreqMap.
func ConcurrentFrequency(sl []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap, len(sl))
	for _, s := range sl {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}
	for range sl {
		for k, v := range <-ch {
			m[k] += v
		}
	}
	return m
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
