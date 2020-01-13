package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	_, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline error:%s", err)
		os.Exit(1)
	}
}
