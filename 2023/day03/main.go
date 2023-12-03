package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func is_number(r rune) bool {
	return r >= '0' && r <= '9'
}

func is_symbol(r rune) bool {
	return r != '.' && !is_number(r)
}

type Iter struct {
	i, j       int
	schematic  [][]rune
	acc        []rune
	has_symbol bool
}

func (iter *Iter) has_next() bool {
	return iter.i < len(iter.schematic) && iter.j+1 < len(iter.schematic[iter.i]) ||
		iter.i+1 < len(iter.schematic)
}

func (iter *Iter) Next() (int, bool) {
	if !iter.has_next() {
		return 0, false
	}

	if iter.j+1 < len(iter.schematic[iter.i]) {
		iter.j += 1
	} else {
		iter.i += 1
		iter.j = 0
	}

	this := iter.schematic[iter.i][iter.j]
	if is_number(this) {
		iter.acc = append(iter.acc, this)

		// check if any neighbor is symbol
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 ||
					iter.i+i < 0 ||
					iter.i+i >= len(iter.schematic) ||
					iter.j+j < 0 ||
					iter.j+j >= len(iter.schematic[iter.i+i]) {
					continue
				}
				if is_symbol(iter.schematic[iter.i+i][iter.j+j]) {
					iter.has_symbol = true
				}
			}
		}
	} else if len(iter.acc) > 0 {
		num, err := strconv.Atoi(string(iter.acc))
		if err != nil {
			panic(err)
		}
		iter.acc = make([]rune, 0)

		if iter.has_symbol {
			// reset
			iter.has_symbol = false
			return num, true
		}
	}

	return iter.Next()
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var schematic [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	iter := Iter{0, -1, schematic, nil, false}
	result := 0
	for {
		num, has_next := iter.Next()
		result += num
		if !has_next {
			break
		}
	}

	fmt.Println(result)
}
