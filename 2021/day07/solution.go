package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, error) {
	var arr []int

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	if err != nil {
		return arr, err
	}
	for _, val := range strings.Split(line, ",") {
		i, err := strconv.Atoi(val)
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}
	err = scanner.Err()
	return arr, err
}

func median(arr []int) int {
	sort.Ints(arr)
	return arr[(len(arr) / 2)]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distToMedian(arr []int) int {
	med := median(arr)
	total := 0
	for _, x := range arr {
		total += abs(x - med)
	}
	return total
}

func minMax(arr []int) (int, int) {
	sort.Ints(arr)
	return arr[0], arr[len(arr)-1]
}

func customDist(arr []int, medioid int) int {
	total := 0
	for _, x := range arr {
		d := abs(x - medioid)
		// 1 + 2 + ... + n = (n*(n-1)) / 2
		// see: https://math.stackexchange.com/a/60581/114961
		total += d + (d*(d-1))/2
	}
	return total
}

func bestCustomDistance(arr []int) int {
	// if we were to go over it means we are not using enough precision
	best := math.MaxInt32
	min, max := minMax(arr)
	// no need to check the extreme values, but it is fast enough
	// so it's easier to iterate over all the cases than handle edge cases
	for medioid := min; medioid <= max; medioid++ {
		d := customDist(arr, medioid)
		if d < best {
			best = d
		}
	}
	return best
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	arr, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// median minimizes absolute error
	result1 := distToMedian(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := bestCustomDistance(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
