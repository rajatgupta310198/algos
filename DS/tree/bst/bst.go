package bst

import "fmt"

type T struct {
	root *node
	count int
}
type node struct {
	item int
	left, right *node
}

func insert(root *node, item int) ( newNode *node, added bool)  {
	if newNode == root {
		newNode = &node{item:item}
		added = true
	} else if item < root.item {
		root.left, added = insert(root.left, item)
	} else if item > root.item {
		root.right, added = insert(root.right, item)
	}
	return 
}

func (t * T) Insert(item int)  (added bool)  {
	t.root, added = insert(t.root, item)
	return added
}

func inOrder(root *node) {
	if root !=nil {
		inOrder(root.left)
		fmt.Print(root.item)
		inOrder(root.right)
	}
}