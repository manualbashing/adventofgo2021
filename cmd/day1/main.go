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

func part1() int {

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

	var larger int = 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			larger++
		}
	}
	return larger
}

func main() {

	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())
}
