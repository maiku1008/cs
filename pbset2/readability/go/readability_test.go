package main

import "testing"

func TestReadability(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "sample 1",
			text:     "One fish. Two fish. Red fish. Blue fish.",
			expected: "Before Grade 1",
		},
		{
			name:     "sample 2",
			text:     "Would you like them here or there? I would not like them here or there. I would not like them anywhere.",
			expected: "Grade 2",
		},
		{
			name:     "sample 3",
			text:     "When he was nearly thirteen, my brother Jem got his arm badly broken at the elbow. When it healed, and Jem's fears of never being able to play football were assuaged, he was seldom self-conscious about his injury. His left arm was somewhat shorter than his right; when he stood or walked, the back of his hand was at right angles to his body, his thumb parallel to his thigh.",
			expected: "Grade 8",
		},
		{
			name:     "sample 4",
			text:     "It was a bright cold day in April, and the clocks were striking thirteen. Winston Smith, his chin nuzzled into his breast in an effort to escape the vile wind, slipped quickly through the glass doors of Victory Mansions, though not quickly enough to prevent a swirl of gritty dust from entering along with him.",
			expected: "Grade 10",
		},
		{
			name:     "sample 5",
			text:     "A large class of computational problems involve the determination of properties of graphs, digraphs, integers, arrays of integers, finite families of finite sets, boolean formulas and elements of other countable domains.",
			expected: "Grade 16+",
		},
	}
	for _, tt := range tests {
		got := readability(tt.text)
		if got != tt.expected {
			t.Fatalf("test %s: expected %s, got %s", tt.name, tt.expected, got)
		}
	}
}
