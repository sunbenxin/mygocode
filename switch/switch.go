package switchment

//https://github.com/golang/go/wiki/Switch

import (
	"fmt"
	"strings"
)

func switchA() {
	switch c {
	case '&':
		esc = "&amp;"
	case '\'':
		esc = "&apos;"
	case '<':
		esc = "&lt;"
	case '>':
		esc = "&gt;"
	case '"':
		esc = "&quot;"
	default:
		panic("unrecognized escape character")
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func anotherSwitch() {
	argv := strings.Fields(command)
	switch argv[0] {
	case "echo":
		fmt.Print(argv[1:]...)
	case "cat":
		if len(argv) <= 1 {
			fmt.Println("Usage: cat <filename>")
			break
		}
		PrintFile(argv[1])
	default:
		fmt.Println("Unknown command; try 'echo' or 'cat'")
	}
}

func switch2() {
	// Unpack 4 bytes into uint32 to repack into base 85 5-byte.
	var v uint32
	switch len(src) {
	default:
		v |= uint32(src[3])
		fallthrough
	case 3:
		v |= uint32(src[2]) << 8
		fallthrough
	case 2:
		v |= uint32(src[1]) << 16
		fallthrough
	case 1:
		v |= uint32(src[0]) << 24
	}
}

func switch3() {
	//The 'fallthrough' must be the last thing in the case; you can't write something like

	switch {
	case f():
		if g() {
			fallthrough // Does not work!
		}
		h()
	default:
		error()
	}

	//However, you can work around this by using a 'labeled' fallthrough:
	switch {
	case f():
		if g() {
			goto nextCase // Works now!
		}
		h()
		break
	nextCase:
		fallthrough
	default:
		error()
	}
}

func letterOp(code int) bool {
	switch chars[code].category {
	case "Lu", "Ll", "Lt", "Lm", "Lo":
		return true
	}
	return false
}
