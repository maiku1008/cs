package main

import "testing"

func TestCaesar(t *testing.T) {
	tests := []struct {
		phrase   string
		key      int
		expected string
	}{
		{phrase: "a", key: 1, expected: "b"},
		{phrase: "barfoo", key: 23, expected: "yxocll"},
		{phrase: "BARFOO", key: 3, expected: "EDUIRR"},
		{phrase: "BaRFoo", key: 4, expected: "FeVJss"},
		{phrase: "barfoo", key: 65, expected: "onesbb"},
		{phrase: "world, say hello!", key: 12, expected: "iadxp, emk tqxxa!"},
	}
	for _, tt := range tests {
		got := caesar(tt.phrase, tt.key)
		if got != tt.expected {
			t.Fatalf("failed test %s: expected %s, got %s", tt.phrase, tt.expected, got)
		}
	}
}
