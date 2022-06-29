package main

import "fmt"

type Tree struct {
	name     string
	children map[string]Tree
}

func output(tree *Tree, depth int) {
	for i := 0; i < depth-1; i++ {
		fmt.Print(" ")
	}
	if tree.name != "" {
		fmt.Print(tree.name, '\n')
	}
	depth++
	for _, v := range tree.children {
		output(&v, depth)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
}
