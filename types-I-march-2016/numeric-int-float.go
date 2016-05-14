package main

import (
	"fmt"
)

func main() {
	var sum int

	sum = 4 + 6
	fmt.Println("4 + 6 =", sum)

	var fahr float32
	var cels float32
	fahr = 451

	cels = (fahr - 32) * 5 / 9
	fmt.Println(cels)
}
