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
	grid [][]int
	prev []Path
}

func NewExplorer(grid [][]int) Explorer {
	return Explorer{grid, []Path{}}
}

func (e *Explorer) FindBest() int {
	var (
		right, down Path
		current     []Path
	)

	n := len(e.grid)
	k := len(e.grid[n-1])

	// initialize for row 0, the: x[0][0] -> x[0][1] -> x[0][2] -> ... path
	e.prev = []Path{{[]Point{{0, 0}}, 0}}
	for j := 1; j < k; j++ {
		path := append(e.prev[j-1].path, Point{0, j})
		risk := e.prev[j-1].risk + e.grid[0][j]
		e.prev = append(e.prev, Path{path, risk})
	}

	// other rows
	for i := 1; i < n; i++ {
		current = nil
		for j := 0; j < k; j++ {
			// path x[i-1][j] -> x[i][j]
			path := append(e.prev[j].path, Point{i - 1, j})
			risk := e.prev[j].risk + e.grid[i][j]
			down = Path{path, risk}

			// path x[i][j-1] -> x[i][j]
			if j > 0 {
				path := append(current[j-1].path, Point{i, j - 1})
				risk := current[j-1].risk + e.grid[i][j]
				right = Path{path, risk}
			} else {
				right = Path{nil, math.MaxInt}
			}

			// pick the path with lowest risk
			// collect current row in temporary slice
			if right.risk < down.risk {
				current = append(current, right)
			} else {
				current = append(current, down)
			}
		}
		// next row, we care only about the previous row
		e.prev = current
	}

	// the final x[n-1][k-1] position
	return current[len(current)-1].risk
}

func wrap(x int) int {
	return (x-1)%9 + 1
}

func printDims(grid [][]int) {
	n := len(grid)
	k := len(grid[n-1])
	fmt.Println(n, k)
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

	// for _, row := range grid {
	// 	for _, x := range row {
	// 		fmt.Print(x)
	// 	}
	// 	fmt.Println()
	// }

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

	// printDims(arr)
	// printDims(expandGrid(arr, 5))

	explorer := NewExplorer(arr)

	result1 := explorer.FindBest()
	fmt.Printf("Puzzle 1: %v\n", result1)

	biggerExplorer := NewExplorer(expandGrid(arr, 5))

	result2 := biggerExplorer.FindBest()
	fmt.Printf("Puzzle 2: %v\n", result2)
}
