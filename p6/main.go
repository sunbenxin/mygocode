package main

import (
	"fmt"
	"regexp"
	"unicode"
	"unicode/utf8"
)

var a = []string{
	"孙本新",
}

func main() {
	str := "中文文文文"
	var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$")
	fmt.Println(hzRegexp.MatchString(str))
	fmt.Println(IsChineseChar(str))
	fmt.Println(utf8.RuneCountInString(a[0]))

	var b = []rune("孙")
	if []rune(a[0])[0] == b[0] {
		fmt.Println(([]rune(a[0])[0]))
	}
}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
