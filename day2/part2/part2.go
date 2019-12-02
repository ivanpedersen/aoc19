package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func intCode(instructionPointer int, program []int, instructions map[string]int) int {
	instruction := program[instructionPointer]
	for {
		parameter1 := program[instructionPointer+1]
		parameter2 := program[instructionPointer+2]
		address := program[instructionPointer+3]

		switch instruction {
		case instructions["add"]:
			program[address] = program[parameter1] + program[parameter2]
		case instructions["multiply"]:
			program[address] = program[parameter1] * program[parameter2]
		case instructions["halt"]:
			return program[0]
		default:
			// error
			fmt.Printf("Unsupported instruction at: %d\n", instructionPointer)
			fmt.Println(program)
			os.Exit(1)
		}
		// increment program counter
		instructionPointer += 4
		instruction = program[instructionPointer]
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: takes a file of integers as only argument.")
		os.Exit(1)
	}

	bytes, _ := ioutil.ReadFile(os.Args[1])
	strings := strings.Split(string(bytes), ",")
	program := make([]int, len(strings))

	for i, s := range strings {
		program[i], _ = strconv.Atoi(s)
	}

	instructions := map[string]int{
		"add":      1,
		"multiply": 2,
		"halt":     99,
	}

	// first instruction at [0]
	instructionPointer := 0

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program[1] = noun
			program[2] = verb
			variable := intCode(instructionPointer, program, instructions)
			if variable == 19690720 { // Magic value
				fmt.Println((100 * noun) + verb)
				os.Exit(0)
			} else {
				// reset program
				program = make([]int, len(strings))
				for i, s := range strings {
					program[i], _ = strconv.Atoi(s)
				}
			}
		}
	}
}
