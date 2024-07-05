package main

import (
	"fmt"
	"os"
	"strconv"
)

func two_sum(x, y float64) float64 {
	return (x + y)
}

func main() {
	args := os.Args[1:]
	if len(args) >= 2 {
		a, _ := strconv.ParseFloat(args[0], 64)
		b, _ := strconv.ParseFloat(args[1], 64)
		fmt.Println(two_sum(a, b))
	} else {
		fmt.Println("Not enought data")
	}
}
