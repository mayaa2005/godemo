package main

import (
	"bytes"
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

	var text string
	text = "中国"
	u := []rune(text)
	fmt.Printf("%x\n", text)
	fmt.Printf("%x\n", u)

	fmt.Printf(intsToString([]int{2, 3, 5, 8, 13}))
}

func intsToString(a []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range a {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')

	return buf.String()
}
