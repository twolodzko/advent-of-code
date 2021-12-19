package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([][]Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		arr    [][]Point
		coords []Point
	)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}

		if strings.TrimSpace(line) == "" {
			arr = append(arr, coords)
			coords = nil
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}

		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			return nil, fmt.Errorf("invalid input: %v", line)
		}
		point := [3]int{}
		for i, field := range fields {
			x, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			point[i] = x
		}
		coords = append(coords, Point{point[0], point[1], point[2]})
	}
	arr = append(arr, coords)

	err = scanner.Err()
	return arr, err
}

type Point struct {
	x, y, z int
}

func square(x int) int {
	return x * x
}

func (a *Point) Dist(b Point) int {
	dist := square(a.x-b.x) + square(a.y-b.y) + square(a.z-b.z)
	return dist
}

type TriangMatrix [][]int

func Zeros(n int) TriangMatrix {
	var (
		row []int
		mtx [][]int
	)
	for i := 0; i < n; i++ {
		row = nil
		for j := 0; j < i; j++ {
			row = append(row, 0)
		}
		mtx = append(mtx, row)
	}
	return mtx
}

func NewDistMatrix(points []Point) TriangMatrix {
	dist := Zeros(len(points))
	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			dist[i][j] = points[i].Dist(points[j])
			// dist[j][i] = dist[i][j]
		}
	}
	return dist
}

func (m *TriangMatrix) Add(val int) {
	n := len(*m)
	row := []int{}
	for j := 0; j < n; j++ {
		row = append(row, 0)
	}
	row = append(row, val)
	*m = append(*m, row)
}

type Position struct {
	i, j int
}

func (m TriangMatrix) Find(val int) (Position, bool) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			if m[i][j] == val {
				return Position{i, j}, true
			}
		}
	}
	return Position{}, false
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

	dist := NewDistMatrix(arr[0])
	fmt.Println(dist)

	// - No matter of coordinates origin and rotation, the distances between points would be the same
	// - Given two points on two distance matrices, we can match the coordinates i1,j1 -> i2,j2
	// - We could collect all the points into single distance matrix between all the points
	// - The dimensions of the matrix would let us to identify unique points

	// fmt.Println(len(dist))

	// result1 := packet.VersionNumberSum()
	// fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := packet.Value()
	// fmt.Printf("Puzzle 2: %v\n", result2)

}
