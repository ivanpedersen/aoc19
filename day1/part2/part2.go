package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func extraFuelFunc(x int) int {
	fuel := int(math.Floor(float64(x/3))) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + extraFuelFunc(fuel)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: takes a file of integers as only argument.")
		os.Exit(1)
	}

	contents, _ := ioutil.ReadFile(os.Args[1])

	sum := 0
	extraFuelSum := 0
	for _, l := range strings.Split(string(contents), "\n") {
		mass, _ := strconv.ParseFloat(l, 64)
		fuel := int(math.Floor(mass/3)) - 2
		extraFuel := extraFuelFunc(fuel)
		sum += fuel
		extraFuelSum += extraFuel
	}

	fmt.Printf("Total fuel needed: %d. Total extra fuel needed: %d. All in all: %d.\n", sum, extraFuelSum, sum+extraFuelSum)
}
