package main

import (
	"fmt"
	"strings"
	"unicode"
)

const chars = 26

func main() {

}

func substitution(phrase, key string) (string, error) {
	err := validateKey(key)
	if err != nil {
		return "", err
	}
	kr := []rune(key)
	var b strings.Builder
	var i int
	for _, c := range phrase {
		if unicode.IsLetter(c) {
			if unicode.IsLower(c) {
				i = int(c - 'a')
				c = unicode.ToLower(kr[i])
			}
			if unicode.IsUpper(c) {
				i = int(c - 'A')
				c = unicode.ToUpper(kr[i])
			}
		}
		b.WriteRune(c)
	}
	return b.String(), nil
}

func validateKey(key string) error {
	if len(key) != chars {
		return fmt.Errorf("key must have exactly %d characters, got: %d", chars, len(key))
	}
	store := make(map[rune]bool)
	for _, c := range key {
		if !unicode.IsLetter(c) {
			return fmt.Errorf("key must have alphabetical characters only, got: %c", c)
		}
		if store[c] {
			return fmt.Errorf("key has duplicate characters, got: %c", c)
		}
		store[c] = true
	}
	return nil
}
