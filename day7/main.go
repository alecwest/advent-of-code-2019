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

// IntCode processes the instructions and returns the result and last opcode
func IntCode(instructions []int, inputChan, outputChan, result chan int) {
	processInstructions(instructions, inputChan, outputChan, result)
}

func parseInstructions(input string) []int {
	memory := strings.Split(input, ",")
	instructions := make([]int, len(memory))
	for i := range instructions {
		instructions[i], _ = strconv.Atoi(memory[i])
	}
	return instructions
}

func getOpCode(instruction int) int {
	return instruction % 100
}

// getMode parses the given instruction and returns the mode of the target parameter
func getMode(instruction, paramNumber int) int {
	return instruction / int(math.Pow(10, float64(1+paramNumber))) % 10
}

func processInstructions(instructions []int, inputChan, outputChan, result chan int) {
	var output int = math.MinInt64
	var code int
	for i := 0; i < len(instructions); {
		code = getOpCode(instructions[i])
		switch code {
		case ADD:
			instructions = add(instructions, i)
			i += 4
		case MULTIPLY:
			instructions = multiply(instructions, i)
			i += 4
		case INPUT:
			instructions = readInput(instructions, i, inputChan)
			i += 2
		case OUTPUT:
			instructions, output = printOutput(instructions, i, outputChan)
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
	result <- output
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

func readInput(instructions []int, opcodeIndex int, inputChan chan int) []int {
	input := <-inputChan
	return setParam(instructions, opcodeIndex, 1, input)
}

func printOutput(instructions []int, opcodeIndex int, outputChan chan int) ([]int, int) {
	output := getParam(instructions, opcodeIndex, 1)
	outputChan <- output
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

func thrusterChain(instructions []int, thrusterPermutation []int) int {
	channels := make([]chan int, len(thrusterPermutation))
	resultChannels := make([]chan int, len(thrusterPermutation))
	for i := range channels {
		channels[i] = make(chan int, 2)
		resultChannels[i] = make(chan int, 1)
	}
	thrusterInstructions := make([][]int, len(thrusterPermutation))
	for i := range thrusterInstructions {
		thrusterInstructions[i] = make([]int, len(instructions)+1)
		copy(thrusterInstructions[i], instructions)
		thrusterInstructions[i][len(instructions)] = i
	}
	for i, channel := range channels {
		channel <- thrusterPermutation[i]
	}
	channels[0] <- 0

	for i := 0; i < len(channels); i++ {
		go IntCode(thrusterInstructions[i], channels[i], channels[(i+1)%len(channels)], resultChannels[i])
	}
	return <-resultChannels[len(thrusterPermutation)-1]
}

func maxThrusterSignal(input string) int {
	instructions := parseInstructions(input)
	maxSignal := math.MinInt64
	var output int
	for _, perm := range advent.Permutations([]int{5, 6, 7, 8, 9}) {
		output = thrusterChain(instructions, perm)
		if output > maxSignal {
			maxSignal = output
		}
	}
	return maxSignal
}

func main() {
	println(maxThrusterSignal(advent.ReadStringInput()))
}
