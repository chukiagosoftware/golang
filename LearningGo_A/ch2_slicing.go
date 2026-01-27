package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s is a slice of 5 ints  and the first two items are:", s[:2])
	fmt.Println("s is a slice of 5 ints and the last two items are:", s[3:])

	fmt.Println("s before modifying derived slices:", s)
	x := s[:3]
	y := s[2:5]
	x[0] = 100
	y[0] = 200
	fmt.Println("s after modifying x[0] which is s[0] and y[0] which is s[2]:", s)

	// when subslicing, the new slice capacity is equal to the original slice minus the starting position of the subslice
	t := make([]string, 0, 5)
	t = append(t, "a", "b", "c", "d")
	fmt.Println("t: ", t)
	u := t[:2]
	fmt.Println("u: ", u)
	v := t[2:]
	fmt.Println("v: ", v)
	u = append(u, "i", "j", "k")
	fmt.Println("u after appending ijk: ", u)
	fmt.Println("t after appending ijk to u: ", t)
	fmt.Println("v after appending ijk to u: ", v)
	t = append(t, "x")
	fmt.Println("t after appending x to t: ", t)
	fmt.Println("u after appending x to t: ", u)
	fmt.Println("v after appending x to t: ", v)
	v = append(v, "y")
	fmt.Println("v after appending y to v: ", v)

	// to avoid this weird behavior with subslice extra capacity overwriting the original slice use full subslice
	p := make([]string, 0, 5)
	p = append(p, "a", "b", "c", "d")
	q := p[:2:2]  // no extra capacity is copied to new subslice
	r := p[2:4:4] // again, no extra cap is copied. new sublice cap = 4 -2 = 2
	q = append(q, "i", "j", "k")
	fmt.Println("p unmodified after appending ijk to q:", p)
	fmt.Println("q: ", q)
	r = append(r, "r")
	fmt.Println("r after appending r to r:", r)
	fmt.Println("p unmodified after appending r to r:", p)

}
