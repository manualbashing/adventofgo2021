package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	p2 "github.com/manualbashing/adventofgo2021/cmd/day2/part2"
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

func part2() int {

	c := p2.SubmarineControls{
		Aim:        0,
		Horizontal: 0,
		Depth:      0,
	}

	plannedCourse := readPlannedCourse()

	for _, instruction := range plannedCourse {

		x, err := strconv.Atoi(instruction[1])
		check(err)

		switch instruction[0] {

		case "forward":
			c.Forward(x)
		case "up":
			c.Up(x)
		case "down":
			c.Down(x)
		}
	}

	return c.Horizontal * c.Depth
}

func main() {

	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())

	fmt.Println("-- Executing part two")
	fmt.Printf("   Result: %d\n", part2())

}
