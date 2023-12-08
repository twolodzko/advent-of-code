package main

import "fmt"

func part1(network_map Map) {
	navigator := NewNavigator("AAA", network_map)
	for navigator.Next() {
	}
	fmt.Println(navigator.step)
}
