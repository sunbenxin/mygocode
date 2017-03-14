package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("a")
	fmt.Println(err)
}
