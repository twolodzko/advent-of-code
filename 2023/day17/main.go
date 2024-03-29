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
		return other == Right
	case Right:
		return other == Left
	case Up:
		return other == Down
	case Down:
		return other == Up
	default:
		panic("not reachable")
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
	min_straight int
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

func NewPathFinder(heat [][]int, min_straight, max_straight int) PathFinder {
	init := heat[0][0]
	loss := make(map[State]int)
	next := []Candidate{
		{State{0, 1, Right, 1}, heat[0][1]},
		{State{1, 0, Down, 1}, heat[1][0]},
	}
	loss[State{0, 0, Down, 0}] = init
	for _, candidate := range next {
		loss[candidate.State] = candidate.loss
	}
	return PathFinder{heat, loss, next, min_straight, max_straight}
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
			if count > this.max_straight {
				continue
			}
		} else {
			if from.count < this.min_straight {
				continue
			}
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
			if loss < best && state.count >= this.min_straight {
				best = loss
			}
		}
	}
	return best
}

func (this PathFinder) IsFinal(state State) bool {
	max_i, max_j := this.Bounds()
	return state.i == max_i && state.j == max_j && state.count >= this.min_straight
}

func (this PathFinder) FinalLoss() int {
	max_i, max_j := this.Bounds()
	return this.MinLossAt(max_i, max_j)
}

// Use Dijkstra's algorithm to find the best path
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
	finder := NewPathFinder(grid, 0, 3)
	finder.FindPath()
	// finder.Show()
	fmt.Println(finder.FinalLoss())
}

func part2(grid [][]int) {
	finder := NewPathFinder(grid, 4, 10)
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
	part2(grid)
}
