package html

// import "io"

// type NodeType int32

// type Attribute struct{ Key, Val string }

// type Node struct {
// 	Type                    NodeType
// 	Data                    string
// 	Attr                    []Attribute
// 	FirstChild, NextSibling *Node
// }

// // This is an enumeration, ErrorNode's value starts from 0, and automatically increments by one
// const (
// 	ErrorNode NodeType = iota
// 	TextNode
// 	DocumentNode
// 	ElementNode
// 	CommentNode
// 	DoctypeNode
// )

// func Parse(r io.Reader) (*Node, error)
