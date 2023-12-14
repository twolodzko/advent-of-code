package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Item rune

const (
	Round  Item = 'O'
	Cube   Item = '#'
	Ground Item = '.'
)

func (this Item) String() string {
	return string(this)
}

type Platform struct {
	platform [][]Item
}

func (this Platform) String() string {
	var str []string
	for _, row := range this.platform {
		var tmp string
		for _, x := range row {
			tmp += string(x)
		}
		str = append(str, tmp)
	}
	return strings.Join(str, "\n")
}

func (this Platform) TiltNorth() {
	var free int
	for j := 0; j < len(this.platform[0]); j++ {
		free = math.MaxInt
		for i := 0; i < len(this.platform); i++ {
			switch this.platform[i][j] {
			case Round:
				if i > free {
					// tilt
					this.platform[free][j] = Round
					this.platform[i][j] = Ground

					for k := free + 1; k <= i; k++ {
						if this.platform[k][j] == Ground {
							free = k
							break
						}
					}
				}
			case Ground:
				free = min(free, i)
			case Cube:
				free = math.MaxInt
			}
		}
	}
}

func (this Platform) TiltSouth() {
	var free int
	for j := 0; j < len(this.platform[0]); j++ {
		free = 0
		for i := len(this.platform) - 1; i >= 0; i-- {
			switch this.platform[i][j] {
			case Round:
				if i < free {
					// tilt
					this.platform[free][j] = Round
					this.platform[i][j] = Ground

					for k := free - 1; k >= i; k-- {
						if this.platform[k][j] == Ground {
							free = k
							break
						}
					}
				}
			case Ground:
				free = max(free, i)
			case Cube:
				free = 0
			}
		}
	}
}

func (this Platform) TiltWest() {
	var free int
	for i := 0; i < len(this.platform); i++ {
		free = math.MaxInt
		for j := 0; j < len(this.platform[0]); j++ {
			switch this.platform[i][j] {
			case Round:
				if j > free {
					// tilt
					this.platform[i][free] = Round
					this.platform[i][j] = Ground

					for k := free + 1; k <= j; k++ {
						if this.platform[i][k] == Ground {
							free = k
							break
						}
					}
				}
			case Ground:
				free = min(free, j)
			case Cube:
				free = math.MaxInt
			}
		}
	}
}

func (this Platform) TiltEast() {
	var free int
	for i := 0; i < len(this.platform); i++ {
		free = 0
		for j := len(this.platform[0]) - 1; j >= 0; j-- {
			switch this.platform[i][j] {
			case Round:
				if j < free {
					// tilt
					this.platform[i][free] = Round
					this.platform[i][j] = Ground

					for k := free - 1; k >= j; k-- {
						if this.platform[i][k] == Ground {
							free = k
							break
						}
					}
				}
			case Ground:
				free = max(free, j)
			case Cube:
				free = 0
			}
		}
	}
}

func (this Platform) TiltCycle() {
	this.TiltNorth()
	this.TiltWest()
	this.TiltSouth()
	this.TiltEast()
}

func (this Platform) Load() int {
	var load int
	for i := 0; i < len(this.platform); i++ {
		for j := 0; j < len(this.platform[0]); j++ {
			if this.platform[i][j] == Round {
				load += len(this.platform) - i
			}
		}
	}
	return load
}

func parse(scanner *bufio.Scanner) Platform {
	var platform [][]Item
	for scanner.Scan() {
		line := scanner.Text()
		var row []Item
		for _, r := range line {
			row = append(row, Item(r))
		}
		platform = append(platform, row)
	}
	return Platform{platform}
}

func (this Platform) Copy() Platform {
	var new Platform
	for _, row := range this.platform {
		var tmp []Item
		for _, x := range row {
			tmp = append(tmp, x)
		}
		new.platform = append(new.platform, tmp)
	}
	return new
}

func part1(platform Platform) {
	this := platform.Copy()
	this.TiltNorth()
	fmt.Println(this.Load())
}

func part2(platform Platform) {
	this := platform.Copy()
	// because Go doesn't support slices as map keys :(
	memory := make(map[string]int)

	// fmt.Println()
	// fmt.Println("Initial state:")
	// fmt.Println(this)
	// fmt.Println()

	var (
		i        int = 0
		max_iter int = 1000000000
	)
	for i < max_iter {
		this.TiltCycle()

		// fmt.Printf("After %d cycles:\n", i+1)
		// fmt.Println(this)
		// fmt.Println()

		if val, ok := memory[this.String()]; ok {
			iters_left := max_iter - i
			cycle_size := i - val
			number_jumps := iters_left / cycle_size
			jump := number_jumps * cycle_size

			if i+jump <= max_iter {
				// skip repeating cycles
				i += jump + 1
			} else {
				i++
			}
		} else {
			memory[this.String()] = i
			i++
		}
	}

	fmt.Println(this.Load())
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	platform := parse(scanner)

	part1(platform)
	part2(platform)
}
