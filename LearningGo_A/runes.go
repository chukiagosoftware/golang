package main

import "fmt"

func main() {
	s := 'r'
	var t rune = 'x' // alias for int32
	fmt.Println(s)   // prints 114

	fmt.Printf("%c\n", s) // prints the actual letter 'r' using the character format string
	fmt.Printf("%x\n", s) //prints 72 for some reason
	fmt.Println(t)        // prints the ascii/unicode value for 'x' which is 120
}
