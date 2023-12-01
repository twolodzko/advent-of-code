package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extractNumbers(line string) int {
	var x, y int
	p := &x
	for _, r := range line {
		if r >= '0' && r <= '9' {
			*p = int(r - 48)
			y = *p
			p = &y
		}
	}
	return x*10 + y
}

func extractNumbers2(line string) int {
	var digits = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	done := false
	var x, y, tmp int
	p := &x
	for i := 0; i < len(line); i++ {
		r := []rune(line)[i]
		if r >= '0' && r <= '9' {
			tmp = int(r - 48)
			done = true
		} else {
			for n, d := range digits {
				if strings.HasPrefix(line[i:], d) {
					tmp = n + 1
					done = true
					break
				}
			}
		}
		if done {
			*p = tmp
			y = *p
			p = &y
			done = false
		}
	}
	return x*10 + y
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num := extractNumbers2(scanner.Text())
		fmt.Println(num)
		result += num
	}
	fmt.Println(result)
}
