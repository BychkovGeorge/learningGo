package main

import "fmt"

func fibonacci() func() int {
	previous := 0
	last := 0
	return func() int {
		result := previous + last
		if result == 0 {
			result = 1
		}
		previous = last
		last = result
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
