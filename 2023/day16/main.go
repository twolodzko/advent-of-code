package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile rune

const (
	Empty              Tile = '.'
	LeftMirror         Tile = '/'
	RightMirror        Tile = '\\'
	HorizontalSplitter Tile = '-'
	VerticalSplitter   Tile = '|'
)

func (this Tile) String() string {
	return string(this)
}

type From struct {
	i, j int
	Direction
}

func (this From) Turn(direction Direction) From {
	return From{this.i, this.j, direction}
}

type BeamPaths struct {
	grid      [][]Tile
	energized [][]bool
	memory    map[From]bool
}

func (this BeamPaths) Next(from From) []From {
	var directions []From

	switch from.Direction {
	case Left:
		from.j++
		if from.j >= len(this.grid[0]) {
			return nil
		}
		switch this.grid[from.i][from.j] {
		case Empty, HorizontalSplitter:
			directions = append(directions, from)
		case RightMirror:
			directions = append(directions, from.Turn(Down))
		case LeftMirror:
			directions = append(directions, from.Turn(Up))
		case VerticalSplitter:
			directions = append(directions, from.Turn(Up))
			directions = append(directions, from.Turn(Down))
		}
	case Right:
		from.j--
		if from.j < 0 {
			return nil
		}
		switch this.grid[from.i][from.j] {
		case Empty, HorizontalSplitter:
			directions = append(directions, from)
		case RightMirror:
			directions = append(directions, from.Turn(Up))
		case LeftMirror:
			directions = append(directions, from.Turn(Down))
		case VerticalSplitter:
			directions = append(directions, from.Turn(Up))
			directions = append(directions, from.Turn(Down))
		}
	case Up:
		from.i--
		if from.i < 0 {
			return nil
		}
		switch this.grid[from.i][from.j] {
		case Empty, VerticalSplitter:
			directions = append(directions, from)
		case RightMirror:
			directions = append(directions, from.Turn(Right))
		case LeftMirror:
			directions = append(directions, from.Turn(Left))
		case HorizontalSplitter:
			directions = append(directions, from.Turn(Right))
			directions = append(directions, from.Turn(Left))
		}
	case Down:
		from.i++
		if from.i >= len(this.grid) {
			return nil
		}
		switch this.grid[from.i][from.j] {
		case Empty, VerticalSplitter:
			directions = append(directions, from)
		case RightMirror:
			directions = append(directions, from.Turn(Left))
		case LeftMirror:
			directions = append(directions, from.Turn(Right))
		case HorizontalSplitter:
			directions = append(directions, from.Turn(Right))
			directions = append(directions, from.Turn(Left))
		}
	}

	return directions
}

func (this BeamPaths) Explore(from From) {
	if _, ok := this.memory[from]; ok {
		// so we do not hit infinite loop
		return
	} else {
		this.memory[from] = true
	}

	for _, from := range this.Next(from) {
		this.energized[from.i][from.j] = true
		this.Explore(from)
	}
}

func Simulate(grid [][]Tile, from From) [][]bool {
	var energized [][]bool
	for i := 0; i < len(grid); i++ {
		energized = append(energized, make([]bool, len(grid[0])))
	}
	memory := make(map[From]bool)
	paths := BeamPaths{grid, energized, memory}
	paths.Explore(from)
	return paths.energized
}

type Direction rune

const (
	Left  Direction = '>'
	Right Direction = '<'
	Up    Direction = '^'
	Down  Direction = 'v'
)

func (this Direction) String() string {
	return string(this)
}

func parse(scanner *bufio.Scanner) [][]Tile {
	var grid [][]Tile
	for scanner.Scan() {
		var row []Tile
		for _, r := range scanner.Text() {
			row = append(row, Tile(r))
		}
		grid = append(grid, row)
	}
	return grid
}

func count(energized [][]bool) int {
	count := 0
	for _, row := range energized {
		for _, x := range row {
			if x {
				count++
			}
		}
	}
	return count
}

func part1(grid [][]Tile) {
	energized := Simulate(grid, From{0, -1, Left})
	fmt.Println(count(energized))
}

func part2(grid [][]Tile) {
	best := 0
	for i := 0; i < len(grid); i++ {
		for _, from := range []From{
			{i, -1, Left},
			{i, len(grid[0]), Right},
		} {
			energized := Simulate(grid, from)
			best = max(best, count(energized))
		}
	}
	for j := 0; j < len(grid[0]); j++ {
		for _, from := range []From{
			{-1, j, Down},
			{len(grid), j, Up},
		} {
			energized := Simulate(grid, from)
			best = max(best, count(energized))
		}
	}
	fmt.Println(best)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := parse(scanner)

	part1(grid)
	part2(grid)
}
