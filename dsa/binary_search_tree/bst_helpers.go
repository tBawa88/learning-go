package binarysearchtree

import "fmt"

func bst_search(node *Treenode, val int) bool {
	if node == nil {
		return false
	} else if node.Data == val {
		return true
	} else if val < node.Data {
		return bst_search(node.Left, val)
	} else {
		return bst_search(node.Right, val)
	}
}

func insert_node(node *Treenode, val int) {
	if val < node.Data {
		if node.Left == nil {
			node.Left = New_Treenode(val)
		} else {
			insert_node(node.Left, val)
		}
	} else if val > node.Data {
		if node.Right == nil {
			node.Right = New_Treenode(val)
		} else {
			insert_node(node.Right, val)
		}
	} else {
		//do nothing
	}
}

// using double pointer to directly update the value of a node (since we don't have references in Go)
func remove_node(node **Treenode, val int) error {
	if *(node) == nil {
		return fmt.Errorf("Tree empty OR Node not found")
	} else if val < (*node).Data {
		return remove_node(&(*node).Left, val)
	} else if val > (*node).Data {
		return remove_node(&(*node).Right, val)
	} else { // val == node.Data
		// check the possible cases for the current node
		switch {
		case (*node).Left == nil && (*node).Right == nil:
			*node = nil
			return nil
		case (*node).Left == nil && (*node).Right != nil:
			*node = (*node).Right
			return nil
		case (*node).Right == nil && (*node).Left != nil:
			*node = (*node).Left
			return nil
		default: // when the node has both left and right subtree
			minFromRight := findMin((*node).Right)
			(*node).Data = minFromRight
			return remove_node(&(*node).Right, minFromRight)
		}
	}
}

func findMin(node *Treenode) int {
	if node.Left == nil {
		return node.Data
	} else {
		return findMin(node.Left)
	}
}

func printSideways(node *Treenode, indent string) {
	if node != nil {
		printSideways(node.Right, indent+"   ")
		fmt.Println(indent, node.Data)
		printSideways(node.Left, indent+"   ")
	}
}

func printInorder(node *Treenode) {
	if node != nil {
		printInorder(node.Left)
		fmt.Printf("%d, ", node.Data)
		printInorder(node.Right)
	}
}
func printPre(node *Treenode) {
	if node != nil {
		fmt.Printf("%d, ", node.Data)
		printInorder(node.Left)
		printInorder(node.Right)
	}
}
func printPost(node *Treenode) {
	if node != nil {
		printInorder(node.Left)
		printInorder(node.Right)
		fmt.Printf("%d, ", node.Data)
	}
}
