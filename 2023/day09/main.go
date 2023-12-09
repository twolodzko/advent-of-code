package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Forecast(arr []int) (int, int) {
	if AllZeros(arr) {
		return 0, 0
	}
	first, last := Forecast(Diff(arr))
	return arr[0] - first, arr[len(arr)-1] + last
}

func AllZeros(arr []int) bool {
	for _, x := range arr {
		if x != 0 {
			return false
		}
	}
	return true
}

func Diff(arr []int) []int {
	var result []int
	for i := 1; i < len(arr); i++ {
		result = append(result, arr[i]-arr[i-1])
	}
	return result
}

func parseLine(line string) []int {
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

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	part1, part2 := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := parseLine(line)
		first, last := Forecast(arr)
		part1 += first
		part2 += last
	}
	fmt.Println(part1, part2)
}
