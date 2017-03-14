// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

type test struct {
	x string
}

func main() {
	a := test{x: tests("abc")}
	fmt.Println(a)

	//s := fmt.Sprintf("%.4f", 3.123456)
	s := fmt.Sprintf("%.4f", 3.1234)
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f)

	testh()
}

func tests(a string) string {
	return "def"
}

func testh() {
	fmts := fmt.Sprintf("%%.%vf", 2)
	s := fmt.Sprintf(fmts, 3.1234)
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f)
}
