package part2

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

type Count struct {
	Key   Card
	Value int
}

type Counter struct {
	counts []Count
}

func (counter *Counter) Add(card Card) {
	for i, count := range counter.counts {
		if count.Key == card {
			counter.counts[i].Value++
			return
		}
	}
	counter.counts = append(counter.counts, Count{card, 1})
}

func (counter Counter) Len() int {
	return len(counter.counts)
}

func (counter Counter) Less(i, j int) bool {
	return counter.counts[i].Value > counter.counts[j].Value
}

func (counter Counter) Swap(i, j int) {
	counter.counts[i], counter.counts[j] = counter.counts[j], counter.counts[i]
}

func (counter *Counter) Pop(card Card) int {
	for i, count := range counter.counts {
		if count.Key == card {
			counter.counts = append(counter.counts[:i], counter.counts[i+1:]...)
			return count.Value
		}
	}
	return 0
}

func (counter Counter) First() Count {
	if len(counter.counts) == 0 {
		return Count{}
	}
	return counter.counts[0]
}

func (card Card) Value() int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'T':
		return 9
	case 'J':
		return 0
	default:
		return int(card - '1')
	}
}

func (card Card) String() string {
	return string(rune(card))
}

func (hand Hand) String() string {
	var str string
	for _, card := range hand {
		str += string(card)
	}
	return str
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
	var counts Counter
	for _, card := range hand {
		counts.Add(card)
	}
	sort.Sort(counts)

	if n := counts.Pop('J'); n != 0 {
		if n == 5 {
			counts.counts = append(counts.counts, Count{'A', n})
		} else {
			counts.counts[0].Value += n
		}
	}
	sort.Sort(counts)

	switch counts.Len() {
	case 1:
		// Five of a kind
		// all five cards have the same label
		return 6
	case 2:
		switch counts.First().Value {
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
		switch counts.First().Value {
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
