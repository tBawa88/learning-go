package binarysearchtree

type Treenode struct {
	Data  int
	Left  *Treenode
	Right *Treenode
}

func New_Treenode(val int) *Treenode {
	return &Treenode{Data: val, Left: nil, Right: nil}
}
