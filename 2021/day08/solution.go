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

type Digit struct {
	value    int
	segments []rune
}

func newDigit(digit int, str string) *Digit {
	runes := []rune(str)
	return &Digit{digit, runes}
}

func (d *Digit) numSegments() int {
	return len(d.segments)
}

func (d *Digit) isUsed(segment rune) bool {
	for _, r := range d.segments {
		if segment == r {
			return true
		}
	}
	return false
}

func (d *Digit) String() string {
	return fmt.Sprintf("Digit(%d)", d.value)
}

func (d *Digit) print() {
	var str, display string
	str += fmt.Sprintf("\n%d:\n", d.value)
	for _, segment := range "abcdefg" {
		display = string(segment)
		if !d.isUsed(segment) {
			display = " "
		}
		switch segment {
		case 'a', 'd', 'g':
			str += fmt.Sprintf(" %s \n", strings.Repeat(display, 2))
		case 'b', 'e':
			str += fmt.Sprintf("%s  ", display)
		default:
			str += fmt.Sprintf("%s\n", display)
		}
	}
	fmt.Println(str)
}

type Digits struct {
	digits []*Digit
}

func newDigits() *Digits {
	digits := []*Digit{
		newDigit(0, "abcefg"),
		newDigit(1, "cf"),
		newDigit(2, "acdeg"),
		newDigit(3, "acdfg"),
		newDigit(4, "bcdf"),
		newDigit(5, "abdfg"),
		newDigit(6, "abdefg"),
		newDigit(7, "acf"),
		newDigit(8, "abcdefg"),
		newDigit(9, "abcdfg"),
	}
	return &Digits{digits}
}

func (d *Digits) print() {
	for _, d := range d.digits {
		fmt.Printf("%s\n", d)
	}
}

func (d *Digits) lengths() map[int][]*Digit {
	lengths := make(map[int][]*Digit)
	for _, d := range d.digits {
		n := d.numSegments()
		if _, ok := lengths[n]; !ok {
			lengths[n] = []*Digit{d}
		} else {
			lengths[n] = append(lengths[n], d)
		}
	}
	return lengths
}

type Candidates struct {
	pattern string
	digits  []*Digit
}

func findObviousCandidates(fields []string) []Candidates {
	digits := newDigits()
	lengths := digits.lengths()
	var discovered []Candidates
	for _, field := range fields {
		n := len(field)
		matched := lengths[n]
		if len(matched) == 1 {
			discovered = append(discovered, Candidates{field, matched})
		}
	}
	return discovered
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
