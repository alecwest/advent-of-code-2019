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
	instructions = processInstructions(instructions)
	return instructions[0]
}

func getOpCode(instruction int) int {
	return instruction % 100
}

// getMode parses the given instruction and returns the mode of the target parameter
func getMode(instruction, paramNumber int) int {
	return instruction / int(math.Pow(10, float64(1+paramNumber))) % 10
}

func processInstructions(instructions []int) []int {
	var numParameters int
	for i := 0; i < len(instructions); i += numParameters + 1 {
		code := getOpCode(instructions[i])
		switch code {
		case ADD:
			numParameters = 3
			instructions = add(instructions, i)
		case MULTIPLY:
			numParameters = 3
			instructions = multiply(instructions, i)
		case INPUT:
			numParameters = 1
			instructions = readInput(instructions, i)
		case OUTPUT:
			numParameters = 1
			instructions = printOutput(instructions, i)
		case EXIT:
			numParameters = 1
			return instructions
		default:
			fmt.Printf("unexpected code %d\n", code)
		}
	}
	return instructions
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
	input := 1
	return setParam(instructions, opcodeIndex, 1, input)
}

func printOutput(instructions []int, opcodeIndex int) []int {
	fmt.Printf("%d\n", getParam(instructions, opcodeIndex, 1))
	return instructions
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
