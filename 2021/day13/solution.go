package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]Point, []Fold, error) {
	var (
		points []Point
		folds  []Fold
	)

	file, err := os.Open(filename)
	if err != nil {
		return points, folds, err
	}
	defer file.Close()

	readingPoints := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return points, folds, err
		}

		if strings.TrimSpace(line) == "" {
			readingPoints = false
			continue
		}

		if readingPoints {
			fields := strings.Split(line, ",")
			if len(fields) != 2 {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			x, err := strconv.Atoi(fields[0])
			if err != nil {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			y, err := strconv.Atoi(fields[1])
			if err != nil {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			points = append(points, Point{x, y})
		} else {
			fields := strings.Fields(line)
			if len(fields) != 3 {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			fold := strings.Split(fields[2], "=")
			if len(fold) != 2 {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			pos, err := strconv.Atoi(fold[1])
			if err != nil {
				return points, folds, fmt.Errorf("invalid input: %v", line)
			}
			folds = append(folds, Fold{pos, fold[0] == "y"})
		}
	}
	err = scanner.Err()
	return points, folds, err
}

type Point struct {
	x, y int
}

type Fold struct {
	position   int
	horizontal bool
}

func (f *Fold) Fold(p Point) Point {
	if f.horizontal && p.y >= f.position {
		p.y = f.position + (f.position - p.y)
	} else if !f.horizontal && p.x >= f.position {
		p.x = f.position + (f.position - p.x)
	}
	return p
}

func countPoints(points []Point) int {
	uniquePoints := make(map[Point]bool)
	for _, point := range points {
		uniquePoints[point] = true
	}
	return len(uniquePoints)
}

func doFolding(points []Point, folds []Fold) int {
	var count int
	for fondId, fold := range folds {
		for i, point := range points {
			points[i] = fold.Fold(point)
		}
		if fondId == 0 {
			count = countPoints(points)
		}
	}
	return count
}

func print(points []Point) {
	maxX, maxY := 0, 0
	for _, point := range points {
		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
	}
	var matrix [][]int
	for i := 0; i <= maxY; i++ {
		row := []int{}
		for j := 0; j <= maxX; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	for _, point := range points {
		matrix[point.y][point.x] += 1
	}
	for _, row := range matrix {
		for _, x := range row {
			if x > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	points, folds, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1 := doFolding(points, folds)
	fmt.Printf("Puzzle 1: %v\n", result1)

	fmt.Println("Puzzle 2:")
	print(points)
}
