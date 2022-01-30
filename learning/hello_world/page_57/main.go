package main

import (
	"fmt"
)

func main() {

	var length float64 = 1.2
	var width int = 2

	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width ?", length > float64(width))

	length = float64(width)
	fmt.Println(length)

	length = 3.75
	width = 5

	width = int(length)
	fmt.Println(width)
}
