package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
)

type Direction string

const (
	Left  Direction = "<"
	Right Direction = ">"
	Up    Direction = "^"
	Down  Direction = "v"
)

func (this Direction) IsReverse(other Direction) bool {
	switch this {
	case Left:
		return this == Right
	case Right:
		return this == Left
	case Up:
		return this == Down
	default:
		return this == Up
	}
}

type MoveHistory struct {
	direction Direction
	count     int
	loss      int
}

type Point struct {
	i, j int
	MoveHistory
}

type PathFinder struct {
	heat     [][]int
	min_loss [][]int
	next     []Point
}

func NewPathFinder(grid [][]int) PathFinder {
	n := len(grid)
	k := len(grid[0])

	var min_loss [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < k; j++ {
			row = append(row, math.MaxInt)
		}
		min_loss = append(min_loss, row)
	}
	min_loss[0][0] = 0
	min_loss[0][1] = grid[0][1]
	min_loss[1][0] = grid[1][0]

	next := []Point{
		{0, 1, MoveHistory{Right, 1, grid[0][1]}},
		{1, 0, MoveHistory{Down, 1, grid[1][0]}},
	}
	return PathFinder{grid, min_loss, next}
}

// Check where we can go from this point
func (this *PathFinder) Explore(point Point) {
	var i, j, count, loss int
	max_i, max_j := this.Bounds()

	for _, direction := range []Direction{Left, Right, Up, Down} {

		if point.direction.IsReverse(direction) {
			// cannot go back
			continue
		}

		i = point.i
		j = point.j

		switch direction {
		case Up:
			i--
			if i < 0 {
				continue
			}
		case Down:
			i++
			if i > max_i {
				continue
			}
		case Left:
			j--
			if j < 0 {
				continue
			}
		case Right:
			j++
			if j > max_j {
				continue
			}
		}

		if point.direction == direction {
			count = point.count + 1
			// too many moves in the same direction
			if count > 3 {
				continue
			}
		} else {
			count = 1
		}

		loss = point.loss + this.heat[i][j]
		if loss <= this.min_loss[i][j] {
			this.min_loss[i][j] = loss
		} else {
			// other variation already reached it with smaller loss
			continue
		}

		hist := MoveHistory{direction, count, loss}
		this.next = append(this.next, Point{i, j, hist})
	}
}

func (this PathFinder) Bounds() (int, int) {
	return len(this.heat) - 1, len(this.heat[0]) - 1
}

func (this PathFinder) IsFinal(point Point) bool {
	max_i, max_j := this.Bounds()
	return point.i == max_i && point.j == max_j
}

func (this PathFinder) Loss() int {
	i, j := this.Bounds()
	return this.min_loss[i][j]
}

func (this *PathFinder) FindPath() {
	var current Point
	for {
		slices.SortFunc(this.next, func(a, b Point) int {
			return cmp.Compare(a.loss, b.loss)
		})

		// fmt.Println(current)
		// fmt.Println(this.next)

		current = this.next[0]
		this.next = this.next[1:]

		if len(this.next) == 0 {
			return
		}

		this.Explore(current)
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

	finder := NewPathFinder(grid)

	finder.FindPath()
	// fmt.Println(finder.Loss())
	for _, row := range finder.min_loss {
		for _, x := range row {
			fmt.Printf("%3.d ", x)
		}
		fmt.Println()
	}
}
