package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData() []int {
	file, err := os.Open("data.dat")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []int
	for scanner.Scan() {
		var current int
		current, err = strconv.Atoi(scanner.Text())
		check(err)
		data = append(data, current)
	}
	return data
}

func part1() int {

	var data []int = readData()
	var larger int = 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			larger++
		}
	}
	return larger
}

func part2() int {

	var data []int = readData()
	var larger int = 0

	/*
		The loop was initially constructed like this:

		var current int = data[i] + data[i+1] + data[i+2]
		var next int = data[i+1] + data[i+2] + data[i+3]

		As this iterations `next` is next iterations `current`, the calculation of this value became
		redundant.

		The following approach gets rid of this calculation.
	*/
	var next int = data[0] + data[1] + data[2]
	for i := 0; i < len(data)-3; i++ {

		var current int = next
		next = data[i+1] + data[i+2] + data[i+3]
		if current < next {
			larger++
		}
	}
	return larger
}

func main() {

	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())
	fmt.Println("-- Executing part two")
	fmt.Printf("   Result: %d\n", part2())
}
