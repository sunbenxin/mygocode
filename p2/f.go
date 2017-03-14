package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	in := `{"firstName":"John","lastName":"Dow"}`

	var p Person
	err := json.Unmarshal(in, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", p)
}
