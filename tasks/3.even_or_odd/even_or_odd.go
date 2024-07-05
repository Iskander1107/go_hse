package main

import (
	"fmt"
	"os"
	"strconv"
)

func is_even(x int) bool {
	return x%2 == 0
}

func main() {
	args := os.Args[1:]
	number, _ := strconv.Atoi(args[0])
	if is_even(number) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
