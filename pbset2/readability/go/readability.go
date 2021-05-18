package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {

}

func readability(text string) string {
	var words float64 = 1
	var letters, sentences float64
	for _, c := range text {
		c = unicode.ToLower(c)
		if unicode.IsLetter(c) {
			letters++
		}
		if unicode.IsSpace(c) {
			words++
		}
		if c == '.' || c == '!' || c == '?' {
			sentences++
		}
	}
	index := math.Round(0.0588*((letters/words)*100) - 0.296*((sentences/words)*100) - 15.8)

	switch {
	case index < 1:
		return "Before Grade 1"
	case index > 16:
		return "Grade 16+"
	default:
		return fmt.Sprintf("Grade %d", int(index))
	}
}
