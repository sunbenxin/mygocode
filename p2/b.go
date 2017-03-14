package main

import (
	"fmt"

	pg "gopkg.in/pg.v5"
)

func main() {
	db := pg.Connect(&pg.Options{
		User: "postgres",
		//Password: "",
		Database: "postgres",
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
