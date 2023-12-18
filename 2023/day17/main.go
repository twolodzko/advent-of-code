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
	None  Direction = "."
)

var Directions = [4]Direction{Left, Right, Up, Down}

func (this Direction) IsReverse(other Direction) bool {
	switch this {
	case Left:
		return other == Right
	case Right:
		return other == Left
	case Up:
		return other == Down
	case Down:
		return other == Up
	default:
		return false
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
	heat         [][]int
	loss         map[State]int
	next         []Candidate
	max_straight int
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

func NewPathFinder(heat [][]int, max_straight int) PathFinder {
	init := heat[0][0]
	loss := make(map[State]int)
	next := []Candidate{
		{State{0, 0, None, 0}, 0},
	}
	for _, candidate := range next {
		loss[candidate.State] = init
	}
	return PathFinder{heat, loss, next, max_straight}
}

// Check where we can go from this point
func (this *PathFinder) Explore(from Candidate) {
	max_i, max_j := this.Bounds()

	for _, direction := range Directions {

		if from.IsReverse(direction) {
			// cannot go back
			continue
		}

		var count int
		if from.Direction == direction {
			count = from.count + 1
			if count > 3 {
				continue
			}
		} else {
			count = 1
		}
		current := State{from.i, from.j, direction, count}

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

		loss := from.loss + this.heat[current.i][current.j]
		if old_loss, ok := this.loss[current]; ok && old_loss <= loss {
			continue
		}

		this.loss[current] = loss
		this.next = append(this.next, Candidate{current, loss})
	}
}

func (this PathFinder) Bounds() (int, int) {
	n := len(this.heat) - 1
	k := len(this.heat[n]) - 1
	return n, k
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
	for len(this.next) > 0 {
		slices.SortFunc(this.next, func(a, b Candidate) int {
			return cmp.Compare(a.loss, b.loss)
		})

		current := this.next[0]
		this.next = this.next[1:]

		if this.IsFinal(current.State) {
			return
		}

		// fmt.Println(current)
		// fmt.Println(this.next)

		this.Explore(current)
	}
}

func (this PathFinder) Show() {
	max_i, max_j := this.Bounds()
	for i := 0; i <= max_i; i++ {
		for j := 0; j <= max_j; j++ {
			loss := this.MinLossAt(i, j)
			if loss == math.MaxInt {
				fmt.Print("  ? ")
			} else {
				fmt.Printf("%3.d ", loss)
			}
		}
		fmt.Println()
	}
}

func part1(grid [][]int) {
	finder := NewPathFinder(grid, 3)
	finder.FindPath()
	// finder.Show()
	fmt.Println(finder.FinalLoss())
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

	part1(grid)
}
