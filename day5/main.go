package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/alecwest/advent-of-code-2019/advent"
)

// OPCODES
const (
	// ADD code - 3 params (instruction1, instruction2, output)
	ADD = 1

	// MULTIPLY code - 3 params (instruction1, instruction2, output)
	MULTIPLY = 2

	// INPUT code - 1 param
	INPUT = 3

	// OUTPUT code - 1 param
	OUTPUT = 4

	// JUMPIFTRUE code - 2 params (boolean, new instruction pointer)
	JUMPIFTRUE = 5

	// JUMPIFFALSE code - 2 params (boolean, new instruction pointer)
	JUMPIFFALSE = 6

	// LESSTHAN code - 3 params (value1, value2, output location)
	LESSTHAN = 7

	// EQUALS code - 3 params (value1, value2, output location)
	EQUALS = 8

	// EXIT code
	EXIT = 99
)

// MODES
const (
	// POSITION mode - interpret associated parameter as the location of the value
	POSITION = 0

	// IMMEDIATE mode - interpret associated parameter as the actual value
	IMMEDIATE = 1
)

// IntCode processes the instructions
func IntCode(input string) int {
	memory := strings.Split(input, ",")
	instructions := make([]int, len(memory))
	for i := range instructions {
		instructions[i], _ = strconv.Atoi(memory[i])
	}
	result := processInstructions(instructions)
	return result
}

func getOpCode(instruction int) int {
	return instruction % 100
}

// getMode parses the given instruction and returns the mode of the target parameter
func getMode(instruction, paramNumber int) int {
	return instruction / int(math.Pow(10, float64(1+paramNumber))) % 10
}

func processInstructions(instructions []int) int {
	var output int = math.MinInt64
	for i := 0; i < len(instructions); {
		code := getOpCode(instructions[i])
		switch code {
		case ADD:
			instructions = add(instructions, i)
			i += 4
		case MULTIPLY:
			instructions = multiply(instructions, i)
			i += 4
		case INPUT:
			instructions = readInput(instructions, i)
			i += 2
		case OUTPUT:
			instructions, output = printOutput(instructions, i)
			i += 2
		case JUMPIFTRUE:
			if getParam(instructions, i, 1) != 0 {
				i = getParam(instructions, i, 2)
			} else {
				i += 3
			}
		case JUMPIFFALSE:
			if getParam(instructions, i, 1) == 0 {
				i = getParam(instructions, i, 2)
			} else {
				i += 3
			}
		case LESSTHAN:
			var result int
			if getParam(instructions, i, 1) < getParam(instructions, i, 2) {
				result = 1
			}
			instructions = setParam(instructions, i, 3, result)
			i += 4
		case EQUALS:
			var result int
			if getParam(instructions, i, 1) == getParam(instructions, i, 2) {
				result = 1
			}
			instructions = setParam(instructions, i, 3, result)
			i += 4
		case EXIT:
			i = len(instructions)
		default:
			fmt.Printf("unexpected code %d\n", code)
		}
	}
	if output > math.MinInt64 {
		return output
	}
	return instructions[0]
}

func add(instructions []int, opcodeIndex int) []int {
	val1 := getParam(instructions, opcodeIndex, 1)
	val2 := getParam(instructions, opcodeIndex, 2)
	return setParam(instructions, opcodeIndex, 3, val1+val2)
}

func multiply(instructions []int, opcodeIndex int) []int {
	val1 := getParam(instructions, opcodeIndex, 1)
	val2 := getParam(instructions, opcodeIndex, 2)
	return setParam(instructions, opcodeIndex, 3, val1*val2)
}

func readInput(instructions []int, opcodeIndex int) []int {
	input := 5
	return setParam(instructions, opcodeIndex, 1, input)
}

func printOutput(instructions []int, opcodeIndex int) ([]int, int) {
	output := getParam(instructions, opcodeIndex, 1)
	fmt.Printf("%d\n", output)
	return instructions, output
}

func getParam(instructions []int, opcodeIndex, paramNumber int) int {
	var param int
	paramIndex := opcodeIndex + paramNumber
	mode := getMode(instructions[opcodeIndex], paramNumber)
	switch mode {
	case POSITION:
		position := instructions[paramIndex]
		param = instructions[position]
	case IMMEDIATE:
		param = instructions[paramIndex]
	default:
		fmt.Printf("Unable to get param with given mode %d", mode)
	}
	return param
}

func setParam(instructions []int, opcodeIndex, paramNumber, value int) []int {
	paramIndex := opcodeIndex + paramNumber
	mode := getMode(instructions[opcodeIndex], paramNumber)
	switch mode {
	case POSITION:
		position := instructions[paramIndex]
		instructions[position] = value
	case IMMEDIATE:
		instructions[paramIndex] = value
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
	IntCode(instructions)
}
