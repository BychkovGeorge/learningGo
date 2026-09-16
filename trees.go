package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func WalkRecursive(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func WalkStack(t *tree.Tree, ch chan int) {
	type node struct {
		visited bool
		tree    *tree.Tree
	}

	var stack []node
	current := node{
		tree: t,
	}

	for {
		if !current.visited && current.tree.Left != nil {
			newVal := node{
				visited: true,
				tree:    current.tree,
			}
			stack = append(stack, newVal)
			current = node{
				tree: current.tree.Left,
			}
			continue
		}
		ch <- current.tree.Value
		if current.tree.Right != nil {
			current = node{
				tree: current.tree.Right,
			}
			continue
		}
		if len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			break
		}
	}
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go WalkRecursive(t1, c1)
	go WalkStack(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)

	if Same(t1, t2) {
		fmt.Println("same")
	} else {
		fmt.Println("not same")
	}
}
