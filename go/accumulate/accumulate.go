/*
Package accumulate contains functionality which given a collection and an operation to perform on each element of the collection,
returns a new collection containing the result of applying that operation to each element of the input collection.

	Given the collection of strings:
 		"cat", "Dog", "b4t", "gO"
 	And the operation:
		upcase a string
 	It would produce the collection of strings:
		"CAT", "DOG", "B4T, "GO"
*/
package accumulate

// Converter function type
type Converter func(string) string

// Accumulate given a collection and an operation to perform on each element of the collection, returns a new collection
// containing the result of applying that operation to each element of the input collection.
func Accumulate(given []string, converter Converter) []string {
	var ret []string
	for _, x := range given {
		ret = append(ret, converter(x))
	}
	return ret
}
