package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum_of_numbers(x, y, z float64) float64 {
	return (x + y + z)
}

func main() {
	args := os.Args[1:]
	args1, _ := strconv.ParseFloat(args[0], 64)
	args2, _ := strconv.ParseFloat(args[1], 64)
	args3, _ := strconv.ParseFloat(args[2], 64)
	fmt.Println(sum_of_numbers(args1, args2, args3))
}
