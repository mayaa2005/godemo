package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
)

func main() {
	var v Flags = FlagUp | FlagBroadcast | FlagLoopback
	fmt.Printf("%08b", v)
}
