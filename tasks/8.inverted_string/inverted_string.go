package main

import (
	"fmt"
	"os"
)

func invented_string(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	str := args[0]
	fmt.Println(invented_string(str))
}
