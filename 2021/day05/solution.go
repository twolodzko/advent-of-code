package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func newPoint(arr []int) (Point, error) {
	if len(arr) != 2 {
		return Point{}, fmt.Errorf("invalid input: %v", arr)
	}
	return Point{arr[0], arr[1]}, nil
}

func (p1 Point) Equal(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

type Vector struct {
	from, to Point
}

func newVector(arr []Point) (Vector, error) {
	if len(arr) != 2 {
		return Vector{}, fmt.Errorf("invalid input: %v", arr)
	}
	return Vector{arr[0], arr[1]}, nil
}

type Counter struct {
	counter map[Point]int
}

func newCounter() *Counter {
	counts := make(map[Point]int)
	return &Counter{counts}
}

func (c *Counter) add(p Point) {
	if _, ok := c.counter[p]; ok {
		c.counter[p] += 1
	} else {
		c.counter[p] = 1
	}
}

func (v Vector) interpolate(diag bool) []Point {
	var (
		dx, dy int
		point  Point
		points []Point
	)

	switch {
	case (v.to.x - v.from.x) > 0:
		dx = 1
	case (v.to.x - v.from.x) < 0:
		dx = -1
	default:
		dx = 0
	}

	switch {
	case (v.to.y - v.from.y) > 0:
		dy = 1
	case (v.to.y - v.from.y) < 0:
		dy = -1
	default:
		dy = 0
	}

	point = v.from

	if diag {
		points = append(points, point)
		for point != v.to {
			point.x += dx
			point.y += dy
			points = append(points, point)
		}
	} else {
		if v.from.x == v.to.x {
			points = append(points, point)
			for point.y != v.to.y {
				point.y += dy
				points = append(points, point)
			}
		} else if v.from.y == v.to.y {
			points = append(points, point)
			for point.x != v.to.x {
				point.x += dx
				points = append(points, point)
			}
		}
	}

	return points
}

func parseLine(str string) (Vector, error) {
	raw := strings.Split(str, "->")
	if len(raw) != 2 {
		return Vector{}, fmt.Errorf("invalid input: %v", str)
	}
	var vector []Point
	for _, field := range raw {
		points := strings.Split(field, ",")
		coords := []int{}
		for _, val := range points {
			i, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				return Vector{}, err
			}
			coords = append(coords, i)
		}
		point, err := newPoint(coords)
		if err != nil {
			return Vector{}, err
		}
		vector = append(vector, point)
	}
	return newVector(vector)
}

func readFile(filename string) ([]Vector, error) {
	var arr []Vector

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
		vec, err := parseLine(line)
		if err != nil {
			return arr, err
		}
		// fmt.Printf("%v", vec)
		// fmt.Printf("=> %v\n\n", vec.interpolate())
		arr = append(arr, vec)
	}
	err = scanner.Err()
	return arr, err
}

func overlaps(vectors []Vector, diag bool) int {
	counter := newCounter()
	for _, v := range vectors {
		points := v.interpolate(diag)
		for _, p := range points {
			counter.add(p)
		}
	}

	count := 0
	for _, val := range counter.counter {
		if val > 1 {
			count += 1
		}
	}
	return count
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

	result1 := overlaps(arr, false)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := overlaps(arr, true)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
