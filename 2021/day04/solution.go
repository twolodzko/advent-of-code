package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	var arr []string

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}
		arr = append(arr, line)
	}
	err = scanner.Err()
	return arr, err
}

type (
	Array  []int
	Board  [][]int
	Boards []Board
)

func (b Board) row(pos int) (Array, error) {
	if pos >= len(b) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return b[pos], nil
}

func (b Board) col(pos int) (Array, error) {
	if pos >= len(b[0]) {
		return nil, fmt.Errorf("index out of bounds")
	}
	var col []int
	for _, row := range b {
		col = append(col, row[pos])
	}
	return col, nil
}

func (b Board) Iter() <-chan int {
	ch := make(chan int)
	go func() {
		for _, row := range b {
			for _, x := range row {
				ch <- x
			}
		}
		close(ch)
	}()
	return ch
}

func isIn(x int, set []int) bool {
	for _, y := range set {
		if x == y {
			return true
		}
	}
	return false
}

func (a Array) consistsOf(numbers []int) bool {
	for _, x := range a {
		if !isIn(x, numbers) {
			return false
		}
	}
	return true
}

func strToIntArray(str []string) ([]int, error) {
	var arr []int
	for _, x := range str {
		i, err := strconv.Atoi(x)
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}
	return arr, nil
}

func parse(arr []string) ([]int, Boards, error) {
	var (
		numbers []int
		board   Board
		boards  Boards
	)

	fields := strings.Split(arr[0], ",")
	numbers, err := strToIntArray(fields)
	if err != nil {
		return nil, nil, err
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] == "" {
			if len(board) > 0 {
				boards = append(boards, board)
				board = Board{}
			}
			continue
		}
		row := strings.Fields(arr[i])
		values, err := strToIntArray(row)
		if err != nil {
			return nil, nil, err
		}
		board = append(board, values)
	}
	boards = append(boards, board)

	return numbers, boards, nil
}

func (b Board) allNotIn(numbers []int) []int {
	var arr []int
	for x := range b.Iter() {
		if !isIn(x, numbers) {
			arr = append(arr, x)
		}
	}
	return arr
}

func (b Board) isWinning(pattern []int) (bool, error) {
	var (
		arr Array
		err error
	)
	for i := 0; i < len(b); i++ {
		arr, err = b.row(i)
		if err != nil {
			return false, err
		}
		if arr.consistsOf(pattern) {
			return true, nil
		}

		arr, err = b.col(i)
		if err != nil {
			return false, err
		}
		if arr.consistsOf(pattern) {
			return true, nil
		}
	}
	return false, nil
}

func (b Boards) matching(pattern []int) (Array, int, error) {
	for boardNum, board := range b {
		for i := 0; i < len(b[0]); i++ {
			won, err := board.isWinning(pattern)
			if err != nil {
				return nil, boardNum, err
			}
			if won {
				notIn := board.allNotIn(pattern)
				return notIn, boardNum, nil
			}
		}
	}
	return nil, 0, nil
}

func getMatch(numbers []int, boards Boards) (Array, int, error) {
	start := len(boards[0])
	for maxNum := start; maxNum < len(numbers); maxNum++ {
		pattern := numbers[:maxNum]
		called := pattern[len(pattern)-1]

		notIn, _, err := boards.matching(pattern)
		if err != nil {
			return nil, 0, err
		}
		if notIn != nil {
			return notIn, called, nil
		}
	}
	return nil, 0, fmt.Errorf("no match")
}

func (a Array) Sum() int {
	total := 0
	for _, x := range a {
		total += x
	}
	return total
}

func score(numbers []int, boards Boards) (int, error) {
	notIn, called, err := getMatch(numbers, boards)
	if err != nil {
		return 0, err
	}
	total := notIn.Sum()
	return called * total, err
}

func remove(arr []int, x int) []int {
	for i, y := range arr {
		if x == y {
			if i == len(arr)-1 {
				return arr[:i]
			}
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func getLastMatch(numbers []int, boards Boards) (Array, int, error) {
	var boardsLeft []int
	for i := 0; i < len(boards); i++ {
		boardsLeft = append(boardsLeft, i)
	}

	start := len(boards[0])
	for maxNum := start; maxNum < len(numbers); maxNum++ {
		pattern := numbers[:maxNum]
		called := pattern[len(pattern)-1]

		for i, board := range boards {
			won, err := board.isWinning(pattern)
			if err != nil {
				return nil, 0, err
			}
			if won {
				boardsLeft = remove(boardsLeft, i)
			}
			if len(boardsLeft) == 0 {
				notIn := board.allNotIn(pattern)
				return notIn, called, nil
			}
		}
	}
	return nil, 0, fmt.Errorf("no match")
}

func lastScore(numbers []int, boards Boards) (int, error) {
	arr, called, err := getLastMatch(numbers, boards)
	if err != nil {
		return 0, err
	}
	total := arr.Sum()
	return called * total, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	arr, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	numbers, boards, err := parse(arr)
	if err != nil {
		log.Fatal(err)
	}

	result1, err := score(numbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2, err := lastScore(numbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Puzzle 2: %v\n", result2)
}
