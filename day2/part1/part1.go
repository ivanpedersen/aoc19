package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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
	instruction := program[instructionPointer]
	// setting initial program parameters
	program[instructionPointer+1] = 12
	program[instructionPointer+2] = 2

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
			fmt.Println(program[0])
			os.Exit(0)
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
