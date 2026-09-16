package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(10000 * time.Millisecond)
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}
	for {
		select {
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(500 * time.Millisecond)
		}
	}
}
