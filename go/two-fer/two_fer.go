// Package twofer for returning string for - `Two-fer` or `2-fer` is short for two for one. One for you and one for me.
package twofer

// ShareWith returns strings depending of value of `name`
func ShareWith(name string) string {
	temp := "you"
	if name != "" {
		temp = name
	}
	return "One for " + temp + ", one for me."
}
