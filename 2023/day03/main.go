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
	cache     MaybePart
}

type MaybePart struct {
	acc     []rune
	symbols map[Symbol]bool
}

func NewMaybePart() MaybePart {
	return MaybePart{
		make([]rune, 0),
		make(map[Symbol]bool),
	}
}

func (p MaybePart) Unpack() (int, map[Symbol]bool, bool) {
	num, err := strconv.Atoi(string(p.acc))
	return num, p.symbols, len(p.symbols) > 0 && err == nil
}

type Symbol struct {
	value rune
	i, j  int
}

func (iter *Iter) has_next() bool {
	return iter.i < len(iter.schematic) && iter.j+1 < len(iter.schematic[iter.i]) ||
		iter.i+1 < len(iter.schematic)
}

func (iter *Iter) Next() (MaybePart, bool) {
	if !iter.has_next() {
		return iter.cache, false
	}

	if iter.j+1 < len(iter.schematic[iter.i]) {
		iter.j += 1
	} else {
		iter.i += 1
		iter.j = 0
	}

	this := iter.schematic[iter.i][iter.j]
	if is_number(this) {
		iter.cache.acc = append(iter.cache.acc, this)

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
					iter.cache.symbols[symbol] = true
				}
			}
		}
	} else if len(iter.cache.acc) > 0 {
		cache := iter.cache

		// reset
		iter.cache = NewMaybePart()

		return cache, true
	}

	return iter.Next()
}

func part1(iter Iter) {
	result := 0
	for {
		part, has_next := iter.Next()
		if num, _, ok := part.Unpack(); ok {
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
		part, has_next := iter.Next()

		if num, symbols, ok := part.Unpack(); ok {
			for symbol, _ := range symbols {
				if symbol.value == '*' {
					gears[symbol] = append(gears[symbol], num)
				}
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

	iter := Iter{0, -1, schematic, NewMaybePart()}
	part1(iter)
	part2(iter)
}
