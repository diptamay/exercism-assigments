// Package space contains function which given an age in seconds, calculate how old someone would be on:
//
// Earth: orbital period 365.25 Earth days, or 31557600 seconds
// Mercury: orbital period 0.2408467 Earth years
// Venus: orbital period 0.61519726 Earth years
// Mars: orbital period 1.8808158 Earth years
// Jupiter: orbital period 11.862615 Earth years
// Saturn: orbital period 29.447498 Earth years
// Uranus: orbital period 84.016846 Earth years
// Neptune: orbital period 164.79132 Earth years
//
// So if you were told someone were 1,000,000,000 seconds old, you should be able to say that they're 31.69 Earth-years old.
//
// If you're wondering why Pluto didn't make the cut, go watch this youtube video.
package space

import "math"

// Planet is name of planet
type Planet string

// Age given an age in seconds, calculate how old someone would be on different planets
func Age(sec float64, planet Planet) float64 {
	const EarthSec = 31557600
	var yearMap = map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	yearFactor, ok := yearMap[planet]
	if ok {
		ret := sec / (EarthSec * yearFactor)
		return math.Round(ret*100) / 100
	}
	return 0.0

}
