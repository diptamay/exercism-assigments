// Package listops implements basic list operations.
// In functional languages list operations like length, map, and reduce are very common.
// Implements a series of basic list operations, without using existing functions.
// Functions Implemented:
// - Foldr
// - Foldl
// - Filter
// - Length
// - Map
// - Reverse
// - Append
// - Concat
package listops

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// IntList is list of ints
type IntList []int

// Foldr folds the values from the right
func (il IntList) Foldr(fn binFunc, initial int) int {
	prev := initial
	for n := il.Length() - 1; n >= 0; n-- {
		prev = fn(il[n], prev)
	}
	return prev
}

// Foldl folds the values from the left
func (il IntList) Foldl(fn binFunc, initial int) int {
	prev := initial
	for _, x := range il {
		prev = fn(prev, x)
	}
	return prev
}

// Filter returns a new list with values that satisfy the filter
func (il IntList) Filter(fn predFunc) IntList {
	ret := IntList([]int{})
	for _, x := range il {
		if fn(x) {
			ret = ret.AppendElem(x)
		}
	}
	return ret
}

// Length returns the list length
func (il IntList) Length() int {
	return len(il)
}

// Map returns a new list with passed Map function applied to the elements
func (il IntList) Map(fn unaryFunc) IntList {
	ret := IntList([]int{})
	for _, x := range il {
		ret = ret.AppendElem(fn(x))
	}
	return ret
}

// Reverse returns a new reversed list
func (il IntList) Reverse() IntList {
	ret := IntList([]int{})
	for n := il.Length() - 1; n >= 0; n-- {
		ret = ret.AppendElem(il[n])
	}
	return ret
}

// AppendElem appends the input element to itself
func (il IntList) AppendElem(in int) IntList {
	return append(il, in)
}

// Append appends the input list to itself
func (il IntList) Append(in IntList) IntList {
	return append(il, in...)
}

// Concat appends the array of lists to itself
func (il IntList) Concat(in []IntList) IntList {
	ret := il
	for _, lst := range in {
		ret = ret.Append(lst)
	}
	return ret
}
