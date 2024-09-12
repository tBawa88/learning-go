package main

// A named struct S cannot contain a field with a type S. Since struct types are aggeregate types, they cannot contain themselves
// But it's possible to contain a pointer to the type *S

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head, tail *Node
	length     int
}

func (l *LinkedList) append(value int) {}
