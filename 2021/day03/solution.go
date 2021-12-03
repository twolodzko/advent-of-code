package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func readFile(filename string) ([][]int, error) {
	var arr [][]int

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}

		var row []int
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		arr = append(arr, row)
	}
	err = scanner.Err()
	return arr, err
}

func binToDec(binary []int) int {
	n := len(binary)
	dec := 0
	for i := 0; i < n; i++ {
		dec += binary[n-i-1] * int(math.Pow(2, float64(i)))
	}
	return dec
}

func powerConsumption(bits [][]int) int {
	var (
		gammaBits   []int
		epsilonBits []int
		total       []int
	)
	for i := 0; i < len(bits[0]); i++ {
		total = append(total, 0)
	}
	for _, line := range bits {
		if len(line) != len(total) {
			panic(fmt.Errorf("Invalid input: %v", line))
		}
		for i, x := range line {
			total[i] += x
		}
	}
	n := len(bits)
	for i := 0; i < len(total); i++ {
		if total[i] >= (n / 2) {
			gammaBits = append(gammaBits, 1)
		} else {
			gammaBits = append(gammaBits, 0)
		}
		epsilonBits = append(epsilonBits, 1-gammaBits[i])
	}
	gamma := binToDec(gammaBits)
	epsilon := binToDec(epsilonBits)
	return gamma * epsilon
}

func group(arr [][]int, pos int) (one [][]int, zero [][]int) {
	for _, row := range arr {
		if row[pos] == 1 {
			one = append(one, row)
		} else {
			zero = append(zero, row)
		}
	}
	return one, zero
}

func rating(arr [][]int, gt bool) int {
	tmp := arr
	n := len(arr[0])
	for i := 0; i < n; i++ {
		one, zero := group(tmp, i)
		if gt {
			if len(zero) <= len(one) {
				tmp = one
			} else {
				tmp = zero
			}
		} else {
			if len(zero) > len(one) {
				tmp = one
			} else {
				tmp = zero
			}
		}

		if len(tmp) == 1 {
			break
		}
	}
	return binToDec(tmp[0])
}

func lifeSupportRating(arr [][]int) int {
	oxygenGeneratorRating := rating(arr, true)
	co2ScrubberRating := rating(arr, false)
	return oxygenGeneratorRating * co2ScrubberRating
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

	result1 := powerConsumption(arr)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := lifeSupportRating(arr)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
