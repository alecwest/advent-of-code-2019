package advent

import (
	"bufio"
	"os"
	"strconv"
)

// ReadIntArrayInput reads multiple lines of integers
func ReadIntArrayInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	input := []int{}
	for scanner.Scan() {
		next, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		input = append(input, next)
	}
	return input
}

// ReadStringInput reads input as a single string
func ReadStringInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// ReadStringArrayInput reads input as an array of strings
func ReadStringArrayInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	for scanner.Scan() {
		next := scanner.Text()
		if next == "" {
			break
		}

		input = append(input, next)
	}
	return input
}

// Max returns the greater of two ints
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the lesser of two ints
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Abs returns the absolute value of an int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// NumberIsBetween determines if the target val is within the range of the two given values (may not be sorted)
func NumberIsBetween(val1, val2, target int) bool {
	return target >= Min(val1, val2) && target <= Max(val1, val2)
}
