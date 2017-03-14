package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "abadefghi"

	longest := 0
	start := 0
	end := 0
	for i := 1; i <= len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if !hasSameNew(s[i-1 : j]) {
				if j-i+1 > longest {
					longest = j - i + 1
					start = i - 1
					end = j
				}
			}
		}
	}
	fmt.Println(longest)
	fmt.Println(s[start:end])
}

func hasSame(s string) bool {
	d := append([]byte{}, []byte(s)...)
	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	for i := 0; i < len(d); i++ {
		if i > 0 && d[i] == d[i-1] {
			return true
		}
	}

	return false
}

func hasSameNew(s string) bool {
	m := map[rune]bool{}
	for _, v := range s {
		if m[v] {
			return true
		}
		m[v] = true
	}
	return false
}
