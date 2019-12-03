package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func (p *point) taxiCabDistance(r *point) int {
	return abs(p.x-r.x) + abs(p.y-r.y)
}

type wire struct {
	width  int
	height int
	points [][]int
}

func (w *wire) getMiddlePoint() point {
	return point{w.width / 2, w.height / 2}
}

func (w *wire) addLine(a, b point, distanceFromMiddle int) {
	if a.x == b.x {
		distance := abs(a.y - b.y)
		sense := 1
		if a.y > b.y {
			sense = -1
		}
		for i := 1; i <= distance; i++ {
			if w.points[a.x][a.y+(i*sense)] == 0 {
				w.points[a.x][a.y+(i*sense)] = distanceFromMiddle
			}
			distanceFromMiddle++
		}
	} else {
		distance := abs(a.x - b.x)
		sense := 1
		if a.x > b.x {
			sense = -1
		}
		for i := 1; i <= distance; i++ {
			if w.points[a.x+(i*sense)][a.y] == 0 {
				w.points[a.x+(i*sense)][a.y] = distanceFromMiddle
			}
			distanceFromMiddle++
		}
	}
}

func newWire(width, height int, segments []string) wire {
	points := make([][]int, width)
	for i := 0; i < width; i++ {
		points[i] = make([]int, height)
	}
	w := wire{
		width,
		height,
		points}
	position := w.getMiddlePoint()
	distanceFromMiddle := 0
	for _, segment := range segments {
		var direction string
		var length int
		fmt.Sscanf(segment, "%1s%d", &direction, &length)
		var finishPoint point
		switch direction {
		case "U":
			finishPoint = point{position.x, position.y + length}
		case "D":
			finishPoint = point{position.x, position.y - length}
		case "R":
			finishPoint = point{position.x + length, position.y}
		case "L":
			finishPoint = point{position.x - length, position.y}
		}
		w.addLine(position, finishPoint, distanceFromMiddle)
		position = finishPoint
		distanceFromMiddle += length
	}
	return w
}

func (w *wire) getIntersections(other *wire) []point {
	var intersections []point
	if w.width != other.width || w.height != other.height {
		return intersections
	}
	for i := 0; i < w.width; i++ {
		for j := 0; j < w.height; j++ {
			if w.points[i][j] != 0 && other.points[i][j] != 0 {
				intersections = append(intersections, point{i, j})
			}
		}
	}
	return intersections
}

func (w *wire) getClosestIntersectionDistanceFromMiddle(other *wire) int {
	intersections := w.getIntersections(other)
	middlePoint := w.getMiddlePoint()
	closestIntersectionDistance := math.MaxInt64
	for _, intersection := range intersections {
		distance := middlePoint.taxiCabDistance(&intersection)
		if distance < closestIntersectionDistance {
			closestIntersectionDistance = distance
		}
	}
	return closestIntersectionDistance
}

func (w *wire) getClosestIntersectionWalkFromMiddle(other *wire) int {
	intersections := w.getIntersections(other)
	minWalkDistance := math.MaxInt64
	for _, intersection := range intersections {
		walkDistance := w.points[intersection.x][intersection.y] + other.points[intersection.x][intersection.y]
		if walkDistance < minWalkDistance {
			minWalkDistance = walkDistance
		}
	}
	return minWalkDistance + 2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wires []wire
	for scanner.Scan() {
		segments := strings.Split(scanner.Text(), ",")
		wires = append(wires, newWire(30000, 30000, segments))
	}
	fmt.Println(wires[0].getClosestIntersectionDistanceFromMiddle(&wires[1]))
	fmt.Println(wires[0].getClosestIntersectionWalkFromMiddle(&wires[1]))
}
