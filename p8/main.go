package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(a["a"])
	fmt.Println(((-10) % 8))
	fmt.Println(math.Mod(float64(-10), float64(8)))
}

var a = map[string]bool{
	"abc": true,
	"def": true,
}
