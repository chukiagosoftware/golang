package main

import (
	"fmt"
)

func FindIndex(arr []int, x int) int {
	// TODO: find peak

	if len(arr) == 0 {
		return -1
	}

	peak := FindPeak(arr)

	// TODO: search to the left of the peak
	// TODO: search to the right of the peak
	if i := BinarySearch(arr, x, 0, peak, true); i != -1 {
		return i
	} else if i := BinarySearch(arr, x, peak+1, len(arr)-1, false); i != -1 {
		return i
	} else {
		return -1
	}
}

func FindPeak(arr []int) int {
	low, high := 0, len(arr)-1
	// TODO: implement this
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func BinarySearch(arr []int, x, low, high int, ascending bool) int {
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		} else if ascending {
			if arr[mid] < x {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			// TODO: implement the descending binary search logic
			if arr[mid] < x {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func main() {
	arr := []int{-3, -2, 4, 6, 10, 8, 7, 1}
	x := 7
	position := FindIndex(arr, x)
	if position == -1 {
		fmt.Println("Element Not Present")
	} else {
		fmt.Printf("Element Present at Index %d\n", position)
	}
}
