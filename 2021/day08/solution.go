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

func simpleMatches(rows []Row) int {
	var lengths = map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	total := 0
	for _, row := range rows {
		for _, field := range row.output {
			if _, ok := lengths[len(field)]; ok {
				total += 1
			}
		}
	}
	return total
}

// ================================================

func uniqueWords(row Row) []string {
	words := make(map[string]bool)
	for _, w := range row.patterns {
		if _, ok := words[w]; !ok {
			words[w] = true
		}
	}
	for _, w := range row.output {
		if _, ok := words[w]; !ok {
			words[w] = true
		}
	}
	var out []string
	for w, _ := range words {
		out = append(out, w)
	}
	return out
}

type Counter struct {
	counts map[rune]int
}

func newCounter() *Counter {
	counts := make(map[rune]int)
	return &Counter{counts}
}

func (c *Counter) Add(v rune) {
	if _, ok := c.counts[v]; ok {
		c.counts[v] += 1
	} else {
		c.counts[v] = 1
	}
}

func countChars(words []string) *Counter {
	counter := newCounter()
	for _, word := range words {
		for _, r := range word {
			counter.Add(r)
		}
	}
	return counter
}

func referenceCounts() map[rune]int {
	reference := countChars([]string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdf",
	})
	return reference.counts
}

func lifeSupportRating(row Row) int {
	words := uniqueWords(row)
	fmt.Println(words)
	reference := referenceCounts()
	for k, v := range reference {
		fmt.Printf("%q => %d\n", k, v)
	}
	fmt.Println("=========")
	counter := countChars(words)
	for k, v := range counter.counts {
		fmt.Printf("%q => %d\n", k, v)
	}
	return -1
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

	row, _ := parse("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	lifeSupportRating(row)

	// result2 := lifeSupportRating(arr)
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
