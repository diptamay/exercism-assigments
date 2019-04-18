// Package flatten provides function to take a nested list and return a single flattened list with all values except nil/null.
package flatten

import "fmt"

func Flatten(input ...interface{}) []interface{} {
	var ret []interface{}
	if input != nil {
		ln := len(input)
		fmt.Printf("(%v, %T) - len=%v\n", input, input, ln)
		if ln > 1 {
			for _, el := range input {
				ret = append(ret, Flatten(el))
			}
		} else {
			ret = append(ret, input)
		}

	}

	return ret
}
