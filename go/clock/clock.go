// Package clock implements the Clock API
// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock
package clock

import "fmt"

// Clock type
type Clock struct {
	hour   int
	minute int
}

// New returns a new Clock
func New(hr, min int) Clock {
	tHr := min / 60
	hr = (hr + tHr) % 24
	tMin := min % 60
	if tMin < 0 {
		tMin = 60 + tMin
		hr--
	}
	if hr < 0 {
		hr = 24 + hr
	}
	return Clock{hour: hr, minute: tMin}
}

// String returns string representation
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes and returns a new Clock instance
func (c Clock) Add(minutes int) Clock {
	min := minutes + c.minute
	return New(c.hour, min)
}

// Subtract substracts minutes and returns a new Clock instance
func (c Clock) Subtract(minutes int) Clock {
	min := c.minute - minutes
	return New(c.hour, min)
}
