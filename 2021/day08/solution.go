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

func obviousCandidates(words []string) (map[int]string, int) {
	var lengths = map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	candidates := make(map[int]string)
	total := 0
	for _, field := range words {
		n := len(field)
		if digit, ok := lengths[n]; ok {
			candidates[digit] = field
			total += 1
		}
	}
	return candidates, total
}

func simpleMatches(rows []Row) int {
	total := 0
	for _, row := range rows {
		_, count := obviousCandidates(row.output)
		total += count
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

func remove(arr []string, key string) []string {
	for i, x := range arr {
		if x == key {
			switch i {
			case 0:
				return arr[1:]
			case len(arr) - 1:
				return arr[:len(arr)-1]
			default:
				return append(arr[:i], arr[i+1:]...)
			}
		}
	}
	return arr
}

func lifeSupportRating(row Row) int {
	words := uniqueWords(row)

	// 1, 4, 7, 8
	candidates, _ := obviousCandidates(words)
	for _, c := range candidates {
		words = remove(words, c)
	}

	fmt.Println(candidates)
	fmt.Println(words)

	// 9
	for i := 0; i < len(words); i++ {
		w := words[i]
		if len(w) != 5 {
			continue
		}
		c, ok := candidates[4]
		if ok && len(intersection(c, w)) == 4 {
			candidates[9] = w
			words = remove(words, w)
		}
	}

	// 6
	for i := 0; i < len(words); i++ {
		w := words[i]
		if len(w) != 6 {
			continue
		}
		c, ok := candidates[1]
		if ok && len(intersection(c, w)) == 1 {
			candidates[6] = w
			words = remove(words, w)
		} else {
			c, ok := candidates[4]
			if ok && len(intersection(c, w)) == 3 {
				candidates[6] = w
				words = remove(words, w)
			}
		}
	}

	// 0
	for i := 0; i < len(words); i++ {
		w := words[i]
		if len(w) != 6 {
			continue
		}
		c, ok := candidates[8]
		if ok && len(intersection(c, w)) == 6 {
			c, ok := candidates[1]
			if ok && len(intersection(c, w)) == 2 {
				candidates[0] = w
				words = remove(words, w)
			}
		} else {
			c, ok := candidates[4]
			if ok && len(intersection(c, w)) == 3 {
				candidates[0] = w
				words = remove(words, w)
			}
		}
	}

	// 3
	for i := 0; i < len(words); i++ {
		w := words[i]
		if len(w) != 5 {
			continue
		}
		c, ok := candidates[1]
		if ok && len(intersection(c, w)) == 2 {
			candidates[3] = w
			words = remove(words, w)
		}
	}

	// 5
	for i := 0; i < len(words); i++ {
		w := words[i]
		if len(w) != 5 {
			continue
		}
		c, ok := candidates[4]
		if ok && len(intersection(c, w)) == 3 {
			candidates[5] = w
			words = remove(words, w)
		}
	}

	fmt.Println(candidates)

	if len(words) > 1 {
		log.Fatalf("invalid input: %v", words)
	}

	candidates[2] = words[0]

	fmt.Println(candidates)
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
