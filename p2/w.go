package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	Id     string
	name   string
	Gender string
}

func main() {
	a := test{
		Id:     "1",
		name:   "abc",
		Gender: "male",
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
