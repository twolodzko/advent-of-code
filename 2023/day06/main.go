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

// Integer square root (using binary search)
func Isqrt(y int) int {
	// See https://en.wikipedia.org/wiki/Integer_square_root
	var (
		L int = 0
		M int
		R int = y + 1
	)
	for L != R-1 {
		M = (L + R) / 2
		if M*M <= y {
			L = M
		} else {
			R = M
		}
	}
	return L
}

func Solve(limit, record int) (int, int) {
	// record > hold_time * (limit - hold_time) = limit * hold_time - hold_time^2
	// 0 = -record + limit * hold_time - hold_time^2
	//
	// 0 = c + bx + ax^2
	// x = (-b +/- sqrt(b^2 - 4ac)) / 2a
	//
	// See:
	// https://www.mathsisfun.com/algebra/quadratic-equation.html

	b := limit
	c := record + 1
	upper := (-b - Isqrt(b*b-4*c)) / -2
	lower := (-b + Isqrt(b*b-4*c)) / -2
	return lower, upper
}

func ExploreSolutions(limit, record int) int {
	count := 0
	for hold_time := 1; hold_time < limit; hold_time++ {
		distance := Distance(hold_time, limit)
		if distance > record {
			count++
		}
	}
	return count
}

func SolutionsCount(limit, record int) int {
	lower, upper := Solve(limit, record)

	// rounding error corrections
	if Distance(lower-1, limit) > record {
		lower--
	} else if Distance(lower, limit) <= record && Distance(lower+1, limit) > record {
		lower++
	}

	if Distance(upper+1, limit) > record {
		upper++
	} else if Distance(upper, limit) <= record && Distance(upper-1, limit) > record {
		upper--
	}

	return upper - lower + 1
}

func parseLine(line string) []int {
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

func parse1() ([]int, []int) {
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
		number_of_strategies := SolutionsCount(time_limit, distance_record)
		result *= number_of_strategies
	}
	fmt.Println(result)
}

func parse2() (int, int) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var line string
	scanner.Scan()
	line = strings.Replace(scanner.Text(), " ", "", -1)
	time := parseLine(line)[0]

	scanner.Scan()
	line = strings.Replace(scanner.Text(), " ", "", -1)
	distance := parseLine(line)[0]

	return time, distance
}

func part2() {
	time_limit, distance_record := parse2()
	number_of_strategies := SolutionsCount(time_limit, distance_record)
	fmt.Println(number_of_strategies)
}

func main() {
	part1()
	part2()
}
