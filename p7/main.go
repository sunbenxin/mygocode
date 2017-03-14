package main

import "fmt"

func main() {
	c := make(chan int, 8)
	bingo := false
	for { // send random sequence of bits to c
		select {
		//case c <- 0: // note: no statement, no fallthrough, no folding of cases
		case c <- 1:
			fmt.Println("in case")
		default:
			bingo = true
		}
		if bingo {
			break
		}
	}
	fmt.Println("saf")
}
