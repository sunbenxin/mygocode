package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

	var validDE = regexp.MustCompile(`\d{3}-\d{1,10}$`)
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(validDE.MatchString(invalid))
	fmt.Println(validDE.MatchString("12345a"))
}
