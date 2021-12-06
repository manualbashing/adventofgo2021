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
	fmt.Println("-- Executing part two")
	fmt.Printf("   Result: %d\n", part2())
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

func countBits(data []string) []int {

	bitCountArray := make([]int, len(data[0]))
	for _, line := range data {

		chars := strings.Split(line, "")
		for i := 0; i < len(chars); i++ {

			currentBit, err := strconv.Atoi(chars[i])
			check(err)
			bitCountArray[i] += currentBit
		}
	}
	return bitCountArray
}

func getRates(data []string) (string, string) {

	bitCountArray := countBits(data)
	var gammaRate string
	var epsilonRate string

	for _, n := range bitCountArray {
		if n > len(data)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}
	return gammaRate, epsilonRate
}

func part1() int64 {
	data := readData()
	gammaRate, epsilonRate := getRates(data)

	gammaRateInt, err := strconv.ParseInt(gammaRate, 2, 0)
	check(err)
	epsilonRateInt, err := strconv.ParseInt(epsilonRate, 2, 0)
	check(err)
	return gammaRateInt * epsilonRateInt
}

func part2() int {

	data := readData()
	return len(data)
}
