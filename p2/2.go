package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(5 * time.Second)
	idx := 0
	for {
		select {
		case <-tick:
			idx = 0
		default:
			if idx < 100000 {
				fmt.Println(idx)
				idx++
			}
		}
	}
}
