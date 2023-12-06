package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Distance(hold_time, limit int) int {
	if hold_time <= 0 || hold_time >= limit {
		return 0
	}
	return hold_time * (limit - hold_time)
}

func parse1() ([]int, []int) {
	parseLine := func(line string) []int {
		line = strings.Split(line, ":")[1]

		var result []int
		for _, s := range strings.Fields(line) {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			result = append(result, num)
		}
		return result
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := parseLine(scanner.Text())
	scanner.Scan()
	distance := parseLine(scanner.Text())
	return time, distance
}

func part1() {
	times, distances := parse1()

	result := 1
	n := len(times)
	for i := 0; i < n; i++ {
		time_limit := times[i]
		distance_record := distances[i]

		number_of_strategies := 0
		for hold_time := 1; hold_time < time_limit; hold_time++ {
			distance := Distance(hold_time, time_limit)
			if distance > distance_record {
				number_of_strategies++
			}
		}
		result *= number_of_strategies
	}
	fmt.Println(result)
}

func parse2() (int, int) {
	parseLine := func(line string) int {
		line = strings.Split(line, ":")[1]
		line = strings.Replace(line, " ", "", -1)
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		return num
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	time := parseLine(scanner.Text())
	scanner.Scan()
	distance := parseLine(scanner.Text())
	return time, distance
}

func part2() {
	time_limit, distance_record := parse2()
	number_of_strategies := 0
	for hold_time := 1; hold_time < time_limit; hold_time++ {
		distance := Distance(hold_time, time_limit)
		if distance > distance_record {
			number_of_strategies++
		}
	}
	fmt.Println(number_of_strategies)
}

func main() {
	part1()
	part2()
}
