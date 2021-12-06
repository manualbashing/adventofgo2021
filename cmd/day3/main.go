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

func countBitsOnPosition(data []string, pos int) int {

	var bitCount int
	for _, line := range data {

		currentBitString := strings.Split(line, "")[pos]
		currentBit, err := strconv.Atoi(currentBitString)
		check(err)
		bitCount += currentBit
	}
	return bitCount
}

func part1() int {
	data := readData()
	gammaRate, epsilonRate := getRates(data)
	return binaryMultiplication(gammaRate, epsilonRate)
}

func getMostCommonBit(bitCount int, dataLength int) string {
	mostCommonBit := "0"
	if float32(bitCount) >= float32(dataLength)/2 {
		mostCommonBit = "1"
	}
	return mostCommonBit
}

func getLeastCommonBit(bitCount int, dataLength int) string {
	leastCommonBit := "1"
	if float32(bitCount) >= float32(dataLength)/2 {
		leastCommonBit = "0"
	}
	return leastCommonBit
}

func reduceDataSet(data []string, matchBitFunc func(int, int) string) string {
	pos := 0
	for matchFound := true; matchFound; matchFound = len(data) > 1 {

		bitCount := countBitsOnPosition(data, pos)
		matchBit := matchBitFunc(bitCount, len(data))
		var reducedData []string
		for _, v := range data {
			if string(v[pos]) == matchBit {
				reducedData = append(reducedData, v)
			}
		}
		data = reducedData
		pos++
	}
	return data[0]
}

func binaryMultiplication(x string, y string) int {
	xInt, err := strconv.ParseInt(x, 2, 0)
	check(err)
	yInt, err := strconv.ParseInt(y, 2, 0)
	check(err)
	return int(xInt * yInt)
}

func part2() int {

	data := readData()
	oxGeneratorRating := reduceDataSet(data, getMostCommonBit)
	co2ScrubberRating := reduceDataSet(data, getLeastCommonBit)

	return binaryMultiplication(oxGeneratorRating, co2ScrubberRating)
}
