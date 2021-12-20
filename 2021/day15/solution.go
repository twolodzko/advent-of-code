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

func (p Path) Last() Point {
	return p.path[len(p.path)-1]
}

type PathsQueue []Path

func (p PathsQueue) Less(i, j int) bool {
	return p[i].risk < p[j].risk
}

func (p PathsQueue) Len() int {
	return len(p)
}

func (p PathsQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PathsQueue) Pop() (Path, bool) {
	if len(*p) > 0 {
		val := (*p)[0]
		*p = (*p)[1:]
		return val, true
	}
	return Path{}, false
}

func (p PathsQueue) MinRisk() int {
	best := math.MaxInt
	for _, path := range p {
		if path.risk < best {
			best = path.risk
		}
	}
	return best
}

func (p *PathsQueue) Insert(path Path) {
	for i := 0; i < len(*p); i++ {
		if path.risk < (*p)[i].risk {
			tmp := append((*p)[:i], path)
			*p = append(tmp, (*p)[i:]...)
			return
		}
	}
	*p = append(*p, path)
}

type Explorer struct {
	risks [][]int
	paths [][]Path
	queue PathsQueue
}

func (e Explorer) Neighbors(p Point) []Path {
	n := len(e.risks)
	k := len(e.risks[n-1])

	var points []Point
	// if p.i > 0 {
	// 	points = append(points, Point{p.i - 1, p.j})
	// }
	if p.i+1 < n {
		points = append(points, Point{p.i + 1, p.j})
	}
	// if p.j > 0 {
	// 	points = append(points, Point{p.i, p.j - 1})
	// }
	if p.j+1 < k {
		points = append(points, Point{p.i, p.j + 1})
	}

	var neighbors []Path
	for _, dest := range points {
		neighbors = append(neighbors, e.NewPath(p, dest))
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
			row = append(row, Path{nil, math.MaxInt})
		}
		paths = append(paths, row)
	}

	paths[0][0].risk = 0

	var queue PathsQueue
	for _, point := range []Point{{0, 1}, {1, 0}} {
		path := Path{[]Point{{0, 0}, point}, grid[point.i][point.j]}
		queue.Insert(path)
	}
	return Explorer{grid, paths, queue}
}

func (e *Explorer) FindBest() int {
	n := len(e.risks)
	k := len(e.risks[n-1])

	for {
		path, emptyQueue := e.queue.Pop()
		if !emptyQueue {
			break
		}
		here := path.Last()
		// fmt.Println(e.queue)
		// fmt.Println(path)
		// if i > 10 {
		// 	// fmt.Println(e.queue)
		// 	break
		// }
		if path.risk <= e.paths[here.i][here.j].risk {
			e.paths[here.i][here.j] = path
		} else {
			continue
		}
		for _, neighbor := range e.Neighbors(here) {
			e.queue.Insert(neighbor)
		}
		if here.i == n-1 && here.j == k-1 {
			if e.paths[here.i][here.j].risk <= e.queue.MinRisk() {
				break
			}
		}
	}
	fmt.Println(e.paths[n-1][k-1])
	return e.paths[n-1][k-1].risk
}

func (e *Explorer) NewPath(from, to Point) Path {
	risk := e.paths[from.i][from.j].risk + e.risks[to.i][to.j]
	path := append(e.paths[from.i][from.j].path, to)
	return Path{path, risk}
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
