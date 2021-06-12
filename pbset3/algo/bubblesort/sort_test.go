package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "numbers",
			input:    []int{5, 4, 6, 1, 7, 3, 8, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		bubblesort(tt.input)
		if !cmp.Equal(tt.input, tt.expected) {
			t.Fatalf(cmp.Diff(tt.input, tt.expected))
		}
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "numbers",
			input:    []int{5, 4, 6, 1, 7, 3, 8, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		mergesort(tt.input)
		if !cmp.Equal(tt.input, tt.expected) {
			t.Fatalf(cmp.Diff(tt.input, tt.expected))
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "numbers",
			input:    []int{5, 4, 6, 1, 7, 3, 8, 2},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		selectionsort(tt.input)
		if !cmp.Equal(tt.input, tt.expected) {
			t.Fatalf(cmp.Diff(tt.input, tt.expected))
		}
	}
}

// add benchmarks
func BenchmarkBubblesort(b *testing.B) {
	input := []int{5, 4, 6, 1, 7, 3, 8, 2}
	for i := 0; i < b.N; i++ {
		bubblesort(input)
	}
}

func BenchmarkMergesort(b *testing.B) {
	input := []int{5, 4, 6, 1, 7, 3, 8, 2}
	for i := 0; i < b.N; i++ {
		mergesort(input)
	}
}

func BenchmarkSelectionsort(b *testing.B) {
	input := []int{5, 4, 6, 1, 7, 3, 8, 2}
	for i := 0; i < b.N; i++ {
		selectionsort(input)
	}
}
