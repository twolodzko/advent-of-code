package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	dest Point
	risk int
}

type Paths []Path

func (p Paths) Less(i, j int) bool {
	return p[i].risk < p[j].risk
}

func (p Paths) Len() int {
	return len(p)
}

func (p Paths) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Paths) Pop() (Path, bool) {
	if len(*p) > 0 {
		path := (*p)[0]
		*p = (*p)[1:]
		return path, true
	}
	return Path{}, false
}

func (p Paths) MinRisk() int {
	best := math.MaxInt
	for _, path := range p {
		if path.risk < best {
			best = path.risk
		}
	}
	return best
}

type Explorer struct {
	risks [][]int
	paths [][]Path
	queue Paths
}

func (e *Explorer) Neighbors(p Point) []Point {
	n := len(e.risks)
	k := len(e.risks[n-1])
	var neighbors []Point
	if p.i > 0 {
		neighbors = append(neighbors, Point{p.i - 1, p.j})
	}
	if p.i+1 < n-1 {
		neighbors = append(neighbors, Point{p.i + 1, p.j})
	}
	if p.j > 0 {
		neighbors = append(neighbors, Point{p.i, p.j - 1})
	}
	if p.j+1 < k-1 {
		neighbors = append(neighbors, Point{p.i, p.j + 1})
	}
	return neighbors
}

func NewExplorer(grid [][]int) Explorer {
	var paths [][]Path
	n := len(grid)
	k := len(grid[n-1])
	for i := 0; i < n; i++ {
		row := []Path{}
		for j := 0; j < k; j++ {
			row = append(row, Path{Point{0, 0}, math.MaxInt})
		}
		paths = append(paths, row)
	}
	paths[0][0] = Path{Point{0, 0}, 0}
	queue := Paths{{Point{0, 1}, grid[0][1]}, {Point{1, 0}, grid[1][0]}}
	sort.Sort(queue)
	return Explorer{grid, paths, queue}
}

func (e *Explorer) FindBest() int {
	n := len(e.risks)
	k := len(e.risks[n-1])

	for len(e.queue) > 0 {
		here, ok := e.queue.Pop()
		// fmt.Println(here)
		if !ok {
			break
		}
		if here.risk < e.paths[here.dest.i][here.dest.j].risk {
			e.paths[here.dest.i][here.dest.j] = here
		}
		if here.dest.i == n-1 && here.dest.j == k-1 {
			if here.risk < e.queue.MinRisk() {
				break
			}
		}

		for _, neighbor := range e.Neighbors(here.dest) {
			to := e.NewPath(here.dest, neighbor)
			if to.risk < e.paths[to.dest.i][to.dest.j].risk {
				e.queue = append(e.queue, to)
			}
		}
		sort.Sort(e.queue)
	}
	return e.paths[n-1][k-1].risk
}

func (e *Explorer) NewPath(from, to Point) Path {
	risk := e.paths[from.i][from.j].risk + e.risks[to.i][to.j]
	return Path{to, risk}
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
