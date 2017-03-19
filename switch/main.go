package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println(fmt.Sprintf(`true%v`, "adb"))
	case false:
		fmt.Println(fmt.Sprintf(`false%v`, "adb"))
	}

	for i := 0; true; i++ {
		fmt.Println(i)
	}
}
