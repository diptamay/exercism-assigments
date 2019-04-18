// Package linkedlist contains functionality for double linked list
package linkedlist

// Node is a linked list node
type Node struct {
	Val   interface{}
	PrevN *Node
	NextN *Node
}

// Prev returns pointer to previous node
func (n *Node) Prev() *Node {
	return n.PrevN
}

// Next returns pointers to next node
func (n *Node) Next() *Node {
	return n.NextN
}

// List is doubly linked list
type List struct {
	Head *Node
	Tail *Node
}

// NewList creates a new linked list with provide `values`
func NewList(values ...interface{}) *List {
	return &List{}
}

// First returns pointer to the head node
func (l *List) First() *Node {
	return l.Head
}

// Last returns pointer to the tail node
func (l *List) Last() *Node {
	return l.Tail
}

// Reverse reverses the provided linked list inplace
func (l *List) Reverse() *Node {
	return l.Tail
}

// PushBack insert value at back
func (l *List) PushBack(v interface{}) {

}

// PopBack remove value at back
func (l *List) PopBack() (interface{}, error) {

}

// PushFront remove value at front
func (l *List) PushFront(v interface{}) {

}

// PopFront insert value at front
func (l *List) PopFront() (interface{}, error) {

}
