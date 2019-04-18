// Package flatten provides function to take a nested list and return a single flattened list with all values except nil/null.
package flatten

// Flatten takes a nested list and return a single flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {
	ret := []interface{}{}

	if input != nil {
		switch value := input.(type) {
		case []interface{}:
			for _, el := range value {
				ret = append(ret, Flatten(el)...)
			}
		default:
			ret = append(ret, input)
		}
	}

	return ret
}
