package main

import (
	"fmt"
	binarysearchtree "tbawa/dsa_go/binary_search_tree"
	prioqueue "tbawa/dsa_go/priority_queue"
)

func main2() {
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

func main() {
	queue := prioqueue.Queue{}
	queue.Enqueue("Go to gym", 3)
	queue.Enqueue("Enroll into BCA degree", 1)
	queue.Enqueue("Get good at maths", 2)
	queue.Enqueue("Get good at core CS concepts", 2)
	queue.Enqueue("Regrow your hair", 4)

	queue.Print()

}
