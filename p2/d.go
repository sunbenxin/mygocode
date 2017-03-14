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

	rawIn := json.RawMessage(in)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var p Person
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", p)
}
