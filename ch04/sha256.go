package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	x1 := sha256.Sum256([]byte("x"))
	x2 := sha256.Sum256([]byte("X"))
	fmt.Printf("x hash:%x\nX hash:%x\nequal:%t\n",
		x1, x2, x1 == x2)
}
