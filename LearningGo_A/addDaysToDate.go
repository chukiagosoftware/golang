package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	Integer, _ := strconv.Atoi(s)
	return Integer
}

func LeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}

var monthMap = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func monthMapByYear(month int, year int) int {
	if month == 2 {
		if LeapYear(year) {
			return 29
		} else {
			return 28
		}
	}
	return monthMap[month]
}

func AddDays(date string, n int) string {
	// TODO: Implement your solution here

	dateParts := strings.Split(date, "-")
	year, month, day := StringToInt(dateParts[0]), StringToInt(dateParts[1]), StringToInt(dateParts[2])
	totalDays := day + n
	daysInMonth := monthMapByYear(month, year)

	for totalDays-daysInMonth > 0 {
		totalDays = totalDays - daysInMonth
		month += 1
		if month > 12 {
			month = 1
			year += 1
		}
		daysInMonth = monthMapByYear(month, year)
	}

	return fmt.Sprintf("%04d-%02d-%02d", year, month, totalDays)
}
