package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, error) {
	var arr []int

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	if err != nil {
		return arr, err
	}
	for _, val := range strings.Split(line, ",") {
		i, err := strconv.Atoi(val)
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}
	err = scanner.Err()
	return arr, err
}

type Lanternfish struct {
	timer int
}

func newLanternfish() *Lanternfish {
	return &Lanternfish{8}
}

func (l *Lanternfish) breed() (*Lanternfish, bool) {
	if l.timer == 0 {
		l.timer = 6
		return newLanternfish(), true
	}
	return nil, false
}

func simulate(arr []int, days int) int {
	var lanternfishes []*Lanternfish
	for _, i := range arr {
		lanternfishes = append(lanternfishes, &Lanternfish{i})
	}

	for t := 0; t < days; t++ {
		for _, l := range lanternfishes {
			if new, ok := l.breed(); ok {
				lanternfishes = append(lanternfishes, new)
			} else {
				l.timer -= 1
			}
		}
	}

	return len(lanternfishes)
}

type LanternfishSimulator struct {
	memory map[int]int
	days   int
}

func newLanternfishSimulator(days int) *LanternfishSimulator {
	memory := make(map[int]int)
	return &LanternfishSimulator{memory, days}
}

func (s *LanternfishSimulator) next(day int) int {
	// read from memory to speed up
	if v, ok := s.memory[day]; ok {
		return v
	}

	if day > s.days {
		// final node
		return 1
	}

	forward := s.next(day + 7)
	breed := s.next(day + 9)
	result := forward + breed

	// memoize
	s.memory[day] = result

	return result
}

func calculate(arr []int, days int) int {
	total := 0
	sim := newLanternfishSimulator(days)

	for _, t := range arr {
		start := t + 1
		total += sim.next(start)
	}
	return total
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

	result1 := simulate(arr, 80)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := calculate(arr, 256)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
