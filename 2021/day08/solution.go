package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFile(filename string) ([]Row, error) {
	var arr []Row

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
		row, err := parse(line)
		if err != nil {
			return arr, err
		}
		arr = append(arr, row)
	}
	err = scanner.Err()
	return arr, err
}

type Row struct {
	patterns []string
	output   []string
}

func parse(row string) (Row, error) {
	fields := strings.Split(row, "|")
	if len(fields) != 2 {
		return Row{}, fmt.Errorf("wrong input: %v", row)
	}
	return Row{strings.Fields(fields[0]), strings.Fields(fields[1])}, nil
}

func findObviousCandidates(fields []string) []int {
	var lengths = map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	var candidates []int
	for _, field := range fields {
		if val, ok := lengths[len(field)]; ok {
			candidates = append(candidates, val)
		}
	}
	return candidates
}

func simpleMatches(rows []Row) int {
	total := 0
	for _, row := range rows {
		total += len(findObviousCandidates(row.output))
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

	result1 := simpleMatches(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	// row, _ := parse("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	// fmt.Println(row)

	// result2 := lifeSupportRating(arr)
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
