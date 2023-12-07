package part1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	Card rune
	Hand []Card
)

func (hand Hand) Counts() []int {
	unique_counts := make(map[Card]int)
	for _, card := range hand {
		if _, ok := unique_counts[card]; !ok {
			unique_counts[card] = 0
		}
		unique_counts[card]++
	}

	var counts []int
	for _, count := range unique_counts {
		counts = append(counts, count)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	return counts
}

func (card Card) Value() int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return 9
	case 'T':
		return 8
	default:
		return int(card - '2')
	}
}

func (lhs Hand) Less(rhs Hand) bool {
	lhs_s := lhs.Strength()
	rhs_s := rhs.Strength()

	if lhs_s == rhs_s {
		for i := 0; i < 5; i++ {
			a := lhs[i].Value()
			b := rhs[i].Value()
			if a != b {
				return a < b
			}
		}
		// they are the same strength
		return false
	}
	return lhs_s < rhs_s
}

func (hand Hand) Strength() int {
	counts := hand.Counts()

	switch len(counts) {
	case 1:
		// Five of a kind
		// all five cards have the same label
		return 6
	case 2:
		switch counts[0] {
		case 4:
			// Four of a kind
			// four cards have the same label and one card has a different label
			return 5
		case 3:
			// Full house
			// three cards have the same label, and the remaining two cards share a different label
			return 4
		}
	case 3:
		switch counts[0] {
		case 3:
			// Three of a kind
			// three cards have the same label, and the remaining two cards are each different from any other card in the hand
			return 3
		case 2:
			// Two pair
			// two cards share one label, two other cards share a second label, and the remaining card has a third label
			return 2
		}
	case 4:
		// One pair
		// where two cards share one label, and the other three cards have a different label from the pair and each other
		return 1
	case 5:
		// High card
		// all cards' labels are distinct
		return 0
	}

	panic(fmt.Sprintf("impossible hand: %v", hand))
}

type Bids struct {
	hands []Hand
	bids  []int
}

func (bids Bids) Len() int {
	return len(bids.hands)
}

func (bids Bids) Swap(i, j int) {
	bids.hands[i], bids.hands[j] = bids.hands[j], bids.hands[i]
	bids.bids[i], bids.bids[j] = bids.bids[j], bids.bids[i]
}

func (bids Bids) Less(i, j int) bool {
	return bids.hands[i].Less(bids.hands[j])
}

func parseLine(line string) (Hand, int) {
	fields := strings.Fields(line)
	bid, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return Hand(fields[0]), bid
}

func parse() Bids {
	var (
		hands []Hand
		bids  []int
	)

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		hand, bid := parseLine(line)
		hands = append(hands, hand)
		bids = append(bids, bid)
	}
	return Bids{hands, bids}
}

func Solution() {
	bids := parse()

	sort.Sort(bids)
	result := 0
	for i := 0; i < len(bids.bids); i++ {
		// fmt.Printf("Hand: %v with bid %d has rank %d\n", bids.hands[i], bids.bids[i], i+1)
		result += (i + 1) * bids.bids[i]
	}
	fmt.Println(result)
}
