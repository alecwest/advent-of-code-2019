package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alecwest/advent-of-code-2019/advent"
)

// NumPossiblePasswords returns all possible passwords given within the input range
func NumPossiblePasswords(input string) int {
	splitInput := strings.Split(input, "-")
	start, _ := strconv.Atoi(splitInput[0])
	end, _ := strconv.Atoi(splitInput[1])

	numPasswords := 0
	for i := start; i <= end; i++ {
		currPassword := fmt.Sprintf("%06d", i)

		// Initialize criteria
		adjacentSame := false
		leftRightDecrease := false
		for j := 1; j < 6; j++ {
			if currPassword[j-1] == currPassword[j] {
				if j == 5 || (j < 5 && currPassword[j+1] != currPassword[j]) {
					adjacentSame = true
				} else {
					for {
						if j < 5 && currPassword[j+1] == currPassword[j] {
							j++
						} else {
							break
						}
					}
				}
			}
			if strings.Compare(string(currPassword[j-1]), string(currPassword[j])) > 0 {
				leftRightDecrease = true
			}
		}
		if adjacentSame && !leftRightDecrease {
			numPasswords++
		}
	}
	return numPasswords
}

func main() {
	input := advent.ReadStringInput()
	println(NumPossiblePasswords(input))
}
