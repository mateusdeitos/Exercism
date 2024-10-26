package binarysearchtree

type BinarySearchTree struct {
	len   int
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i, len: 1}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data && bst.left == nil {
		bst.left = NewBst(i)
		bst.len++
		return
	}

	if i >= bst.data && bst.right == nil {
		bst.right = NewBst(i)
		bst.len++
		return
	}

	if bst.left != nil && i <= bst.data {
		bst.left.Insert(i)
		bst.len++
		return
	}

	if bst.right != nil && i >= bst.data {
		bst.right.Insert(i)
		bst.len++
		return
	}

	panic("Could not insert node")
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	s := []int{}

	if bst.left != nil {
		s = append(s, bst.left.SortedData()...)
	}

	s = append(s, bst.data)

	if bst.right != nil {
		s = append(s, bst.right.SortedData()...)
	}

	return s
}
