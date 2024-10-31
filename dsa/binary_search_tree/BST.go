package binarysearchtree

type BS_Tree struct {
	Root  *Treenode
	tsize int
}

// contains(int) : searches the tree for a value, returns true if the value is found, false if it is not
func (tree *BS_Tree) contains(val int) bool {
	return bst_search(tree.Root, val)
}

// insert(int) : inserts the value in the tree. Duplicate values are ignored
func (tree *BS_Tree) insert(val int) {
	if tree.Root == nil {
		tree.Root = New_Treenode(val)
		tree.tsize++
		return
	}

	insert_node(tree.Root, val)
	tree.tsize++
}

// remove(int) : finds the value in the tree, then removes it
// if the value is not found, it returns an error
func (tree *BS_Tree) remove(val int) error {
	err := remove_node(&tree.Root, val)
	if err != nil {
		return err
	}
	tree.tsize--
	return nil
}
