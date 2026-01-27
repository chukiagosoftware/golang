package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {

	//var nilMap map[string]int
	//totalWins := map[string]int{} //empty but not nil
	//totalWins["Orcas"] = 1
	//totalWins["Lions"] = 2
	//fmt.Println(totalWins["Orcas"])
	//fmt.Println(totalWins["Kittens"])
	//fmt.Println("nilMap: ", nilMap)

	m := map[string][]int{
		"yo":     {0, 1, 2},
		"wassup": {3, 4, 5},
	}
	n := make(map[string][]int, 4)
	n["yo"] = []int{0, 1, 2}
	n["wassup"] = []int{3, 4, 8}
	// fmt.Println("Does n equal m?", maps.Equal(m, n)) works for int values not []int values, see below

	// tell it to make them equal based on keys
	fmt.Println("Does n keys and value equal m keys and values?", maps.EqualFunc(n, m, slices.Equal[[]int]))
}
