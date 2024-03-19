package binarytree

import "testing"

func preorderTraversal(tree *Tree) []int {
	out := []int{}
	out = append(out, tree.root.Value)

	if tree.root.Left != nil {
		left := preorderTraversal(tree)
		out = append(out, left...)
	}

	if tree.root.Right != nil {
		right := preorderTraversal(tree)
		out = append(out, right...)
	}

	return out
}

func TestPreorderTraversal(t *testing.T) {

}
