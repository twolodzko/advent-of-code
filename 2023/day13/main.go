package main

import (
	"bufio"
	"fmt"
	"os"
)

type Terrain rune

const (
	Ash  Terrain = '.'
	Rock Terrain = '#'
)

func (this Terrain) String() string {
	return string(this)
}

func (this Terrain) Flip() Terrain {
	if this == Ash {
		return Rock
	} else {
		return Ash
	}
}

type Pattern [][]Terrain

func (this Pattern) LeftRightEqual(left, right int) bool {
	for i := 0; i < len(this); i++ {
		if this[i][left] != this[i][right] {
			return false
		}
	}
	return true
}

func (this Pattern) IsLeftRightMirror(left, right int) bool {
	for {
		if left < 0 || right >= len(this[0]) {
			break
		}
		if !this.LeftRightEqual(left, right) {
			return false
		}
		left--
		right++
	}
	return true
}

func (this Pattern) LeftRightReflection() (int, bool) {
	for start := 0; start < len(this[0])-1; start++ {
		left := start
		right := start + 1
		if this.IsLeftRightMirror(left, right) {
			return left + 1, true
		}
	}
	return 0, false
}

func (this Pattern) LeftRightFind(diff int) (int, bool) {
	for start := 0; start < len(this[0])-1; start++ {
		left := start
		right := start + 1
		count := 0
		for {
			if left < 0 || right >= len(this[0]) || count > diff {
				break
			}
			for i := 0; i < len(this); i++ {
				if this[i][left] != this[i][right] {
					count++
				}
				if count > diff {
					break
				}
			}
			left--
			right++
		}
		if count == diff {
			return start + 1, true
		}
	}
	return 0, false
}

func (this Pattern) TopDownEqual(top, down int) bool {
	for j := 0; j < len(this[top]); j++ {
		if this[top][j] != this[down][j] {
			return false
		}
	}
	return true
}

func (this Pattern) IsTopDownMirror(top, down int) bool {
	for {
		if top < 0 || down >= len(this) {
			break
		}
		if !this.TopDownEqual(top, down) {
			return false
		}
		top--
		down++
	}
	return true
}

func (this Pattern) TopDownReflection() (int, bool) {
	for start := 0; start < len(this)-1; start++ {
		top := start
		down := start + 1
		if this.IsTopDownMirror(top, down) {
			return top + 1, true
		}
	}
	return 0, false
}

func (this Pattern) TopDownFind(diff int) (int, bool) {
	for start := 0; start < len(this)-1; start++ {
		top := start
		down := start + 1
		count := 0
		for {
			if top < 0 || down >= len(this) || count > diff {
				break
			}
			for j := 0; j < len(this[top]); j++ {
				if this[top][j] != this[down][j] {
					count++
				}
				if count > diff {
					break
				}
			}
			top--
			down++
		}
		if count == diff {
			return start + 1, true
		}
	}
	return 0, false
}

// func (this Pattern) FixSmudge() (int, int) {
// 	vertical, _ := this.LeftRightReflection()
// 	horizontal, _ := this.TopDownReflection()

// 	for i := 0; i < len(this); i++ {
// 		for j := 0; j < len(this[0]); j++ {
// 			// flit the pattern
// 			this[i][j] = this[i][j].Flip()

// 			if num, ok := this.LeftRightReflection(); ok {
// 				if num != vertical {
// 					return num, 0
// 				}
// 			}
// 			if num, ok := this.TopDownReflection(); ok {
// 				if num != horizontal {
// 					return 0, num
// 				}
// 			}

// 			// flip it back
// 			this[i][j] = this[i][j].Flip()
// 		}
// 	}

// 	panic("no smudge found")
// }

func parsePattern(scanner *bufio.Scanner) Pattern {
	var pattern Pattern
	for scanner.Scan() {
		var row []Terrain
		line := scanner.Text()

		if line == "" {
			break
		}

		for _, r := range line {
			row = append(row, Terrain(r))
		}
		pattern = append(pattern, row)
	}
	return pattern
}

func parse(scanner *bufio.Scanner) []Pattern {
	var patterns []Pattern
	for {
		pattern := parsePattern(scanner)
		if pattern == nil {
			break
		}
		patterns = append(patterns, pattern)
	}
	return patterns
}

func part1(patterns []Pattern) {
	var vertical, horizontal int
	for _, pattern := range patterns {
		// for _, row := range pattern {
		// 	fmt.Println(row)
		// }
		// fmt.Println()
		if num, ok := pattern.LeftRightReflection(); ok {
			vertical += num
		}
		if num, ok := pattern.TopDownReflection(); ok {
			horizontal += num
		}

	}
	fmt.Println(vertical + horizontal*100)
}

func part2(patterns []Pattern) {
	var vertical, horizontal int
	for _, pattern := range patterns {
		// for _, row := range pattern {
		// 	fmt.Println(row)
		// }
		// fmt.Println()
		if num, ok := pattern.LeftRightFind(1); ok {
			vertical += num
		}
		if num, ok := pattern.TopDownFind(1); ok {
			horizontal += num
		}

	}
	fmt.Println(vertical + horizontal*100)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := parse(scanner)

	part1(patterns)
	part2(patterns)
}
