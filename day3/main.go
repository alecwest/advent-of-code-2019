package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alecwest/advent-of-code-2019/advent"
)

const (
	// UP movement
	UP = "U"

	// RIGHT movement
	RIGHT = "R"

	// DOWN movement
	DOWN = "D"

	// LEFT movement
	LEFT = "L"
)

type point struct {
	x int
	y int
}

type line struct {
	point1    point
	point2    point
	direction string
	length    int
}

type path struct {
	lines []line
}

func printPath(p path) {
	fmt.Printf("%+v\n", p)
}

func printLine(l line) {
	fmt.Printf("%+v\n", l)
}

func printPoint(p point) {
	fmt.Printf("%+v\n", p)
}

func parsePath(input []string) path {
	wire := make([]line, len(input))

	wire[0].point1 = point{0, 0}
	for i := range wire {
		if i != 0 {
			wire[i].point1 = wire[i-1].point2
		}

		movement := input[i]
		distance, err := strconv.Atoi(string(movement[1:]))
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not parse %v to an int. Assigning 0", movement[1:])
			distance = 0
		}

		switch wire[i].direction = string(movement[0]); wire[i].direction {
		case UP:
			wire[i].point2 = point{wire[i].point1.x, wire[i].point1.y + distance}
		case RIGHT:
			wire[i].point2 = point{wire[i].point1.x + distance, wire[i].point1.y}
		case DOWN:
			wire[i].point2 = point{wire[i].point1.x, wire[i].point1.y - distance}
		case LEFT:
			wire[i].point2 = point{wire[i].point1.x - distance, wire[i].point1.y}
		default:
			fmt.Fprintf(os.Stderr, "invalid direction given: %s", wire[i].direction)
		}
		wire[i].length = distance
	}
	return path{wire}
}

func linesArePerpendicular(l1, l2 line) bool {
	if isVertical(l1) && !isVertical(l2) {
		return true
	} else if !isVertical(l1) && isVertical(l2) {
		return true
	}
	return false
}

func linesIntersect(hLine, vLine line) bool {
	potentialIntersection := point{vLine.point1.x, hLine.point1.y}
	return advent.NumberIsBetween(hLine.point1.x, hLine.point2.x, potentialIntersection.x) && advent.NumberIsBetween(vLine.point1.y, vLine.point2.y, potentialIntersection.y)
}

func getIntersection(hLine, vLine line) *point {
	potentialIntersection := point{vLine.point1.x, hLine.point1.y}
	if advent.NumberIsBetween(hLine.point1.x, hLine.point2.x, potentialIntersection.x) && advent.NumberIsBetween(vLine.point1.y, vLine.point2.y, potentialIntersection.y) {
		return &potentialIntersection
	}
	return nil
}

func comparePointDistance(p1, p2 point) int {
	return advent.Abs((advent.Abs(p1.x) + advent.Abs(p1.y)) - (advent.Abs(p2.x) + advent.Abs(p2.y)))
}

func isVertical(l line) bool {
	return l.direction == UP || l.direction == DOWN
}

func isCentralPort(p point) bool {
	return p.x == 0 && p.y == 0
}

// BestIntersection returns the manhattan distance from the best intersection of the two paths
func BestIntersection(strPath1, strPath2 string) int {
	input1 := strings.Split(strPath1, ",")
	input2 := strings.Split(strPath2, ",")
	path1 := parsePath(input1)
	path2 := parsePath(input2)
	var shortestDistance int

	lineDistance1 := 0
	for _, line1 := range path1.lines {
		lineDistance2 := 0
		for _, line2 := range path2.lines {
			if !linesArePerpendicular(line1, line2) {
				lineDistance2 += line2.length
				continue
			}
			var vLine line
			var hLine line
			if isVertical(line1) {
				vLine = line1
				hLine = line2
			} else {
				vLine = line2
				hLine = line1
			}

			intersection := getIntersection(hLine, vLine)
			if intersection != nil {
				// Get length of lines taken to get to intersection
				partialLineDistance1 := comparePointDistance(hLine.point1, *intersection)
				partialLineDistance2 := comparePointDistance(vLine.point1, *intersection)
				intersectDistance := lineDistance1 + lineDistance2 + partialLineDistance1 + partialLineDistance2
				if shortestDistance == 0 || intersectDistance < shortestDistance {
					shortestDistance = intersectDistance
				}
			}
			lineDistance2 += line2.length
		}
		lineDistance1 += line1.length
	}
	return shortestDistance
}

func main() {
	input := advent.ReadStringArrayInput()
	println(BestIntersection(input[0], input[1]))
}
