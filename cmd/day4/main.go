package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
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

func main() {
	fmt.Println("-- Executing part one")
	fmt.Printf("   Result: %d\n", part1())
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

func playBingo(numbers []string, boards []Board) (string, Board) {
	for _, n := range numbers {
		fmt.Println("Drawn number: ", n)
		for _, b := range boards {

			b.Mark(n)
			if b.CheckBingo() {
				return n, b
			}
		}
	}
	return "", Board{}
}

func part1() int {
	drawnNumbers, boards := readBoards()
	lastNumber, bingoBoard := playBingo(drawnNumbers, boards)
	fmt.Println("Last number: ", lastNumber)
	fmt.Println("Bingo Board: ", bingoBoard)
	return 1
}
