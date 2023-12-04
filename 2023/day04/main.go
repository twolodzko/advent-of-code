package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id            int
	your, winning []int
	count         int
}

func parseLine(line string) Card {
	var your, winning []int

	// Example:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	re := regexp.MustCompile(`Card +([0-9]+): ([0-9 ]+) +\| +([0-9 ]+)`)

	matched := re.FindStringSubmatch(line)
	if len(matched) != 4 {
		panic("failed to parse")
	}

	id, err := strconv.Atoi(matched[1])
	if err != nil {
		panic(err)
	}

	for _, s := range strings.Fields(matched[2]) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		your = append(your, num)
	}

	for _, s := range strings.Fields(matched[3]) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		winning = append(winning, num)
	}

	return Card{id, your, winning, 1}
}

func countMatches(s Card) int {
	count := 0
	for _, x := range s.your {
		if slices.Contains(s.winning, x) {
			count++
		}
	}
	return count
}

func pow(x, n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

func part1(cards []Card) {
	result := 0
	for _, card := range cards {
		result += pow(2, countMatches(card)) / 2
	}
	fmt.Println(result)
}

func part2(cards []Card) {
	for i, card := range cards {
		n := countMatches(card)
		for j := i; j < i+n; j++ {
			// fmt.Printf("With card %d you win card %d\n", card.id, cards[j+1].id)
			cards[j+1].count += card.count
		}
	}

	result := 0
	for _, card := range cards {
		// fmt.Printf("%d cards with number %d\n", card.count, card.id)
		result += card.count
	}
	fmt.Println(result)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cards []Card
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		scratchcard := parseLine(line)
		cards = append(cards, scratchcard)
	}

	part1(cards)
	part2(cards)
}
