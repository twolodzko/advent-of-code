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
	i, j      int
	schematic [][]rune
	acc       []rune
	symbols   map[Symbol]bool
}

type Symbol struct {
	value rune
	i, j  int
}

func (iter *Iter) has_next() bool {
	return iter.i < len(iter.schematic) && iter.j+1 < len(iter.schematic[iter.i]) ||
		iter.i+1 < len(iter.schematic)
}

func (iter *Iter) Next() (int, map[Symbol]bool, bool) {
	if !iter.has_next() {
		return 0, nil, false
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
				r := iter.schematic[iter.i+i][iter.j+j]
				if is_symbol(r) {
					symbol := Symbol{r, iter.i + i, iter.j + j}
					iter.symbols[symbol] = true
				}
			}
		}
	} else if len(iter.acc) > 0 {
		num, err := strconv.Atoi(string(iter.acc))
		if err != nil {
			panic(err)
		}
		symbols := iter.symbols

		// reset
		iter.acc = make([]rune, 0)
		iter.symbols = make(map[Symbol]bool)

		return num, symbols, true
	}

	return iter.Next()
}

func part1(iter Iter) {
	result := 0
	for {
		num, symbols, has_next := iter.Next()
		if len(symbols) > 0 {
			result += num
		}
		if !has_next {
			break
		}
	}

	fmt.Println(result)
}

func part2(iter Iter) {
	gears := make(map[Symbol][]int)

	for {
		num, symbols, has_next := iter.Next()
		for symbol, _ := range symbols {
			if symbol.value == '*' {
				gears[symbol] = append(gears[symbol], num)
			}
		}
		if !has_next {
			break
		}
	}

	result := 0
	for _, nums := range gears {
		if len(nums) > 1 {
			ratio := 1
			for _, x := range nums {
				ratio *= x
			}
			result += ratio
		}
	}
	fmt.Println(result)
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

	iter := Iter{0, -1, schematic, nil, make(map[Symbol]bool)}
	part1(iter)
	part2(iter)
}
