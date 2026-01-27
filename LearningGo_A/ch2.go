package main

import "fmt"

func main() {
	//i := 20
	//f := float32(i)
	//fmt.Println(i, f)

	const value = 20
	i := 20
	f := float32(i)
	fmt.Printf("i=%d\nf=%f\n", i, f) // prints decimal and zeroes for f
	fmt.Println(i, f)                // prints integer for both

	var (
		b      byte   = 255
		smallI int32  = 1<<31 - 1
		bigI   uint64 = 1<<64 - 1
	)
	b++
	smallI++
	bigI++
	fmt.Println(b, smallI, bigI)
	x := 32 + int(50.0)
	fmt.Printf("x is the sum of two literals: %d\n", x)

}
