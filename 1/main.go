package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calculateReqiuredFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel >= 0 {
		return fuel
	}
	return 0
}

func improvedCalculateReqiuredFuel(mass int) int {
	fuel := calculateReqiuredFuel(mass)
	sum := fuel
	for fuel > 0 {
		fuel = calculateReqiuredFuel(fuel)
		sum += fuel
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	var mass int

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d", &mass)
		if err != nil {
			log.Fatal(err)
		}
		sum += improvedCalculateReqiuredFuel(mass)

	}

	fmt.Println(sum)
}
