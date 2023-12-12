package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spring rune

const (
	Operational Spring = '.'
	Damaged     Spring = '#'
	Unknown     Spring = '?'
)

func (this Spring) Matches(other Spring) bool {
	return this == Unknown || this == other
}

func (this Spring) String() string {
	return string(this)
}

type Pattern struct {
	pattern []Spring
	groups  []int
}

type Node struct {
	prefix, size int
}

func (this Node) Size() int {
	return this.prefix + this.size
}

func (this Pattern) Matches(start, end int) bool {
	if end > len(this.pattern) {
		return false
	}
	for i := start; i < end; i++ {
		if this.pattern[i] == Operational {
			return false
		}
	}
	if end < len(this.pattern) {
		return this.pattern[end] != Damaged
	}
	return true
}

func (this Pattern) Explore(start int, positions []int) int {
	if len(this.groups) == len(positions) {
		group := 0
		for i := 0; i < len(this.pattern); i++ {
			if group >= len(this.groups) {
				if this.pattern[i] == Damaged {
					return 0
				}
			} else if i < positions[group] {
				if this.pattern[i] == Damaged {
					return 0
				}
			} else if i >= positions[group]+this.groups[group] {
				if this.pattern[i] == Damaged {
					return 0
				}
				group++
			}
		}
		return 1
	}
	if start >= len(this.pattern) {
		return 0
	}

	group_size := this.groups[len(positions)]
	count := 0
	for i := start; i < len(this.pattern); i++ {
		end := i + group_size
		if this.Matches(i, end) {
			count += this.Explore(end+1, append(positions, i))
		}
	}
	return count
}

func (this Pattern) CountArrangements() int {
	min_size := max(0, len(this.groups)-1)
	for _, n := range this.groups {
		min_size += n
	}
	if len(this.pattern) == min_size {
		return 1
	}
	return this.Explore(0, []int{})
}

func ToStringJoined[T any](arr []T, sep string) string {
	var tmp []string
	for _, x := range arr {
		tmp = append(tmp, fmt.Sprint(x))
	}
	return strings.Join(tmp, sep)
}

func (this Pattern) String() string {
	pattern := ToStringJoined(this.pattern, "")
	groups := ToStringJoined(this.groups, ",")
	return fmt.Sprintf("%s %s", pattern, groups)
}

func ParseSprings(line string) []Spring {
	var pattern []Spring
	for _, r := range line {
		switch r {
		case '.':
			pattern = append(pattern, Operational)
		case '#':
			pattern = append(pattern, Damaged)
		case '?':
			pattern = append(pattern, Unknown)
		default:
			panic(fmt.Sprintf("invalid spring character: %v", r))
		}
	}
	return pattern
}

func ParseGroups(line string) []int {
	var groups []int
	for _, s := range strings.Split(line, ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		groups = append(groups, num)
	}
	return groups
}

func ParseRow(line string) Pattern {
	fields := strings.Fields(line)
	pattern := ParseSprings(fields[0])
	groups := ParseGroups(fields[1])

	var damaged int
	for _, x := range pattern {
		if x == Damaged {
			damaged++
		}
	}

	return Pattern{pattern, groups}
}

func part1() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		pattern := ParseRow(line)
		arrangements := pattern.CountArrangements()
		result += arrangements
		// fmt.Printf("%s - %d arrangements\n", line, arrangements)
	}
	fmt.Println(result)
}

func repeat5(line string) string {
	fields := strings.Fields(line)

	var springs, groups []string
	for i := 0; i < 5; i++ {
		springs = append(springs, fields[0])
		groups = append(groups, fields[1])
	}

	return fmt.Sprintf("%s %s", strings.Join(springs, "?"), strings.Join(groups, ","))
}

func part2() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = repeat5(line)
		pattern := ParseRow(line)
		arrangements := pattern.CountArrangements()
		result += arrangements
		// fmt.Printf("%s - %d arrangements\n", line, arrangements)
	}
	fmt.Println(result)
}

func main() {
	part1()
	// part2()
}
