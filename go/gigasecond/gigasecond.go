// Package gigasecond returns the date time when someone has lived for 10^9 seconds given a person's birthdate
// A gigasecond is 10^9 (1,000,000,000) seconds.
package gigasecond

// import path for the time package from the standard library
import "time"

//AddGigasecond returns the date time when someone has lived for 10^9 seconds given a person's birthdate/time
func AddGigasecond(t time.Time) time.Time {
	const giga = 1000000000
	t = t.Add(time.Second * time.Duration(giga))
	return t
}
