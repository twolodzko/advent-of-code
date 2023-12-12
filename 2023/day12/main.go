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
	pattern           []Spring
	groups            []int
	damaged, expected int
}

type Groups struct {
	groups []int
}

func NewGroups(groups []int) Groups {
	return Groups{groups}
}

func (this *Groups) Contains(size int) bool {
	for i, n := range this.groups {
		if size == n {
			i++
			if i < len(this.groups) {
				this.groups = this.groups[i:]
			} else {
				this.groups = nil
			}
			return true
		}
	}
	return false
}

// Input has groups of the expected sizes
func (this Pattern) Matches(input []Spring) bool {
	var count, group int
	for _, x := range input {
		switch x {
		case Operational:
			if count > 0 {
				if group >= len(this.groups) || count != this.groups[group] {
					return false
				}
				count = 0
				group++
			}
		case Damaged:
			count++
		}
	}
	if count > 0 {
		return group == len(this.groups)-1 && count == this.groups[group]
	} else {
		return group == len(this.groups)
	}
}

func (this Pattern) Count(variant []Spring, damaged int) int {

	if damaged == this.expected {
		for i := range variant {
			if variant[i] == Unknown {
				variant[i] = Operational
			}
		}

		if this.Matches(variant) {
			return 1
		} else {
			return 0
		}
	} else if damaged > this.expected {
		return 0
	}

	count := 0
	altered := make([]Spring, len(variant))

	for i, field := range this.pattern {
		if field == Unknown {
			copy(altered, variant)
			altered[i] = Damaged
			count += this.Count(altered, damaged+1)

			// for next variants
			variant[i] = Operational
		}
	}
	return count
}

func (this Pattern) CountArrangements() int {
	variant := make([]Spring, len(this.pattern))
	copy(variant, this.pattern)
	return this.Count(variant, this.damaged)
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

func ParseSprings(line string) ([]Spring, int) {
	var (
		pattern []Spring
		damaged int
	)
	for _, r := range line {
		switch r {
		case '.':
			pattern = append(pattern, Operational)
		case '#':
			pattern = append(pattern, Damaged)
			damaged++
		case '?':
			pattern = append(pattern, Unknown)
		default:
			panic(fmt.Sprintf("invalid spring character: %v", r))
		}
	}
	return pattern, damaged
}

func ParseGroups(line string) ([]int, int) {
	var (
		groups   []int
		expected int
	)
	for _, s := range strings.Split(line, ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		groups = append(groups, num)
		expected += num
	}
	return groups, expected
}

func ParseRow(line string) Pattern {
	fields := strings.Fields(line)
	pattern, allocated := ParseSprings(fields[0])
	groups, expected := ParseGroups(fields[1])
	return Pattern{pattern, groups, allocated, expected}
}

func main() {
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
		result += pattern.CountArrangements()
	}
	fmt.Println(result)
}
