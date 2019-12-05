package main

import (
	"../advent"
	"fmt"
	"strings"
	"strconv"
)

const (
	// ADD code - 3 params (instruction1, instruction2, output)
	ADD = 1

	// MULTIPLY code - 3 params (instruction1, instruction2, output)
	MULTIPLY = 2

	// EXIT code
	EXIT = 99
)

// IntCode processes the instructions
func IntCode(input string) int {
	memory := strings.Split(input, ",")
	instructions := make([]int, len(memory))
	for i := range instructions {
		instructions[i], _ = strconv.Atoi(memory[i])
	}
	instructions = processInstructions(instructions)
	return instructions[0]
}

func processInstructions(instructions []int) []int {
	for i := 0; i < len(instructions); i += 4 {
		if instructions[i] == EXIT {
			break
		}
		code := instructions[i]
		pos1 := instructions[i+1]
		pos2 := instructions[i+2]
		pos3 := instructions[i+3]
		val1 := instructions[pos1]
		val2 := instructions[pos2]
		switch code {
			case ADD: {
				instructions[pos3] = val1 + val2
			}
			case MULTIPLY: {
				instructions[pos3] = val1 * val2
			}
			default: {
				println(fmt.Errorf("unexpected code %d", code))
			}
		}
	}
	return instructions
}

func injectNounAndVerb(instructions string, noun int, verb int) string {
	codes := strings.Split(instructions, ",")
	codes[1] = strconv.Itoa(noun)
	codes[2] = strconv.Itoa(verb)
	return strings.Join(codes, ",")
}

func main() {
	instructions := advent.ReadStringInput()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			instructions = injectNounAndVerb(instructions, i, j)
			result := IntCode(instructions)
			if result == 19690720 {
				fmt.Printf("noun: %d, verb: %d\n", i, j)
				break
			}
		}
	}
}