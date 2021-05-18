package cash

import "testing"

func TestCash(t *testing.T) {
	tests := []struct {
		change   float64
		expected int
	}{
		{change: 0.41, expected: 4},
		{change: 0.01, expected: 1},
		{change: 0.15, expected: 2},
		{change: 1.6, expected: 7},
		{change: 23, expected: 92},
		{change: 4.2, expected: 18},
	}
	for _, tt := range tests {
		got := cash(tt.change)
		if got != tt.expected {
			t.Fatalf("test input: %f, Got: %d, expected: %d", tt.change, got, tt.expected)
		}
	}
}
