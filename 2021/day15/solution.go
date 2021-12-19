package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}
		if err != nil {
			return arr, err
		}
		row := []int{}
		for _, r := range line {
			d, err := strconv.Atoi(string(r))
			if err != nil {
				return arr, err
			}
			row = append(row, d)
		}
		arr = append(arr, row)
	}
	err = scanner.Err()
	return arr, err
}

type Point struct {
	i, j int
}

type Path struct {
	path []Point
	risk int
}

type Explorer struct {
	grid  [][]int
	paths [][]Path
}

func (e *Explorer) Neighbors(p Point) []Point {
	// n := len(e.grid)
	// k := len(e.grid[n-1])
	var neighbors []Point
	if p.i > 0 {
		neighbors = append(neighbors, Point{p.i - 1, p.j})
	}
	// if p.i+1 < n-1 {
	// 	neighbors = append(neighbors, Point{p.i + 1, p.j})
	// }
	if p.j > 0 {
		neighbors = append(neighbors, Point{p.i, p.j - 1})
	}
	// if p.j+1 < k-1 {
	// 	neighbors = append(neighbors, Point{p.i, p.j + 1})
	// }
	return neighbors
}

func NewExplorer(grid [][]int) Explorer {
	var paths [][]Path
	n := len(grid)
	k := len(grid[n-1])
	for i := 0; i < n; i++ {
		row := []Path{}
		for j := 0; j < k; j++ {
			row = append(row, Path{nil, math.MaxInt})
		}
		paths = append(paths, row)
	}
	paths[0][0] = Path{[]Point{{0, 0}}, 0}
	return Explorer{grid, paths}
}

func (e *Explorer) FindBest() int {
	n := len(e.grid)
	k := len(e.grid[n-1])

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if e.paths[i][j].risk == 0 {
				continue
			}
			best := Path{nil, math.MaxInt}
			here := Point{i, j}
			for _, neighbor := range e.Neighbors(here) {
				path := e.NewPath(neighbor, here)
				if path.risk < best.risk {
					best = path
				}
			}
			e.paths[i][j] = best
		}
	}

	return e.paths[n-1][k-1].risk
}

func (e *Explorer) NewPath(from, to Point) Path {
	prev := e.paths[from.i][from.j]
	risk := prev.risk + e.grid[to.i][to.j]
	return Path{append(prev.path, to), risk}
}

func wrap(x int) int {
	return (x-1)%9 + 1
}

func expandGrid(grid [][]int, times int) [][]int {

	for i, row := range grid {
		for k := 1; k < times; k++ {
			for _, x := range row {
				grid[i] = append(grid[i], wrap(x+k))
			}
		}
	}

	n := len(grid)

	for k := 1; k < times; k++ {
		for i := 0; i < n; i++ {
			row := grid[i]
			tmp := []int{}
			for _, x := range row {
				tmp = append(tmp, wrap(x+k))
			}
			grid = append(grid, tmp)
		}
	}
	return grid
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

	explorer := NewExplorer(arr)

	result1 := explorer.FindBest()
	fmt.Printf("Puzzle 1: %v\n", result1)

	biggerExplorer := NewExplorer(expandGrid(arr, 5))

	result2 := biggerExplorer.FindBest()
	fmt.Printf("Puzzle 2: %v\n", result2)
}
