package main

import (
	"fmt"
	"math"
)

func search(cycle_lengths []int) int {
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

	fmt.Println(search(cycle_lengths))
}
