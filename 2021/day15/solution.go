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

func (x Point) Equal(y Point) bool {
	return x.i == y.i && x.j == y.j
}

type Path struct {
	path []Point
	risk int
}

type Explorer struct {
	graph [][]int
	paths [][]Path
}

func NewExplorer(graph [][]int) Explorer {
	var paths [][]Path
	n := len(graph)
	k := len(graph[n-1])
	for i := 0; i < n; i++ {
		row := []Path{}
		for j := 0; j < k; j++ {
			row = append(row, Path{nil, math.MaxInt})
		}
		paths = append(paths, row)
	}
	paths[0][0] = Path{[]Point{{0, 0}}, 0}
	return Explorer{graph, paths}
}

func (e *Explorer) FindBest() int {
	right := Path{nil, math.MaxInt}
	down := Path{nil, math.MaxInt}

	n := len(e.graph)
	k := len(e.graph[n-1])

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if e.paths[i][j].risk == 0 {
				continue
			}
			if i > 0 {
				right = e.NewPath(Point{i - 1, j}, Point{i, j})
			}
			if j > 0 {
				down = e.NewPath(Point{i, j - 1}, Point{i, j})
			}
			if right.risk < down.risk {
				e.paths[i][j] = right
			} else {
				e.paths[i][j] = down
			}
		}
	}

	return e.paths[n-1][k-1].risk
}

func (e *Explorer) NewPath(from, to Point) Path {
	prev := e.paths[from.i][from.j]
	risk := prev.risk + e.graph[to.i][to.j]
	return Path{append(prev.path, to), risk}
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

	// result2 := middleScore(arr)
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
