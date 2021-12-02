package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readPlannedCourse() [][]string {
	file, err := os.Open("data.dat")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var plannedCourse [][]string
	for scanner.Scan() {

		instruction := strings.Split(scanner.Text(), " ")
		plannedCourse = append(plannedCourse, instruction)
	}
	return plannedCourse
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() int {

	var horizontal int = 0
	var depth int = 0
	plannedCourse := readPlannedCourse()

	for _, instruction := range plannedCourse {

		value, err := strconv.Atoi(instruction[1])
		check(err)
		switch instruction[0] {

		case "forward":
			horizontal += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	return horizontal * depth
}

func main() {

	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())

}
