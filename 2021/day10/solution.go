package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readFile(filename string) ([]string, error) {
	var arr []string

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
		if err != nil {
			return arr, err
		}
		arr = append(arr, line)
	}
	err = scanner.Err()
	return arr, err
}

type Stack struct {
	elem []rune
}

func (s *Stack) Push(x rune) {
	s.elem = append(s.elem, x)
}

func (s *Stack) Pop() rune {
	var x rune
	var last = len(s.elem) - 1
	x, s.elem = s.elem[last], s.elem[:last]
	return x
}

func (s *Stack) Last() (rune, bool) {
	if len(s.elem) == 0 {
		return rune(0), false
	}
	return s.elem[len(s.elem)-1], true
}

func isClosing(r rune) bool {
	switch r {
	case ')', ']', '}', '>':
		return true
	default:
		return false
	}
}

func checkLine(line string) int {
	var s Stack
	var scores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var matching = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	runes := []rune(line)
	s.Push(runes[0])
	for i := 1; i < len(runes); i++ {
		r := runes[i]
		if isClosing(r) {
			if last, ok := s.Last(); ok && matching[last] != r {
				return scores[r]
			}
			s.Pop()
		} else {
			s.Push(r)
		}
	}
	return 0
}

func scoreLines(lines []string) int {
	total := 0
	for _, line := range lines {
		total += checkLine(line)
	}
	return total
}

func correctLine(line string) int {
	var s Stack
	var scores = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	var matching = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	total := 0
	runes := []rune(line)
	s.Push(runes[0])
	for i := 1; i < len(runes); i++ {
		r := runes[i]
		if isClosing(r) {
			if last, ok := s.Last(); ok && matching[last] != r {
				log.Fatalf("unexpected character: %q (pos=%d)", r, i)
			}
			s.Pop()
		} else {
			s.Push(r)
		}
	}
	for len(s.elem) > 0 {
		r := s.Pop()
		total *= 5
		total += scores[r]
	}
	return total
}

func correctedLinesScores(lines []string) []int {
	var scores []int
	for _, line := range lines {
		if checkLine(line) > 0 {
			// skip
			continue
		}
		scores = append(scores, correctLine(line))
	}
	return scores
}

func middleScore(lines []string) int {
	scores := correctedLinesScores(lines)
	sort.Ints(scores)
	return scores[len(scores)/2]
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

	result1 := scoreLines(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := middleScore(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
