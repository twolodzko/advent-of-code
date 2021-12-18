package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}
		arr = append(arr, line)
	}
	err = scanner.Err()
	return arr, err
}

// func Split(num int) Pair {
// 	left := int(math.Floor(float64(num) / 2))
// 	right := int(math.Ceil(float64(num) / 2))
// 	return Pair{left, right}
// }

const (
	START int = -1
	END       = -2
)

func isSeparator(r rune) bool {
	switch r {
	case '[', ',', ']':
		return true
	default:
		return false
	}
}

type Pair []int

func (p Pair) String() string {
	var s string
	for i, v := range p {
		switch v {
		case START:
			if i > 0 && p[i-1] != START {
				s += ","
			}
			s += "["
		case END:
			s += "]"
		default:
			if i > 0 && p[i-1] != START {
				s += ","
			}
			s += fmt.Sprintf("%v", v)
		}
	}
	return s
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func parse(str string) (Pair, error) {
	var pattern []int
	runes := []rune(str)
	i := 0
	for i < len(runes) {
		switch runes[i] {
		case '[':
			pattern = append(pattern, START)
		case ']':
			pattern = append(pattern, END)
		case ',':
		default:
			// number
			var num []rune
			for j := i; j < len(runes); j++ {
				if isSeparator(runes[j]) {
					break
				}
				num = append(num, runes[j])
			}
			i += max(len(num)-1, 0)
			val, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}
			pattern = append(pattern, val)
		}
		i++
	}
	return pattern, nil
}

func reduce(p Pair) Pair {
	depth := 0
	i := 0
	for i < len(p) {
		fmt.Println(p)
		if p[i] == START {
			depth++
		}
		if p[i] == END {
			depth--
		}
		if depth == 4 {
			p = explode(p)
			i = 0
			depth = 0
		}
		if p[i] >= 10 {
			p = split(p)
			i = 0
			depth = 0
		}
		i++
	}
	return p
}

func explode(p Pair) Pair {
	for i := 0; i < len(p); i++ {
		if p[i] > 0 && i+1 < len(p) && p[i+1] > 0 {
			// leftmost pair
			for j := 0; j < i; j++ {
				if p[j] > 0 {
					p[j] += p[i]
				}
			}
			for j := i + 2; j < len(p); j++ {
				if p[j] > 0 {
					p[j] += p[i+1]
				}
			}
			head := append(p[:i-1], 0)
			return append(head, p[i+3:]...)
		}
	}
	return p
}

func split(p Pair) Pair {
	for i := 0; i < len(p); i++ {
		if p[i] > 0 {
			// leftmost regular number
			x := float64(p[i]) / 2
			left := int(math.Floor(x))
			right := int(math.Ceil(x))
			head := append(p[:i], []int{START, left, right, END}...)
			return append(head, p[i+1:]...)
		}
	}
	return p
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	lines, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		pair, err := parse(line)
		if err != nil {
			log.Fatal(err)
		}
		if pair.String() != line {
			log.Fatalf("Parsing error: %v != %v", pair, line)
		}
		fmt.Println(pair)
		fmt.Println(reduce(pair))
		fmt.Println()
	}

	// result1 := packet.VersionNumberSum()
	// fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := packet.Value()
	// fmt.Printf("Puzzle 2: %v\n", result2)

}
