package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rectangle struct {
	width  float64
	height float64
}

func (d Rectangle) Get_area() float64 {
	return d.width * d.height
}
func main() {
	args := os.Args[1:]
	height, _ := strconv.Atoi(args[0])
	width, _ := strconv.Atoi(args[1])
	Rec := Rectangle{width: float64(width), height: float64(height)}
	fmt.Println(Rec.Get_area())
}
