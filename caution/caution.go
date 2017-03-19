package main

import "fmt"

func main() {
	fmt.Println(example(1))
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}
