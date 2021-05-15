package main

import "testing"

func TestGetCardType(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected string
	}{
		{
			name:     "valid VISA",
			number:   "4003600000000014",
			expected: "VISA",
		},
		{
			name:     "valid Mastercard",
			number:   "5555555555554444",
			expected: "MASTERCARD",
		},
		{
			name:     "valid American Express",
			number:   "378282246310005",
			expected: "AMEX",
		},
		{
			name:     "invalid number",
			number:   "1",
			expected: "INVALID",
		},
		{
			name:     "invalid American Express",
			number:   "369421438430814",
			expected: "INVALID",
		},
		{
			name:     "invalid VISA",
			number:   "4062901840",
			expected: "INVALID",
		},
		{
			name:     "invalid Mastercard",
			number:   "5673598276138003",
			expected: "INVALID",
		},
	}
	for _, tt := range tests {
		got := getCardType(tt.number)
		if got != tt.expected {
			t.Fatalf("%s: \n\tExpected: %s\n\t Got: %s", tt.name, tt.expected, got)
		}
	}
}
