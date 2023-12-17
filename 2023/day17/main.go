package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	current  Point
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

	hist := MoveHistory{Up, 0, 0}
	current := Point{0, 0, hist}
	next := []Point{}
	return PathFinder{grid, min_loss, current, next}
}

// Check where we can go from this point
func (this *PathFinder) Explore(point Point) {
	var i, j, count, loss int
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
			if i > len(this.heat)-1 {
				continue
			}
		case Left:
			j--
			if j < 0 {
				continue
			}
		case Right:
			j++
			if j > len(this.heat[0])-1 {
				continue
			}
		}

		if point.direction == direction {
			count = point.count + 1
			// too many moves in the same direction
			if count >= 3 {
				continue
			}
		} else {
			count = 1
		}

		loss = point.loss + this.heat[i][j]
		if loss < this.min_loss[i][j] {
			this.min_loss[i][j] = loss
		} else {
			// other variation already reached it with smaller loss
			continue
		}

		hist := MoveHistory{direction, count, loss}
		this.next = append(this.next, Point{i, j, hist})
	}
}

func (this PathFinder) IsFinal(point Point) bool {
	return point.i == len(this.heat)-1 && point.j == len(this.heat[0])-1
}

// Find index of the candidate point with smallest heat loss
func (this *PathFinder) Next() int {
	var (
		index int
		best  int = math.MaxInt
	)
	for i, candidate := range this.next {
		// we reached the destination
		if this.IsFinal(candidate) {
			return i
		}

		if candidate.loss < best {
			index = i
			best = candidate.loss
		}
	}
	return index
}

func pop[T any](arr []T, index int) (T, []T) {
	if index >= len(arr) {
		return arr[index], arr[:index]
	} else {
		return arr[index], append(arr[:index], arr[index+1:]...)
	}
}

func (this *PathFinder) FindPath() int {
	for {
		fmt.Println(this.current)
		// fmt.Println(this.next)

		this.Explore(this.current)
		index := this.Next()
		this.current, this.next = pop(this.next, index)

		if len(this.next) == 0 {
			return this.min_loss[len(this.min_loss)-1][len(this.min_loss[0])-1]
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

	for _, row := range finder.min_loss {
		fmt.Println(row)
	}
}
