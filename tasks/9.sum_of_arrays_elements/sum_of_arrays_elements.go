package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum_of_arrays_elements(numbers []string) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		num, _ := strconv.Atoi(numbers[i])
		sum += num
	}
	return sum
}

func main() {
	args := os.Args[1:]
	fmt.Printf("%T", args)
	fmt.Println(sum_of_arrays_elements(args))
}
