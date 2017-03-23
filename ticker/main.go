package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)

	i := 0
	for range tick {
		go func(i int) {
			fmt.Println(i)
		}(i)
		i++
	}
}
