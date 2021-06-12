package main

func main() {}

// Runs bubblesort on input array
func bubblesort(arr []int) {
	length := len(arr)
	for i := range arr {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Runs MergeSort algorithm on a arr single
func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2
	return merge(mergesort(arr[:mid]), mergesort(arr[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func selectionsort(arr []int) {

}
