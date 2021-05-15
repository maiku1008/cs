package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var number string
	fmt.Print("number: ")
	fmt.Scan(&number)

	fmt.Println(getCardType(number))
}

const (
	amex       = "AMEX"
	invalid    = "INVALID"
	mastercard = "MASTERCARD"
	visa       = "VISA"
)

func getCardType(input string) string {
	// Remove whitespace
	input = strings.ReplaceAll(input, " ", "")
	// Check length
	if len(input) < 2 {
		return invalid
	}
	var sum int
	even := len(input)%2 == 0

	for _, v := range input {
		// Convert string value to int.
		// If the value is not a number, return false.
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			return invalid
		}
		// Determine if we need to double our digit according to even bool
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		even = !even
		sum += digit
	}
	if sum%10 == 0 { // valid cc number
		if isMastercard(input) {
			return mastercard
		}
		if isVisa(input) {
			return visa
		}
		if isAmex(input) {
			return amex
		}
	}
	return invalid
}

func isMastercard(number string) bool {
	d := number[:2]
	return (d == "51" || d == "52" || d == "53" || d == "54" || d == "55") && len(number) == 16
}

func isVisa(number string) bool {
	d := number[:1]
	return d == "4" && (len(number) == 13 || len(number) == 16)
}

func isAmex(number string) bool {
	d := number[:2]
	return (d == "34" || d == "37") && len(number) == 15
}
