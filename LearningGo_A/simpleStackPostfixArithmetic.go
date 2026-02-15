package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvaluateReversePolishNotation(expression string) int {
	// TODO: Initialize a slice to simulate a stack for holding integer values
	nums := []int{}

	// TODO: Split the expression into tokens using whitespace as the delimiter
	tokens := strings.Fields(expression)

	// TODO: Iterate over each token in the split expression
	for _, t := range tokens {
		num, err := strconv.Atoi(t)
		if err != nil {
			ln := len(nums)
			left := nums[ln-2]
			right := nums[ln-1]
			nums = nums[:ln-2]
			if t == "+" {
				sum := left + right
				nums = append(nums, sum)
			} else if t == "-" {
				diff := left - right
				nums = append(nums, diff)
			}
		} else {
			nums = append(nums, num)
		}
	}
	// TODO: If the token is an operator ('+' or '-'), pop the top two elements from the stack,
	// perform the corresponding operation, and push the result back onto the stack

	// TODO: If the token is an operand, parse it to an integer and push it onto the stack

	// TODO: Return the final result, which should be the only element left in the stack
	return nums[0] // Placeholder return statement
}

func main() {
	// The expression "1 2 + 4 -" is "(1 + 2) - 4"
	fmt.Println(EvaluateReversePolishNotation("1 2 + 4 -")) // Expected output: -1
}
