package main

import (
	"fmt"
	binarysearchtree "tbawa/dsa_go/binary_search_tree"
)

func main() {
	root := binarysearchtree.New_Treenode(10)
	tree := binarysearchtree.BS_Tree{Root: root}

	find := 30
	if tree.Contains(find) {
		fmt.Println(find, "Value FOUND :)")
	} else {
		fmt.Println(find, "Value NOT FOUND :(")
	}

	tree.Insert(25)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(8)
	tree.Insert(15)
	tree.Insert(30)

	tree.PrintSideways()
	tree.Remove(99)
	fmt.Println("===== value removed ======")
	tree.PrintSideways()
}
