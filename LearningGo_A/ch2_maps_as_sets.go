package main

import "fmt"

func main() {
	words := []string{"hey", "boo", "slice", "pizza", "bird", "plane", "day"}
	set := make(map[string]bool)
	for _, v := range words {
		set[v] = true
	}

	fmt.Println("Does the set contain skateboard?", set["skateboard"])
}
