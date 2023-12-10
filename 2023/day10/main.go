package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction string

const (
	North Direction = "North"
	South Direction = "South"
	West  Direction = "West"
	East  Direction = "East"
)

type Pipe rune

const (
	Vertical      Pipe = '|'
	Horizontal    Pipe = '-'
	NorthEastBend Pipe = 'L'
	NorthWestBend Pipe = 'J'
	SouthWestBend Pipe = '7'
	SouthEastBend Pipe = 'F'
	Start         Pipe = 'S'
	Ground        Pipe = '.'
)

func (this Pipe) Directions() (bool, bool, bool, bool) {
	switch this {
	case Vertical:
		return true, true, false, false
	case Horizontal:
		return false, false, true, true
	case NorthEastBend:
		return true, false, true, false
	case NorthWestBend:
		return true, false, false, true
	case SouthWestBend:
		return false, true, false, true
	case SouthEastBend:
		return false, true, true, false
	case Start:
		return true, true, true, true
	default:
		return false, false, false, false
	}
}

func (this Pipe) IsFinal() bool {
	return this == 'S' || this == '.'
}

func (this Pipe) String() string {
	return string(this)
}

type Explorer struct {
	grid                 [][]Pipe
	start_row, start_col int
	distances            [][]int
}

func Matrix(n, k int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	return matrix
}

func NewExplorer(grid [][]Pipe, start_row, start_col int) Explorer {
	distances := Matrix(len(grid), len(grid[0]))
	return Explorer{grid, start_row, start_col, distances}
}

func (e Explorer) move(from Direction, i, j, distance int) {
	if i >= 0 && i < len(e.grid) && j >= 0 && j < len(e.grid[i]) {
		this := e.grid[i][j]
		// fmt.Printf("moving from %-5.5v to %d %d (%v)\n", from, i, j, this)

		if !this.IsFinal() {
			north, south, east, west := this.Directions()

			// the direction where we came from
			// * if there is no entry, we end our trip
			// * otherwise, we close the way back, so we move in one direction
			switch from {
			case North:
				if !north {
					return
				}
				north = false
			case South:
				if !south {
					return
				}
				south = false
			case East:
				if !east {
					return
				}
				east = false
			case West:
				if !west {
					return
				}
				west = false
			}

			if e.distances[i][j] == 0 {
				e.distances[i][j] = distance
			} else {
				e.distances[i][j] = min(e.distances[i][j], distance)
			}

			switch {
			case north:
				e.move(South, i-1, j, distance+1)
			case south:
				e.move(North, i+1, j, distance+1)
			case west:
				e.move(East, i, j-1, distance+1)
			case east:
				e.move(West, i, j+1, distance+1)
			}
		}
	}
}

func (e Explorer) explore() {
	e.move(South, e.start_row-1, e.start_col, 1)
	e.move(North, e.start_row+1, e.start_col, 1)
	e.move(East, e.start_row, e.start_col-1, 1)
	e.move(West, e.start_row, e.start_col+1, 1)
}

func part1(explorer Explorer) {
	largest := 0
	for _, row := range explorer.distances {
		for _, x := range row {
			largest = max(largest, x)
		}
	}
	fmt.Println(largest)
}

func part2(explorer Explorer) {
	dist := explorer.distances
	// so that we count 'S' as part of the path
	dist[explorer.start_row][explorer.start_col] = 1

	// tmp := Matrix(len(dist), len(dist[0]))

	result := 0
	var (
		inside bool
		prev   Pipe
	)
	for i, row := range dist {
		// we always start outside
		inside = false
		for j, d := range row {
			// non-zero distance means this is the valid path
			if d > 0 {
				this := explorer.grid[i][j]
				// we check for boundaries of the polygon
				switch this {
				case '|':
					inside = !inside
					prev = this
				case 'L', 'F':
					prev = this
				case 'J':
					// we ignore '-'
					if prev == 'L' {
						// upwards U-turn
						prev = this
					} else if prev == 'F' {
						inside = !inside
						prev = this
					}
				case '7':
					// we ignore '-'
					if prev == 'F' {
						// downwards U-turn
						prev = this
					} else if prev == 'L' {
						inside = !inside
						prev = this
					}
				}
			} else if inside {
				result++
				// tmp[i][j] = 1
			}
		}
	}

	// for _, row := range tmp {
	// 	fmt.Println(row)
	// }

	fmt.Println(result)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]Pipe

	scanner := bufio.NewScanner(file)
	var i, j, start_row, start_col int

	for scanner.Scan() {
		line := scanner.Text()
		var row []Pipe
		j = 0
		for _, r := range line {
			row = append(row, Pipe(r))
			if r == 'S' {
				start_row, start_col = i, j
			}
			j++
		}
		grid = append(grid, row)
		i++
	}

	explorer := NewExplorer(grid, start_row, start_col)

	explorer.explore()

	// for _, row := range explorer.grid {
	// 	fmt.Println(row)
	// }
	// fmt.Println()

	// for _, row := range explorer.distances {
	// 	fmt.Println(row)
	// }
	// fmt.Println()

	part1(explorer)
	part2(explorer)
}
