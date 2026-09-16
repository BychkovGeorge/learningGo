package main

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

func main() {
	slice := []int{1, 2, 3, -1, -2, -3}
	c := make(chan int, 2)
	go sum(slice[:len(slice)/2], c)
	go sum(slice[len(slice)/2:], c)
	x := <-c
	y := <-c
	fmt.Println(x, y, x+y)
}
