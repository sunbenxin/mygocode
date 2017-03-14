package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{""}
	fmt.Println(len(a))

	i, err := strconv.Atoi("")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v\n", i)
}
