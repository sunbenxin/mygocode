package main

import "fmt"

var (
	a = 1
)

func init() {
	fmt.Println(a)
}

func main() {
	a := 5
	fmt.Println(a)
}
