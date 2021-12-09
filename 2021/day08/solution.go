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

func guessObviousDigit(s string) (int, bool) {
	switch len(s) {
	case 2:
		return 1, true
	case 4:
		return 4, true
	case 3:
		return 7, true
	case 7:
		return 8, true
	default:
		return 0, false
	}
}

func simpleMatches(rows []Row) int {
	total := 0
	for _, row := range rows {
		for _, field := range row.output {
			if _, ok := guessObviousDigit(field); ok {
				total += 1
			}
		}
	}
	return total
}

func (row *Row) words() []string {
	var words []string
	for _, w := range row.patterns {
		words = append(words, w)
	}
	for _, w := range row.output {
		words = append(words, w)
	}
	return words
}

func includes(s string, v rune) bool {
	for _, r := range s {
		if r == v {
			return true
		}
	}
	return false
}

func intersection(a, b string) []rune {
	var common []rune
	for _, r := range b {
		if includes(a, r) {
			common = append(common, r)
		}
	}
	return common
}

type GuessedDigits struct {
	guessed map[int]string
}

func newGuessedDigits() *GuessedDigits {
	guessed := make(map[int]string)
	return &GuessedDigits{guessed}
}

func (g *GuessedDigits) Init(words []string) {
	for _, w := range words {
		if d, ok := guessObviousDigit(w); ok {
			g.guessed[d] = w
		}
	}
}

func (g *GuessedDigits) guessDigit(s string) (int, bool) {
	if d, ok := guessObviousDigit(s); ok {
		return d, true
	}

	// 0, 6, 9
	if len(s) == 6 {
		// 9
		if p, ok := g.guessed[4]; ok && len(intersection(p, s)) == 4 {
			return 9, true
		}
		// 6
		if p, ok := g.guessed[1]; ok && len(intersection(p, s)) == 1 {
			return 6, true
		}
		if p, ok := g.guessed[7]; ok && len(intersection(p, s)) == 2 {
			return 6, true
		}
		// 0
		return 0, true
	}

	// 2, 3, 5
	if len(s) == 5 {
		// 3
		if p, ok := g.guessed[1]; ok && len(intersection(p, s)) == 2 {
			return 3, true
		}
		if p, ok := g.guessed[7]; ok && len(intersection(p, s)) == 3 {
			return 3, true
		}
		// 5
		if p, ok := g.guessed[4]; ok && len(intersection(p, s)) == 3 {
			return 5, true
		}
		// 2
		return 2, true
	}

	return 0, false
}

func decode(row Row) ([]int, error) {
	guess := newGuessedDigits()
	guess.Init(row.words())
	var digits []int
	for _, w := range row.output {
		if d, ok := guess.guessDigit(w); ok {
			digits = append(digits, d)
		} else {
			return nil, fmt.Errorf("unrecognized pattern: %v", w)
		}
	}
	return digits, nil
}

func lifeSupportRating(rows []Row) int {
	total := 0
	for _, row := range rows {
		decoded, err := decode(row)
		if err != nil {
			log.Fatal(err)
		}

		dec := 1
		for i := 0; i < len(decoded); i++ {
			dec *= 10
		}
		for _, d := range decoded {
			dec /= 10
			total += d * dec
		}
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

	result2 := lifeSupportRating(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
