package main

import (
	"fmt"
	"leech/reversebits"
)

func main() {
	actual := reversebits.ReverseBits(43261596)
	// expect 964176192
	fmt.Println(actual)
}
