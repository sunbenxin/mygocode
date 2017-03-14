package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int, 100)
	for i := 0; i < 1; i++ {
		go func(idx int) {
			for range dataChan {
				fmt.Println(idx)
				time.Sleep(10)
			}
		}(i)
	}

	for i := 0; i < 100000; i++ {
		dataChan <- 1
		fmt.Println(len(dataChan))
	}

}
