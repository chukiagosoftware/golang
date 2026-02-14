// My solution to Codesignal Go Algorithms exercise
package main

import (
	"fmt"
)

func main() {
	fmt.Println(areBracketsBalanced("(){}[]")) // Output: true
	fmt.Println(areBracketsBalanced("([)]"))   // Output: false
	fmt.Println(areBracketsBalanced("{"))      // Output: false
}

func areBracketsBalanced(inputStr string) bool {
	// TODO: iterate over all characters in the input

	lookup := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, r := range inputStr {
		if _, match := lookup[r]; match {
			stack = append(stack, r)
		} else if len(stack) < 1 || (lookup[stack[len(stack)-1]] != r) {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
	// TODO: use a slice as a stackto store opening brackets found
	// TODO: check if closing bracked has a matching pair in the stack
	// TODO: return true only if all prackets match
}
