package main

import (
	"fmt"
	"os"
)

func is_vowel(char rune) bool {
	return char == rune('a') || char == rune('e') || char == rune('u') || char == rune('i') || char == rune('o')
}

func main() {
	args := os.Args[1:]
	character := args[0][0]
	fmt.Println(is_vowel(rune(character)))
}
