package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	forward = iota
	down
	up
)

type Course struct {
	direction int
	distance  int
}

func newCourse(line string) (Course, error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return Course{}, fmt.Errorf("Wrong input %v", fields)
	}
	var (
		direction int
		distance  int
		err       error
	)
	switch fields[0] {
	case "forward":
		direction = forward
	case "up":
		direction = up
	case "down":
		direction = down
	default:
		err = fmt.Errorf("Invalid direction: %v", fields[0])
	}
	if err != nil {
		return Course{}, err
	}

	distance, err = strconv.Atoi(fields[1])
	return Course{direction, distance}, err
}

func readFile(filename string) ([]Course, error) {
	var arr []Course

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
		c, err := newCourse(line)
		if err != nil {
			return arr, err
		}
		arr = append(arr, c)
	}
	err = scanner.Err()
	return arr, err
}

func totalDistance(arr []Course) int {
	var (
		horizontal int = 0
		depth      int = 0
	)
	for _, c := range arr {
		switch c.direction {
		case down:
			depth += c.distance
		case up:
			depth -= c.distance
		default:
			horizontal += c.distance
		}
	}
	return horizontal * depth
}

func totalCourse(arr []Course) int {
	var (
		horizontal int = 0
		depth      int = 0
		aim        int = 0
	)
	for _, c := range arr {
		switch c.direction {
		case down:
			aim += c.distance
		case up:
			aim -= c.distance
		default:
			horizontal += c.distance
			depth += c.distance * aim
		}
	}
	return horizontal * depth
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

	result1 := totalDistance(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := totalCourse(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
