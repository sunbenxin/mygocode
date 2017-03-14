package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	a := []time.Time{time.Now(), time.Now(), time.Now()}
	sort.Slice(a, func(i, j int) bool { return a[i].After(a[j]) })
	fmt.Println(a)
}
