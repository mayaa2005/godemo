package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Printf("USD:%s,EUR:%s,GBP:%s,RMB:%s",
		symbol[USD], symbol[EUR], symbol[GBP], symbol[RMB])
}
