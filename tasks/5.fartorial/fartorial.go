package main

import (
	"fmt"
	"os"
	"strconv"
)

func get_factorial(x uint64) uint64 {
	var factorial uint64 = 1
	for i := uint64(1); i <= x; i++ {
		factorial *= i
	}
	return factorial
}

func main() {
	if number, err := strconv.ParseUint(os.Args[1], 10, 64); err == nil {
		fmt.Println(get_factorial(number))
	}

}
