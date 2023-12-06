package main

import "fmt"

func test_almanac(almanac Almanac) {
	var seed_to_soil [100]int
	for i := 0; i < 100; i++ {
		seed_to_soil[i] = i
	}

	for _, r := range almanac.maps["seed"].mappings {
		for i := 0; i < r.length; i++ {
			seed_to_soil[r.source+i] = r.dest + i
		}
	}

	fmt.Println(almanac)
	for _, s := range []int{79, 14, 55, 13} {
		fmt.Printf("Seed number %d corresponds to soil number %d.\n", s, seed_to_soil[s])
	}

	// pos, ok := almanac.maps[0].ranges[0].ContainsSeed(98)
	// fmt.Printf("%v contains seed: %v at position %v", almanac.maps[0].ranges[0], ok, pos)

	// pos, ok = almanac.maps[0].ranges[0].ContainsSeed(99)
	// fmt.Printf("%v contains seed: %v at position %v\n", almanac.maps[0].ranges[0], ok, pos)
}

func example(almanac Almanac) {
	almanac.seeds = almanac.seeds[0:1]
	fmt.Println(almanac)
}
