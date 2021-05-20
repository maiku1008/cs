package main

import "testing"

func TestSubstitution(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		phrase      string
		expected    string
		errExpected bool
	}{
		{
			name:        "valid 1",
			key:         "ZYXWVUTSRQPONMLKJIHGFEDCBA",
			phrase:      "A",
			expected:    "Z",
			errExpected: false,
		},
		{
			name:        "valid 2",
			key:         "ZYXWVUTSRQPONMLKJIHGFEDCBA",
			phrase:      "a",
			expected:    "z",
			errExpected: false,
		},
		{
			name:        "valid 3",
			key:         "NJQSUYBRXMOPFTHZVAWCGILKED",
			phrase:      "ABC",
			expected:    "NJQ",
			errExpected: false,
		},
		{
			name:        "valid 4",
			key:         "NJQSUYBRXMOPFTHZVAWCGILKED",
			phrase:      "XyZ",
			expected:    "KeD",
			errExpected: false,
		},
		{
			name:        "valid 5",
			key:         "NJQSUYBRXMOPFTHZVAWCGILKED",
			phrase:      "XyZ",
			expected:    "KeD",
			errExpected: false,
		},
		{
			name:        "valid 6",
			key:         "YUKFRNLBAVMWZTEOGXHCIPJSQD",
			phrase:      "This is CS50",
			expected:    "Cbah ah KH50",
			errExpected: false,
		},
		{
			name:        "valid 7",
			key:         "yukfrnlbavmwzteogxhcipjsqd",
			phrase:      "This is CS50",
			expected:    "Cbah ah KH50",
			errExpected: false,
		},
		{
			name:        "valid 8",
			key:         "YUKFRNLBAVMWZteogxhcipjsqd",
			phrase:      "This is CS50",
			expected:    "Cbah ah KH50",
			errExpected: false,
		},
		{
			name:        "wrong length",
			key:         "CIAO",
			errExpected: true,
		},
		{
			name:        "invalid chars",
			key:         "!@#$%YBRXMOPFTHZVAWCGILKED",
			errExpected: true,
		},
		{
			name:        "duplicate character",
			key:         "NNQSUYBRXMOPFTHZVAWCGILKED",
			errExpected: true,
		},
		{
			phrase:      "multiple duplicates",
			key:         "NNNNNNNNNNNNNNNNNNNNNNNNNN",
			errExpected: true,
		},
	}
	for _, tt := range tests {
		got, err := substitution(tt.phrase, tt.key)
		errReceived := err != nil
		if tt.errExpected != errReceived {
			t.Fatalf("name: %s, unexpected error status, got: %v", tt.name, err)
		}
		if got != tt.expected {
			t.Fatalf("name: %s, expected %s, got %s", tt.name, tt.expected, got)
		}
	}
}
