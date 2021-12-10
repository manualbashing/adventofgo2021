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

func (b *Board) CheckBingo() bool {
	// Check Rows
	for _, line := range b.Data {
		marked := 0
		for _, c := range line {
			if strings.HasPrefix(c, "x") {
				marked++
			}
		}
		if marked == 5 {
			return true
		}
		marked = 0
	}

	// Check Cols

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
		line := regexp.MustCompile(`[, ]+`).Split(scanner.Text(), -1)
		boardData = append(boardData, line)
	}
	drawnNumbers := boardCollection[0].Data[0]
	return drawnNumbers, boardCollection[1:]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1() int {
	drawnNumbers, boards := readBoards()
	fmt.Println("Drawn numbers: ", drawnNumbers)

	for _, b := range boards {
		b.Mark("85")
		b.Mark("78")
	}

	fmt.Print("Bingo boards: ", boards)
	return 1
}
