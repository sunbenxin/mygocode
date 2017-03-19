package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	rep, err := conn.Do("PING")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rep)
}
