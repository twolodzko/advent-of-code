package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cube struct {
	qty   int
	color string
}

func parseLine(line string) (int, [][]cube) {
	input := strings.Split(line, ":")
	game_id, err := strconv.Atoi(strings.Split(input[0], " ")[1])
	if err != nil {
		panic(err)
	}
	parts := strings.Split(input[1], ";")
	var game [][]cube
	for _, part := range parts {
		var cubes []cube
		for _, part := range strings.Split(part, ",") {
			part := strings.Split(strings.TrimSpace(part), " ")
			qty, err := strconv.Atoi(part[0])
			if err != nil {
				panic(err)
			}
			color := strings.TrimSpace(part[1])
			cubes = append(cubes, cube{qty, color})
		}
		game = append(game, cubes)
	}
	return game_id, game
}

func part1(file *os.File) {
	bag := map[string]int{
		"red": 12, "green": 13, "blue": 14,
	}

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		id, game := parseLine(line)
		all_valid := true

		for _, cubes := range game {
			for _, cube := range cubes {
				if cube.qty > bag[cube.color] {
					all_valid = false
					break
				}
			}
		}
		if all_valid {
			result += id
		}
	}

	fmt.Println(result)
}

func part2(file *os.File) {
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		_, game := parseLine(line)

		bag := map[string]int{
			"red": 0, "green": 0, "blue": 0,
		}

		for _, cubes := range game {
			for _, cube := range cubes {
				if cube.qty > bag[cube.color] {
					bag[cube.color] = cube.qty
				}
			}
		}

		power := 1
		for _, v := range bag {
			power *= v
		}
		result += power
	}

	fmt.Println(result)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// part1(file)
	part2(file)
}
