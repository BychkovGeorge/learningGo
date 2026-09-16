package main

import (
	"fmt"
	"math/rand"
)

const ITERATIONS_NUMBER = 10000000

func main() {
	calculate := func(change bool) {
		wins := 0
		for i := 0; i < ITERATIONS_NUMBER; i++ {
			var doors [3]bool
			var prizeIndex = rand.Intn(3)
			for j := 0; j < len(doors); j++ {
				if j == prizeIndex {
					doors[j] = true
				} else {
					doors[j] = false
				}
			}

			var playerIndex = rand.Intn(3)

			changeOption := func() {
				var montyIndex int
				for j := 0; j < len(doors); j++ {
					if j != playerIndex && !doors[j] {
						montyIndex = j
					}
				}
				for j := 0; j < len(doors); j++ {
					if j != playerIndex && j != montyIndex {
						if doors[j] {
							wins += 1
						}
					}
				}
			}

			doNotChangeOption := func() {
				if doors[playerIndex] {
					wins += 1
				}
			}

			if change {
				changeOption()
			} else {
				doNotChangeOption()
			}
		}
		var winPercent = (float64(wins) / ITERATIONS_NUMBER) * 100
		fmt.Printf("total wins percent - %v\n", winPercent)
	}

	calculate(true)
}
