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
	lengths := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	candidates := make(map[int]string)
	total := 0
	for _, field := range words {
		if digit, ok := lengths[len(field)]; ok {
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
	for w := range words {
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

func getMapping(row Row) map[int]string {

	words := uniqueWords(row)

	// 1, 4, 7, 8
	candidates, _ := obviousCandidates(words)

	//  aaaa
	// b    c
	// b    c
	//  dddd
	// e    f
	// e    f
	//  dddd

	for _, w := range words {
		if len(w) == 6 {
			// 0, 6, 9
			if c, ok := candidates[1]; ok && len(intersection(c, w)) == 1 {
				candidates[6] = w
			} else if c, ok := candidates[7]; ok && len(intersection(c, w)) == 2 {
				candidates[6] = w
			} else if c, ok := candidates[4]; ok && len(intersection(c, w)) == 4 {
				candidates[9] = w
			} else {
				candidates[0] = w
			}
		} else {
			// 2, 3, 5,
			if c, ok := candidates[1]; ok && len(intersection(c, w)) == 2 {
				candidates[3] = w
			} else if c, ok := candidates[7]; ok && len(intersection(c, w)) == 3 {
				candidates[3] = w
			} else if c, ok := candidates[4]; ok && len(intersection(c, w)) == 3 {
				candidates[5] = w
			} else {
				candidates[2] = w
			}
		}
	}
	return candidates
}

func decodeOutput(output []string, candidates map[int]string) []int {
	var decoded []int
	for _, field := range output {
		for digit, pattern := range candidates {
			if field == pattern {
				decoded = append(decoded, digit)
			}
		}
	}
	return decoded
}

func lifeSupportRating(rows []Row) int {
	total := 0
	for _, row := range rows {
		mapping := getMapping(row)
		decoded := decodeOutput(row.output, mapping)
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

	// row, _ := parse("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	// result2 := lifeSupportRating([]Row{row})
	// fmt.Println(result2)

	result2 := lifeSupportRating(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
