package binarytree

import "testing"

type Node struct {
	Value       int
	Left, Right *Node
}

type Tree struct {
	root *Node
}

func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	root := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		root.Left = levelOrderBinaryTree(arr, left, size)
	}

	if right < size {
		root.Right = levelOrderBinaryTree(arr, right, size)
	}

	return root
}

func TestLevelOrderBinaryTree(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 17, 19}
	levelOrderBinaryTree(arr, 0, len(arr))
}
