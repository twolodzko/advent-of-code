package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

func (this Direction) String() string {
	switch this {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	default:
		return "West"
	}
}

const (
	North Direction = iota
	South
	West
	East
)

type Pipe struct {
	north, south, east, west bool
}

func (this Pipe) IsFinal() bool {
	return (this.north == this.south) && (this.south == this.east) && (this.east == this.west)
}

func (this Pipe) IsGround() bool {
	return !this.north && !this.south && !this.east && !this.west
}

func (this Pipe) String() string {
	switch {
	case this.north && this.south && !this.east && !this.west:
		return "|"
	case !this.north && !this.south && this.east && this.west:
		return "-"
	case this.north && !this.south && this.east && !this.west:
		return "L"
	case this.north && !this.south && !this.east && this.west:
		return "J"
	case !this.north && this.south && !this.east && this.west:
		return "7"
	case !this.north && this.south && this.east && !this.west:
		return "F"
	case this.north && this.south && this.east && this.west:
		return "S"
	default:
		return "."
	}
}

func parse(r rune) Pipe {
	switch r {
	case '|':
		return Pipe{true, true, false, false}
	case '-':
		return Pipe{false, false, true, true}
	case 'L':
		return Pipe{true, false, true, false}
	case 'J':
		return Pipe{true, false, false, true}
	case '7':
		return Pipe{false, true, false, true}
	case 'F':
		return Pipe{false, true, true, false}
	case '.':
		return Pipe{false, false, false, false}
	case 'S':
		return Pipe{true, true, true, true}
	default:
		panic(fmt.Sprintf("invalid symbol: %v", r))
	}
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
			// the direction where we came from is not available
			switch from {
			case North:
				if !this.north {
					return
				}
				this.north = false
			case South:
				if !this.south {
					return
				}
				this.south = false
			case East:
				if !this.east {
					return
				}
				this.east = false
			case West:
				if !this.west {
					return
				}
				this.west = false
			}

			if e.distances[i][j] == 0 {
				e.distances[i][j] = distance
			} else {
				e.distances[i][j] = min(e.distances[i][j], distance)
			}

			switch {
			case this.north:
				e.move(South, i-1, j, distance+1)
			case this.south:
				e.move(North, i+1, j, distance+1)
			case this.west:
				e.move(East, i, j-1, distance+1)
			case this.east:
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

	tmp := Matrix(len(explorer.grid), len(explorer.grid[0]))

	result := 0
	tmp_count := 0
	should_count := false
	for i, row := range explorer.grid {
		should_count = false
		tmp_count = 0
		for j, x := range row {
			if x.north || x.south {
				result += tmp_count
				tmp_count = 0
				should_count = !should_count
			} else if should_count && x.IsGround() {
				tmp_count++
				tmp[i][j] = 1
			}
		}
	}

	for _, row := range tmp {
		fmt.Println(row)
	}

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
			row = append(row, parse(r))
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
