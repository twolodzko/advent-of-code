package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) ([]int, error) {
	var arr []int

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}

	if err := scanner.Err(); err != nil {
		return arr, err
	}
	return arr, err
}

func countIncreases(arr []int) int {
	count := 0
	var previous int

	for i, current := range arr {
		if i == 0 {
			previous = current
		}
		if current > previous {
			count += 1
		}
		previous = current
	}
	return count
}

func slidingWindowIncreases(arr []int) int {
	if len(arr) < 4 {
		log.Fatal("Invalid input")
	}
	var (
		first  int = 0
		second int = 0
		count  int = 0
	)
	for i := 0; i < len(arr)-1; i++ {
		first += arr[i]
		second += arr[i+1]

		if i < 2 {
			// initialization only
			continue
		}

		if first < second {
			count += 1
		}

		first -= arr[i-2]
		second -= arr[i-1]
	}

	return count
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
	result1 := countIncreases(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := slidingWindowIncreases(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
