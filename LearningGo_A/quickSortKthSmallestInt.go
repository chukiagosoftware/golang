// from codesignal Go Algorithms

package main

import (
	"fmt"
	"math"
)

// quicksort partition using last element, returns the pivotIndex position of this element
// in the final sorted array
func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1

	// Rearrange elements based on pivot
	for j := start; j < end; j++ {
		if arr[j] <= pivot {
			i++
			// Swap elements to place element <= pivot on the left
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Place pivot in its correct position
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1 // return pivot position
}

// partition function returns the final index of element, so we check
// if k = 1, the smallest element is at final index 0
// 0 - 0 == k - 1 == 1 - 1  for k == 1
// 1 - 0 == k - 1 == 2 - 1  for k == 2
func kthSmallest(arr []int, start, end, k int) int {
	if k > 0 && k <= end-start+1 {
		pos := partition(arr, start, end) // Partition the array
		// Check if pivot position is the k-1 index
		if pos-start == k-1 {
			return arr[pos] // Found the k-th smallest element
		}
		// Search left subarray
		if pos-start > k-1 {
			return kthSmallest(arr, start, pos-1, k)
		}
		// Search right subarray
		return kthSmallest(arr, pos+1, end, k-pos+start-1)
	}
	return math.MaxInt // Return max int if no valid k-th smallest found
}

func findKthSmallest(numbers []int, k int) (int, error) {
	if numbers == nil || len(numbers) < k {
		return math.MinInt, fmt.Errorf("invalid input")
	}
	return kthSmallest(numbers, 0, len(numbers)-1, k), nil
}
