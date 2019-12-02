package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func doAddition(numbers []int, position int) {
	numbers[numbers[position+3]] = numbers[numbers[position+1]] + numbers[numbers[position+2]]
}

func doMultiplication(numbers []int, position int) {
	numbers[numbers[position+3]] = numbers[numbers[position+1]] * numbers[numbers[position+2]]
}

func processNumbers(numbers []int) int {
	position := 0

	for position < len(numbers) {
		optcode := numbers[position]
		switch optcode {
		case 99:
			return numbers[0]
		case 1:
			doAddition(numbers, position)
			position += 4
		case 2:
			doMultiplication(numbers, position)
			position += 4
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(bufio.NewReader(file))

	stringNumbers, error := r.Read()
	if error != nil {
		log.Fatal(error)
	}
	numbers := make([]int, len(stringNumbers))

	for i := 0; i < len(stringNumbers); i++ {
		numbers[i], _ = strconv.Atoi(stringNumbers[i])
	}

	inputCopy := make([]int, len(numbers))

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			copy(inputCopy, numbers)
			inputCopy[1] = i
			inputCopy[2] = j
			processNumbers(inputCopy)
			if inputCopy[0] == 19690720 {
				fmt.Println(i, ",", j)
				break
			}
		}
	}
	// numbers[1] = 12
	// numbers[2] = 2
	// fmt.Println(processNumbers(numbers))
}
