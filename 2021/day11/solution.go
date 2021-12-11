package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) ([][]int, error) {
	var arr [][]int

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}
		if err != nil {
			return arr, err
		}
		arr = append(arr, []int{})
		for _, r := range line {
			d, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}
			arr[i] = append(arr[i], d)
		}
		i += 1
	}
	err = scanner.Err()
	return arr, err
}

type Point struct {
	i, j int
}

func (p *Point) Neighbors(maxSize int) []Point {
	var neighbors []Point
	for _, di := range []int{-1, 0, +1} {
		for _, dj := range []int{-1, 0, +1} {
			if di == 0 && dj == 0 {
				continue
			}
			i := p.i + di
			j := p.j + dj
			if i < 0 || i >= maxSize || j < 0 || j >= maxSize {
				continue
			}
			neighbors = append(neighbors, Point{i, j})
		}
	}
	return neighbors
}

type Queue struct {
	items []Point
}

func newQueue(size int) Queue {
	var queue Queue
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			queue.items = append(queue.items, Point{i, j})
		}
	}
	return queue
}

func (q *Queue) Push(i, j int) {
	q.items = append(q.items, Point{i, j})
}

func (q *Queue) Pop() Point {
	if len(q.items) == 0 {
		return Point{}
	}
	p := q.items[0]
	if len(q.items) > 1 {
		q.items = q.items[1:]
	} else {
		q.items = nil
	}
	return p
}

func (q *Queue) HasNext() bool {
	return len(q.items) > 0
}

func printBoard(board [][]int) {
	size := len(board)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] > 9 {
				fmt.Print("* ")
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func deepcopy(inp [][]int) [][]int {
	out := make([][]int, len(inp))
	for i := 0; i < len(inp); i++ {
		out[i] = make([]int, len(inp[i]))
		for j := 0; j < len(inp[i]); j++ {
			out[i][j] = inp[i][j]
		}
	}
	return out
}

func run(board [][]int, steps int) int {
	flashes := 0
	size := len(board)

	for step := 0; step < steps; step++ {
		queue := newQueue(size)
		for queue.HasNext() {
			point := queue.Pop()
			board[point.i][point.j] += 1
			if board[point.i][point.j] == 10 {
				flashes += 1
				queue.items = append(queue.items, point.Neighbors(size)...)
			}
		}

		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] > 9 {
					board[i][j] = 0
				}
			}
		}
	}
	return flashes
}

func firstSynchronization(board [][]int) int {
	size := len(board)
	step := 0
	for {
		step += 1
		queue := newQueue(size)
		for queue.HasNext() {
			point := queue.Pop()
			board[point.i][point.j] += 1
			if board[point.i][point.j] == 10 {
				queue.items = append(queue.items, point.Neighbors(size)...)
			}
		}

		flashes := 0
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] > 9 {
					board[i][j] = 0
					flashes += 1
				}
			}
		}

		if flashes == size*size {
			return step
		}
	}
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

	result1 := run(deepcopy(arr), 100)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := firstSynchronization(deepcopy(arr))
	fmt.Printf("Puzzle 2: %v\n", result2)
}
