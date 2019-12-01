package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: takes a file of integers as only argument.")
		os.Exit(1)
	}

	contents, _ := ioutil.ReadFile(os.Args[1])

	sum := 0
	for _, l := range strings.Split(string(contents), "\n") {
		mass, _ := strconv.ParseFloat(l, 64)
		fuel := int(math.Floor(mass/3)) - 2
		sum += fuel
	}

	fmt.Printf("Total fuel needed: %d.\n", sum)
}
