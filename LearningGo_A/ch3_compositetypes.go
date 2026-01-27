package main

import (
	"fmt"
	"math"
)

func main() {
	var x [99]float64
	for i := 0; i < len(x); i++ {
		if i%2 == 0 {
			x[i] = math.Pow10(i)
		}
		fmt.Println("current value at index:", i, x[i])
	}

	var y = [5]int{0, 1, 2, 3, 4}
	fmt.Println(y, len(y), cap(y))

	// sparse array index must be int, value can be computed
	var sparse = [100]int{99, 32: int(x[4]), 100, 101, 65: 23, 24, 99: int(x[32])}
	fmt.Printf("sparse is an array of length:%d\n", len(sparse))

	var slice = []int{1, 2, 3, 4, 5}
	var empty []int
	fmt.Printf("slice is an array of length:%d\n and it contains: %v\n", len(slice), slice)
	if slice == nil {
		fmt.Printf("slice is nil\n")
	} else {
		fmt.Println("the slice 'slice' is not empty")
		fmt.Printf("empty == nil is true? %v\n", empty == nil)
	}
	empty = append(empty, 32)

	for range slice {
		empty = append(empty, 10)
		fmt.Printf("empty has length %d\n", len(empty))
		fmt.Printf("empty capacity is %d\n", cap(empty))
	}

	fmt.Println(empty, cap(empty))

	z := make([]int, 5)
	fmt.Printf("we maked a slice z of length:%d, cap: %d\n", len(z), cap(z))
	fmt.Println("here is z:", z)

	zeroLength := make([]int, 0, 10) // length is 0, capacity is 10
	fmt.Println("Zero length sliice with 10 cap:", zeroLength)
	zeroLength = append(zeroLength, 9)
	fmt.Println("Zero length with one item appended:", zeroLength)

}
