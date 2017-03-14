package main

import "fmt"

type a struct {
	Id string
}

func (a a) Test() string {
	return a.Id
}

type b struct {
	a
}

func main() {
	b := b{a: a{Id: "abc"}}
	s := b.Test()
	fmt.Println(s)
}
