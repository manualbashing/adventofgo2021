package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	Data [][]string
}

func (b *Board) Mark(number string) {

	for i := 0; i < len(b.Data); i++ {
		for j := 0; j < len(b.Data[i]); j++ {
			if b.Data[i][j] == number {
				b.Data[i][j] = ("x" + number)
			}
		}
	}
}

func (b Board) CheckBingo() bool {

	// Skip empty boards
	if b.Data == nil {
		return false
	}

	// Check Rows
	for _, row := range b.Data {
		if checkBingo(row) {
			return true
		}
	}

	// Check Cols
	for i := 0; i < len(b.Data[0]); i++ {
		var col []string
		for j := 0; j < len(b.Data); j++ {
			col = append(col, b.Data[j][i])
		}
		if checkBingo(col) {
			return true
		}
	}
	return false
}

func (b Board) GetFinalScore() int {
	sum := 0
	for _, row := range b.Data {
		for _, col := range row {
			if strings.HasPrefix(col, "x") {
				continue
			}
			val, err := strconv.Atoi(col)
			check(err)
			sum += val
		}
	}
	return sum
}

func main() {
	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())
	fmt.Println("-- Executing part two")
	fmt.Printf("   Result: %d\n", part2())
}

func readBoards() ([]string, []Board) {
	file, err := os.Open("data.dat")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var boardCollection []Board
	var boardData [][]string
	for scanner.Scan() {

		// Start a new board whenever an empty line is encountered.
		if len(bytes.TrimSpace(scanner.Bytes())) == 0 {
			board := Board{Data: boardData}
			boardCollection = append(boardCollection, board)
			boardData = nil
			continue
		}
		line := strings.TrimSpace(scanner.Text())
		lineArr := regexp.MustCompile(`[, ]+`).Split(line, -1)
		boardData = append(boardData, lineArr)
	}
	drawnNumbers := boardCollection[0].Data[0]
	return drawnNumbers, boardCollection[1:]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkBingo(input []string) bool {
	marked := 0
	for _, c := range input {
		if strings.HasPrefix(c, "x") {
			marked++
		}
	}
	return (marked == 5)
}

func winBingo(numbers []string, boards []Board) (string, Board) {
	for _, n := range numbers {

		for _, b := range boards {

			b.Mark(n)
			if b.CheckBingo() {
				return n, b
			}
		}
	}
	return "", Board{}
}

func looseBingo(numbers []string, boards []Board) (string, Board) {
	var winningBoards []Board
	var winningNumbers []string
	for _, n := range numbers {
		for i := 0; i < len(boards); i++ {
			b := boards[i]
			b.Mark(n)
			if b.CheckBingo() {
				winningBoards = append(winningBoards, b)
				winningNumbers = append(winningNumbers, n)
				boards[i] = Board{} // Remove winning from iteration
			}
		}
	}
	return winningNumbers[len(winningNumbers)-1], winningBoards[len(winningBoards)-1]
}

func part1() int {
	drawnNumbers, boards := readBoards()
	lastNumber, bingoBoard := winBingo(drawnNumbers, boards)
	lastNumberInt, err := strconv.Atoi(lastNumber)
	check(err)
	score := bingoBoard.GetFinalScore() * lastNumberInt
	return score
}

func part2() int {
	drawnNumbers, boards := readBoards()
	lastNumber, bingoBoard := looseBingo(drawnNumbers, boards)
	lastNumberInt, err := strconv.Atoi(lastNumber)
	check(err)
	score := bingoBoard.GetFinalScore() * lastNumberInt
	return score
}
