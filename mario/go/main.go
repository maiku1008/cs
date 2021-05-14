package main

import (
	"fmt"
	"strings"
)

const (
	hash  = "#"
	space = " "
)

func main() {
	height := 5
	for i := 1; i <= height; i++ {
		hashes := strings.Repeat(hash, height-(height-i))
		fmt.Printf("%*s  %s\n", height, hashes, hashes)
	}
}
