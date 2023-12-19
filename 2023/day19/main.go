package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	workflows, parts := parse(scanner)

	part1(workflows, parts)
	part2(workflows)

	// fmt.Println(system)
	// fmt.Println(parts)
}
