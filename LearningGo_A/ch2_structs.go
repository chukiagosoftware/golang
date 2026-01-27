package main

import "fmt"

func main() {
	type phone struct {
		name  string
		imei  int
		kind  string
		model int
		maker string
		year  int
	}

	mobiles := []phone{
		{"iphone", 12345678, "mini", 13, "apple", 2022},
		{"iphone", 56784321, "edge", 12, "apple", 2020},
		{"nexus", 45671234, "large", 5, "google", 2018},
		{"galaxy", 55558899, "s6", 6, "samsung", 2018},
	}

	for _, v := range mobiles {
		if v.imei > 30000000 {
			fmt.Println(v.name, v.kind, v.model, v.maker)
		} else {
			fmt.Println("IMEI blocked:", v.imei, v.kind, v.model, v.maker)
		}
	}

}
