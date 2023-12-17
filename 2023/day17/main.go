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

var Directions = [4]Direction{Left, Right, Up, Down}

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

type State struct {
	i, j int
	Direction
	count int
}

type Candidate struct {
	State
	loss int
}

type PathFinder struct {
	heat [][]int
	loss map[State]int
	next []Candidate
}

type Losses map[Direction]int

func NewLosses() Losses {
	this := make(map[Direction]int)
	for _, direction := range Directions {
		this[direction] = math.MaxInt
	}
	return this
}

func (this Losses) Min() int {
	best := math.MaxInt
	for _, val := range this {
		best = min(best, val)
	}
	return best
}

func NewPathFinder(heat [][]int) PathFinder {
	loss := make(map[State]int)
	// for _, direction := range Directions {
	// 	loss[State{0, 0, direction, 1}] = heat[0][0]
	// }

	next := []Candidate{
		{State{0, 1, Right, 1}, heat[0][1]},
		{State{1, 0, Down, 1}, heat[1][0]},
	}
	for _, candidate := range next {
		loss[candidate.State] = candidate.loss
	}

	return PathFinder{heat, loss, next}
}

// Check where we can go from this point
func (this *PathFinder) Explore(from Candidate) {
	max_i, max_j := this.Bounds()

	for _, direction := range Directions {

		if from.IsReverse(direction) {
			// cannot go back
			continue
		}

		current := State{from.i, from.j, direction, 1}

		switch direction {
		case Up:
			current.i--
			if current.i < 0 {
				continue
			}
		case Down:
			current.i++
			if current.i > max_i {
				continue
			}
		case Left:
			current.j--
			if current.j < 0 {
				continue
			}
		case Right:
			current.j++
			if current.j > max_j {
				continue
			}
		}

		if from.Direction == direction {
			current.count = from.count + 1
			// too many moves in the same direction
			if current.count > 3 {
				continue
			}
		}

		loss := from.loss + this.heat[current.i][current.j]
		if prev, ok := this.loss[current]; ok && prev < loss {
			continue
		}

		this.loss[current] = loss
		this.next = append(this.next, Candidate{current, loss})
	}
}

func (this PathFinder) Bounds() (int, int) {
	return len(this.heat) - 1, len(this.heat[0]) - 1
}

func (this PathFinder) MinLossAt(i, j int) int {
	best := math.MaxInt
	for state, loss := range this.loss {
		if state.i == i && state.j == j {
			if loss < best {
				best = loss
			}
		}
	}
	return best
}

func (this PathFinder) IsFinal(state State) bool {
	max_i, max_j := this.Bounds()
	return state.i == max_i && state.j == max_j
}

func (this PathFinder) FinalLoss() int {
	max_i, max_j := this.Bounds()
	return this.MinLossAt(max_i, max_j)
}

func (this *PathFinder) FindPath() {
	var current Candidate
	for {
		// if len(this.next) == 0 {
		// 	return
		// }

		slices.SortFunc(this.next, func(a, b Candidate) int {
			return cmp.Compare(a.loss, b.loss)
		})

		current = this.next[0]
		this.next = this.next[1:]

		if this.IsFinal(current.State) {
			return
		}

		// fmt.Println(current)
		// fmt.Println(this.next)

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

	max_i, max_j := finder.Bounds()
	for i := 0; i <= max_i; i++ {
		for j := 0; j <= max_j; j++ {
			loss := finder.MinLossAt(i, j)
			if loss == math.MaxInt {
				fmt.Print("  ? ")
			} else {
				fmt.Printf("%3.d ", loss)
			}
		}
		fmt.Println()
	}

	fmt.Println(finder.FinalLoss())
}
