package main

import (
	"fmt"
	"math"
)

func main() {
	var a = 1<<2 | 1<<5
	var b = 1<<3 | 1<<6
	fmt.Printf("a=%08b\n", a)
	fmt.Printf("b=%08b\n", b)

	var c = 12345678
	fmt.Printf("c=%d %[1]x %#[1]x %#[1]X\n", c)

	for i := 0; i < 8; i++ {
		fmt.Printf("x=%d,e^x=%8.3f\n", i, math.Exp(float64(i)))
	}

	x := 1 + 2i
	y := 2 + 3i
	fmt.Println(x * y)
}
