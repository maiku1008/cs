package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	key, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("plaintext: ")
	var phrase string
	fmt.Scanln(&phrase)
	fmt.Println("ciphertext:", caesar(phrase, key))
}

func caesar(phrase string, key int) string {
	var b strings.Builder
	var k = rune(key)
	for _, c := range phrase {
		if unicode.IsLetter(c) {
			if unicode.IsLower(c) {
				c = (c+k-'a')%26 + 'a'
			}
			if unicode.IsUpper(c) {
				c = (c+k-'A')%26 + 'A'
			}
		}
		b.WriteRune(c)
	}
	return b.String()
}
