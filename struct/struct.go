package main

import "fmt"

type Tree struct {
	Value       int
	Left, Right *Tree
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.Value = value
		return t
	}
	if value < t.Value {
		t.Left = Add(t.Left, value)
	} else {
		t.Right = Add(t.Right, value)
	}
	return t
}

func Sort(values []int) []int {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	return appendValues(make([]int, 0, len(values)), root)
	// return appendValues([]int{}, root)
	// return appendValues(values[:0], root)
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.Left)
		values = append(values, t.Value)
		values = appendValues(values, t.Right)
	}
	return values
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 1, 213, 12, 42, 45, 23, 45}
	fmt.Println(Sort(arr[:]))
}
