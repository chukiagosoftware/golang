package main

import "fmt"
import "time"

func main() {
	/* Declare a variable locally with implied type or use var to specify a type
	variables with no value set will use the types zero value. */
	// s := "Bob"
	//var j string
	//j = "James"
	//var i int32 = 2
	//var i int64 = 1
	// var i uint8 = 1 //  goes to 0 after overflow when multiplying by 2
	// adding 1 makes infinite loop over u or signed range
	var i uint8 = 1
	fmt.Printf("i=%d\n", i)
	for {
		if i < 0 {
			fmt.Println("negative number")
			break
		}

		if i == 0 {
			fmt.Println("zero number")
			break
		}
		//i = i * 2
		i = i * 2 // reaches zero due to multiply 2^31 by 2 == 2^32 which overflows exactly and 0*0 stays 0
		fmt.Printf("%d\n", i)
		time.Sleep(5000)
	}
	//fmt.Printf("Hello, %s, and hey %s\n", j, s)
}
