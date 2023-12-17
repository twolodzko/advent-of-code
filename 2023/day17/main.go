package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Direction string

const (
	Left  Direction = "Left"
	Right Direction = "Right"
	Down  Direction = "Down"
)

func (this Direction) String() string {
	return string(this)
}

type Point struct {
	i, j int
}

func (this Point) Equal(other Point) bool {
	return this.i == other.i && this.j == other.j
}

// type MoveHistory struct {
// 	direction Direction
// 	count     int
// }

// func (this MoveHistory) To(direction Direction) MoveHistory {
// 	move := MoveHistory{direction, 1}
// 	if this.direction == direction {
// 		move.count += this.count
// 	}
// 	return move
// }

type PathFinder struct {
	grid     [][]int
	loss     [][]int
	explored [][]bool
	current  Point
	final    Point
	next     map[Point]bool
	// history  [][]MoveHistory
}

func Matrix[T any](n, k int) [][]T {
	var mtx [][]T
	for i := 0; i < n; i++ {
		row := make([]T, k)
		mtx = append(mtx, row)
	}
	return mtx
}

func NewPathFinder(grid [][]int) PathFinder {
	n := len(grid)
	k := len(grid[0])

	var loss [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < k; j++ {
			row = append(row, math.MaxInt)
		}
		loss = append(loss, row)
	}
	loss[0][0] = grid[0][0]

	var explored [][]bool
	for i := 0; i < n; i++ {
		explored = append(explored, make([]bool, k))
	}
	explored[0][0] = true

	// var prev [][]MoveHistory
	// for i := 0; i < n; i++ {
	// 	var row []MoveHistory
	// 	for j := 0; j < k; j++ {
	// 		row = append(row, MoveHistory{Direction("Left"), 0})
	// 	}
	// 	prev = append(prev, row)
	// }

	current := Point{0, 0}
	final := Point{len(grid) - 1, len(grid[0]) - 1}
	next := make(map[Point]bool)
	return PathFinder{grid, loss, explored, current, final, next} //, prev}
}

func (this Point) Move(direction Direction) Point {
	switch direction {
	case Left:
		this.j--
	case Right:
		this.j++
	default:
		this.i++
	}
	return this
}

// Check where we can go from `this.current`
func (this *PathFinder) Explore() {
	for _, direction := range []Direction{"Left", "Right", "Down"} {
		location := this.current.Move(direction)

		if location.i < 0 || location.i >= len(this.grid) || location.j < 0 || location.j >= len(this.grid[0]) {
			continue
		}
		if this.explored[location.i][location.j] {
			continue
		}
		// hist := this.history[this.current.x][this.current.y].To(direction)
		// if hist.count >= 3 {
		// 	continue
		// }

		this.next[location] = true

		// this.history[location.x][location.y] = hist
		loss := this.loss[this.current.i][this.current.j] + this.grid[location.i][location.j]
		this.loss[location.i][location.j] = min(this.loss[location.i][location.j], loss)
	}
}

// Switch `this.current` to the candidate from `this.next` with smallest loss
func (this PathFinder) Next() Point {
	best := math.MaxInt
	var move Point
	for candidate, ok := range this.next {
		if ok {
			// we reached the destination
			if candidate.Equal(this.final) {
				return candidate
			}

			loss := this.loss[candidate.i][candidate.j]
			if loss < best {
				move = candidate
				best = loss
			}
		}
	}
	return move
}

func (this *PathFinder) FindPath() int {
	for {
		this.Explore()
		move := this.Next()

		this.next[move] = false
		this.explored[move.i][move.j] = true
		this.current = move

		if move.Equal(this.final) {
			return this.loss[move.i][move.j]
		}
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		var row []int
		for _, r := range scanner.Text() {
			if r < '0' || r > '9' {
				panic("not a number")
			}
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	finder := NewPathFinder(grid)
	fmt.Println(finder.FindPath())
}
