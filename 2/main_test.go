package main

import (
	"fmt"
	"testing"
)

func TestExamples1(t *testing.T) {
	input := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	expectedOutput := 3500
	processNumbers(input)
	fmt.Println(input[0], ":", expectedOutput)
	if input[0] != expectedOutput {
		t.Fail()
	}
}

func TestExamples2(t *testing.T) {
	input := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expectedOutput := 30
	processNumbers(input)
	fmt.Println(input[0], ":", expectedOutput)
	if input[0] != expectedOutput {
		t.Fail()
	}
}
