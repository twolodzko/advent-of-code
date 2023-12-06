package main

import (
	"fmt"
	"math"
	"os"
)

type Almanac struct {
	seeds []int
	maps  map[string]Map
}

type Map struct {
	source, dest string
	mappings     []Mapping
}

func (m Map) Translate(source int) (string, int) {
	for _, r := range m.mappings {
		if val, ok := r.Contains(source); ok {
			return m.dest, val
		}
	}
	return m.dest, source
}

type Mapping struct {
	dest, source, length int
}

func (m Mapping) Contains(seed int) (int, bool) {
	if seed >= m.source && seed <= m.source+m.length {
		return seed + m.Delta(), true
	}
	return 0, false
}

func (m Mapping) Delta() int {
	return m.dest - m.source
}

func (a Almanac) FindLocation(value int) int {
	key := "seed"
	for key != "location" {
		// fmt.Printf("%s %d, ", key, value)
		key, value = a.maps[key].Translate(value)
	}
	// fmt.Printf("%s %d.\n", key, value)
	return value
}

func part1(almanac Almanac) {
	smallest := math.MaxInt
	for _, seed := range almanac.seeds {
		value := almanac.FindLocation(seed)
		smallest = min(smallest, value)
	}
	fmt.Println(smallest)
}

func (m Map) Explore(inputs []Range) (string, []Range) {
	var tmp, results []Range
	for _, mapping := range m.mappings {
		r := NewRange(mapping.source, mapping.source+mapping.length)
		for _, v := range inputs {
			if before, ok := v.Before(r); ok {
				tmp = append(tmp, before)
			}
			if after, ok := v.After(r); ok {
				tmp = append(tmp, after)
			}
			if intersection, ok := v.Intersection(r); ok {
				intersection := intersection.Add(mapping.Delta())
				results = append(results, intersection)
			}
		}
		inputs = tmp
		tmp = nil
	}

	return m.dest, append(results, inputs...)
}

func (a Almanac) ExploreLocations(values []Range) []Range {
	key := "seed"
	for key != "location" {
		// fmt.Printf("%s %d, ", key, values)
		key, values = a.maps[key].Explore(values)
	}
	// fmt.Printf("%s %d.\n", key, values)
	return values
}

func part1b(almanac Almanac) {
	smallest := Range{math.MaxInt, math.MaxInt}
	for _, seed := range almanac.seeds {
		value := almanac.ExploreLocations([]Range{{seed, seed}})
		smallest = smallest.Min(value[0])
	}
	fmt.Println(smallest)
}

func part2(almanac Almanac) {
	smallest := math.MaxInt
	for i := 0; i < len(almanac.seeds); i += 2 {
		start, length := almanac.seeds[i], almanac.seeds[i+1]
		results := almanac.ExploreLocations([]Range{{start, start + length}})

		for _, r := range results {
			smallest = min(smallest, r.min)
		}
	}
	fmt.Println(smallest)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	almanac := parse(file)

	part1(almanac)
	// part1b(almanac)
	part2(almanac)
}
