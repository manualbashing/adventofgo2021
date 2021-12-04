package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData() []string {
	file, err := os.Open("data.dat")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func part1() int64 {
	data := readData()
	bitCountArray := make([]int, len(data[0]))
	for _, line := range data {

		chars := strings.Split(line, "")
		for i := 0; i < len(chars); i++ {

			currentBit, err := strconv.Atoi(chars[i])
			check(err)
			bitCountArray[i] += currentBit
		}
	}

	var gammaRate string
	for _, n := range bitCountArray {
		if n > len(data)/2 {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}
	}
	result, err := strconv.ParseInt(gammaRate, 2, 0)
	check(err)
	return result
}
