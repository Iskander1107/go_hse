package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func is_prime(number int64) bool {
	sqrt_number := int64(math.Sqrt(float64(number)))
	for i := int64(2); i <= sqrt_number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	argument, _ := strconv.ParseInt(args[0], 10, 64)
	fmt.Println(is_prime(argument))
}
