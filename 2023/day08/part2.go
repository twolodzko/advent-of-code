package main

import (
	"fmt"
	"math"
)

func naive_search(cycle_lengths []int) int {
	done := func(x int) bool {
		for _, n := range cycle_lengths {
			if x%n != 0 {
				return false
			}
		}
		return true
	}

	step := math.MaxInt
	for _, x := range cycle_lengths {
		if x < step {
			step = x
		}
	}

	var i = 0
	for {
		i += step
		if done(i) {
			return i
		}
	}
}

// Greatest common divisor (of non-negative numbers)
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Least common multiple (of non-negative numbers)
// See https://en.wikipedia.org/wiki/Least_common_multiple
func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func lcm_search(cycle_lengths []int) int {
	if len(cycle_lengths) == 1 {
		return cycle_lengths[0]
	}
	result := Lcm(cycle_lengths[0], cycle_lengths[1])
	for i := 2; i < len(cycle_lengths); i++ {
		result = Lcm(result, cycle_lengths[i])
	}
	return result
}

func part2(network_map Map) {
	var start_nodes []string
	for name := range network_map.nodes {
		if name[2] == 'A' {
			start_nodes = append(start_nodes, name)
		}
	}

	var cycle_lengths []int
	for _, start := range start_nodes {
		navigator := NewNavigator(start, network_map)
		for navigator.Next() {
		}
		cycle_lengths = append(cycle_lengths, navigator.step)
	}

	fmt.Println(lcm_search(cycle_lengths))
}
