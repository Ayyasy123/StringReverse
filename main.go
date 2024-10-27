package main

import (
	"fmt"
)

func reverseString(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func main() {
	input := "hello world"
	fmt.Println("Reversed:", reverseString(input))
}
