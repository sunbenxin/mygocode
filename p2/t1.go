package main

import "fmt"

type a struct {
	Id string
}

func (a a) Test() string {
	return a.Id
}

type b a

func main() {
	b := b{Id: "abc"}
	s := b.Test()
	fmt.Println(s)
}
