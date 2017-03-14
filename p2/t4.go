package main

import "fmt"

func main() {
	a:=2
	for {
		select a.(type)	{
		case int:
			fmt.Println("adf")
			break
		}
	}
}
