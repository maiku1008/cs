package main

import (
	"fmt"
	"strings"
)

func main() {
	var height int
	for !(0 < height && height < 9) {
		fmt.Print("Height: ")
		fmt.Scan(&height)
	}
	for i := 1; i <= height; i++ {
		repeat := height - (height - i)
		hashes := strings.Repeat("#", repeat)
		fmt.Printf("%*s  %s\n", height, hashes, hashes)
	}
}
