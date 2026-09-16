package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Has(v T) bool {
	if l == nil {
		return false
	}
	if l.val == v {
		return true
	}
	if l.next != nil {
		return l.next.Has(v)
	}
	return false
}

func main() {
	var list *List[int]
	for i := 0; i < 10; i++ {
		if i == 0 {
			element := List[int]{val: i, next: nil}
			list = &element
		} else {
			element := List[int]{val: i, next: list}
			list = &element
		}
	}

	fmt.Println("has - ", list.Has(0))
}
