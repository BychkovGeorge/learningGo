package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (slice [][]uint8) {
	slice = make([][]uint8, dy)
	for i := range slice {
		subSlice := make([]uint8, dx)
		slice[i] = subSlice
		for j := range subSlice {
			slice[i][j] = uint8(i * j)
		}
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
