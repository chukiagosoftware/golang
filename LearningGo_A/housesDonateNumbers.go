package main

import "strconv"
import "fmt"

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func HouseGame(houses []int) []int {
	// TODO: implement your solution here

	n := len(houses)
	step := 1
	var houseStrings []string

	for _, house := range houses {
		hString := strconv.Itoa(house)
		houseStrings = append(houseStrings, hString)
	}
	fmt.Println("House strings: ", houseStrings)

	for {

		donations := []string{}
		houseStringsCopy := []string{}

		for _, house := range houseStrings {
			var houseLen int = len(house)
			fmt.Printf("House length: %d, step: %d\n", houseLen, step)
			if step > houseLen {
				fmt.Println("skipping this house: ", house)
				donations = append(donations, "") // No digit to donate
				houseStringsCopy = append(houseStringsCopy, house)
				continue
			}
			extractI := houseLen - step
			digit := house[extractI]
			donations = append(donations, string(digit))
			if extractI >= houseLen-1 {
				house = house[:extractI]
			} else {
				house = house[:extractI] + house[extractI+1:]
			}

			houseStringsCopy = append(houseStringsCopy, house)
			//fmt.Printf("Houses copy is now: %v\n", houseStringsCopy)

			if slicesEqual(houseStrings, houseStringsCopy) {
				break
				fmt.Println("We are done!! Copy is equal")
			}

		}

		houseStringsCopy[0] = donations[n-1] + houseStringsCopy[0]
		fmt.Println("Donations: ", donations)
		for i := 1; i < len(houseStringsCopy); i++ {
			houseStringsCopy[i] = donations[(i-1)] + houseStringsCopy[i]
		}

		if slicesEqual(houseStrings, houseStringsCopy) {
			break
			fmt.Println("We are done!! Copy is equal")
		}

		houseStrings = houseStringsCopy
		fmt.Printf("Houses copy is now: %v\n", houseStringsCopy)
		step++
	}

	for i, h := range houseStrings {
		hInt, _ := strconv.Atoi(h)
		houses[i] = hInt
	}
	fmt.Println("Done with these houses!")
	return houses
}
