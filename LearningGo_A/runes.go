package main

import (
	"fmt"
	"unicode"
)
import "strconv"

func main() {
	s := 'r'
	var t rune = 'x' // alias for int32
	fmt.Println(s)   // prints 114

	fmt.Printf("%c\n", s) // prints the actual letter 'r' using the character format string
	fmt.Printf("%x\n", s) //prints 72 for some reason
	fmt.Println(t)        // prints the ascii/unicode value for 'x' which is 120

	var z string = "Hello, Sunshine"

	var bs []byte = []byte(z)
	var rs []rune = []rune(z)
	fmt.Println("Byte slice: ", bs) // prints each byte
	fmt.Println("Rune slice: ", rs) // identifies a byte that is an icon
	// bs = append(bs, 127774) error, cannot append because value overflows byte.
	rs = append(rs, 12774) // append the sun as a rune
	fmt.Println("We added the sun as rune value 12774: ", rs)
	newString := string(rs)
	fmt.Println(newString)

	i, err := strconv.Atoi("23")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i + 7) // prints 30
	var test = "joe scored 5 points, while adam scored 10 points and bob scored 2, with an extra 1 point scored by joe"
	fmt.Println(unicode.IsDigit(rune(test[12])))

}
