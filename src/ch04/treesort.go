package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}
	if root.value > value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}

	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Sort(values []int) {
	var root *tree = nil
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func main() {
	a := [...]int{2, 9, 30, 5, 2, 88, 66, 22, 90, 100, 1}
	Sort(a[:])
	fmt.Print(a)
}
