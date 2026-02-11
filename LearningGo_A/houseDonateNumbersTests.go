package main

import "testing"

func TestHouseGame(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{123, 234, 345, 456}, []int{362, 433, 144, 255}},
		{[]int{12, 34, 56}, []int{41, 63, 25}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{141, 4}, []int{44, 11}},
		{[]int{155, 261, 31}, []int{15, 156, 123}},
		{[]int{1, 1, 2, 2, 3, 3}, []int{3, 1, 1, 2, 2, 3}},
		{[]int{111, 222, 333}, []int{231, 312, 123}},
		{[]int{5, 5, 5, 5, 5}, []int{5, 5, 5, 5, 5}},
	}

	for _, test := range tests {
		got := HouseGame(test.input)
		want := test.output
		if !equalIntSlices(got, want) {
			t.Errorf("HouseGame(%v) = %v, want %v", test.input, got, want)
		}
	}
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
