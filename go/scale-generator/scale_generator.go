// Package scale contains function which given a tonic, or starting note, and a set of intervals, generate the musical scale
// starting with the tonic and following the specified interval pattern.
package scale

import "strings"

var chromaticFlats = [12]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
var chromaticSharps = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var increments = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

func indexOf(needle string, haystack []string) int {
	for i, v := range haystack {
		if needle == v {
			return i
		}
	}
	return -1
}

func useFlats(tonic string) bool {
	flatTonics := [5]string{"F", "d", "g", "c", "f"}
	return len(tonic) == 1 &&
		indexOf(tonic[:1], flatTonics[:]) != -1 ||
		strings.HasSuffix(tonic, "b")
}

// Scale given a tonic, or starting note, and a set of intervals, generate the musical scale starting with the tonic
// and following the specified interval pattern.
func Scale(tonic, interval string) []string {
	notes := &chromaticSharps
	if useFlats(tonic) {
		notes = &chromaticFlats
	}

	var scale []string
	initial := indexOf(strings.Title(tonic), notes[:])
	if interval == "" {
		for i := initial; i < initial+12; i++ {
			scale = append(scale, notes[i%12])
		}
	} else {
		i := initial
		for _, c := range interval {
			scale = append(scale, notes[i%12])
			i += increments[c]
		}
	}

	return scale
}
