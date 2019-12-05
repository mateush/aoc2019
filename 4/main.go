package main

import (
	"fmt"
	"strconv"
)

func doDigitsNotDecrease(digits []int) bool {
	current := digits[0]
	for _, d := range digits[1:] {
		if d < current {
			return false
		}
		current = d
	}
	return true
}

func areTwoTheSame(digits []int) bool {
	current := digits[0]
	same := 0
	for _, d := range digits[1:] {
		if d == current {
			same++
		} else {
			if same == 1 {
				return true
			}
			same = 0

		}
		current = d
	}
	return same == 1
}

func checkNumber(n int) bool {
	nString := strconv.Itoa(n)
	var digits []int
	for i := 0; i < len(nString); i++ {
		digit, _ := strconv.Atoi(string(nString[i]))
		digits = append(digits, digit)
	}
	return doDigitsNotDecrease(digits) && areTwoTheSame(digits)
}

func main() {
	counter := 0
	for i := 248345; i <= 746315; i++ {
		if checkNumber(i) {
			counter++
		}
	}
	fmt.Println(counter)
}
