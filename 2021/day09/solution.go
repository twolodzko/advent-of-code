package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readFile(filename string) (Heightmap, error) {
	var arr [][]int

	file, err := os.Open(filename)
	if err != nil {
		return Heightmap{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, ch := range line {
			i, err := strconv.Atoi(string(ch))
			if err != nil {
				return Heightmap{}, err
			}
			row = append(row, i)
		}
		arr = append(arr, row)
	}
	err = scanner.Err()
	return Heightmap{arr}, err
}

type Heightmap struct {
	data [][]int
}

type Position struct {
	i, j int
}

func (h *Heightmap) Neighbors(i, j int) []Position {
	var positions []Position
	if i > 0 {
		positions = append(positions, Position{i - 1, j})
	}
	if i+1 < len(h.data) {
		positions = append(positions, Position{i + 1, j})
	}
	if j > 0 {
		positions = append(positions, Position{i, j - 1})
	}
	if j+1 < len(h.data[i]) {
		positions = append(positions, Position{i, j + 1})
	}
	return positions
}

func (h *Heightmap) isSmallest(i, j int) bool {
	for _, n := range h.Neighbors(i, j) {
		if h.data[i][j] >= h.data[n.i][n.j] {
			return false
		}
	}
	return true
}

type Point struct {
	value int
	Position
}

func (h *Heightmap) lowPoints() []Point {
	var points []Point
	for i := 0; i < len(h.data); i++ {
		for j := 0; j < len(h.data[i]); j++ {
			if h.isSmallest(i, j) {
				point := Point{h.data[i][j], Position{i, j}}
				points = append(points, point)
			}
		}
	}
	return points
}

func (h *Heightmap) riskLevel() int {
	lp := h.lowPoints()
	level := 0
	for _, point := range lp {
		level += point.value + 1
	}
	return level
}

type Explorer struct {
	*Heightmap
	visited [][]bool
}

func newExplorer(h *Heightmap) *Explorer {
	var v [][]bool
	for i := 0; i < len(h.data); i++ {
		v = append(v, []bool{})
		for j := 0; j < len(h.data[i]); j++ {
			v[i] = append(v[i], false)
		}
	}
	return &Explorer{h, v}
}

func (e *Explorer) inBasin(i, j int) bool {
	if e.visited[i][j] || e.data[i][j] == 9 {
		return false
	}
	for _, n := range e.Neighbors(i, j) {
		if e.visited[n.i][n.j] {
			continue
		}
		// not sure why here > works instead of >=
		if e.data[i][j] > e.data[n.i][n.j] {
			return false
		}
	}
	return true
}

func (e *Explorer) exploreNeighbors(i, j int) []Point {
	var points []Point
	for _, n := range e.Neighbors(i, j) {
		if e.inBasin(n.i, n.j) {
			p := Point{e.data[n.i][n.j], Position{n.i, n.j}}
			points = append(points, p)
			e.visited[n.i][n.j] = true
			points = append(points, e.exploreNeighbors(n.i, n.j)...)
		}
	}
	return points
}

func (e *Explorer) findBasins() [][]Point {
	var (
		basin  []Point
		basins [][]Point
	)
	lp := e.lowPoints()
	for _, point := range lp {
		basin = []Point{point}
		e.visited[point.i][point.j] = true
		basin = append(basin, e.exploreNeighbors(point.i, point.j)...)
		basins = append(basins, basin)
	}
	return basins
}

func basinSizesMultiplied(h *Heightmap) int {
	explorer := newExplorer(h)
	basins := explorer.findBasins()

	var sizes []int
	for _, basin := range basins {
		sizes = append(sizes, len(basin))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	mul := 1
	for i := 0; i < 3; i++ {
		mul *= sizes[i]
	}
	return mul
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	heightmap, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1 := heightmap.riskLevel()
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := basinSizesMultiplied(&heightmap)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
