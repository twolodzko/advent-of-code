package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func readFile(filename string) (string, []Rule, error) {
	var (
		template string
		rules    []Rule
	)

	file, err := os.Open(filename)
	if err != nil {
		return template, rules, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	template = scanner.Text()
	if err != nil {
		return template, rules, err
	}

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return template, rules, err
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Split(line, " -> ")
		if len(fields) != 2 {
			return template, rules, fmt.Errorf("invalid input: %v", line)
		}
		rules = append(rules, Rule{fields[0], fields[1]})
	}
	err = scanner.Err()
	return template, rules, err
}

type Rule struct {
	pattern   string
	insertion string
}

type Rules struct {
	rules map[string]string
}

func RulesFromList(lst []Rule) Rules {
	rules := make(map[string]string)
	for _, rule := range lst {
		rules[rule.pattern] = rule.pattern[:1] + rule.insertion
	}
	return Rules{rules}
}

func (r *Rules) Process(template string, numSteps int) RuneCounter {
	for step := 0; step < numSteps; step++ {
		template = r.Transform(template)
	}
	result := RuneCounterFromString(template)
	return result
}

func (r *Rules) Transform(pattern string) string {
	partial := ""
	for i := 0; i < len(pattern)-1; i++ {
		pattern := pattern[i:(i + 2)]
		partial += r.MapPattern(pattern)
	}
	// Rule:
	//   ab -> x
	// is applied as:
	//   ab = ax
	// so we correct by taking:
	//   ax + b
	return partial + pattern[len(pattern)-1:]
}

func (r *Rules) MapPattern(pattern string) string {
	if result, ok := r.rules[pattern]; ok {
		return result
	} else {
		return pattern[:len(pattern)-1]
	}
}

type RuneCounter struct {
	items map[rune]int
}

func NewRuneCounter() RuneCounter {
	items := make(map[rune]int)
	return RuneCounter{items}
}

func RuneCounterFromString(str string) RuneCounter {
	counter := NewRuneCounter()
	for _, r := range str {
		counter.Add(r)
	}
	return counter
}

func (c *RuneCounter) Add(r rune) {
	if _, ok := c.items[r]; ok {
		c.items[r] += 1
	} else {
		c.items[r] = 1
	}
}

// func (c *RuneCounter) Merge(other RuneCounter) {
// 	for key, val := range other.items {
// 		if _, ok := c.items[key]; ok {
// 			c.items[key] += val
// 		} else {
// 			c.items[key] = val
// 		}
// 	}
// }

func (c *RuneCounter) MostCommon() (rune, int) {
	var (
		bestRune  rune
		bestCount int = 0
	)
	for r, c := range c.items {
		if c > bestCount {
			bestRune = r
			bestCount = c
		}
	}
	return bestRune, bestCount
}

func (c *RuneCounter) LeastCommon() (rune, int) {
	var (
		bestRune  rune
		bestCount int = math.MaxInt
	)
	for r, c := range c.items {
		if c < bestCount {
			bestRune = r
			bestCount = c
		}
	}
	return bestRune, bestCount
}

func applySteps(template string, rules Rules, steps int) int {
	counter := rules.Process(template, steps)
	_, mostCommonCount := counter.MostCommon()
	_, leastCommonCount := counter.LeastCommon()
	return mostCommonCount - leastCommonCount
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	template, rawRules, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	rules := RulesFromList(rawRules)

	result1 := applySteps(template, rules, 10)
	fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := applyDepthFirst(template, rules, 10)
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
