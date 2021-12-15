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
		chunk := pattern[i:(i + 2)]
		partial += r.MapPattern(chunk)
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

type DepthExpander struct {
	rules  Rules
	memory map[State]RuneCounter
}

type State struct {
	pattern string
	depth   int
}

func DepthExpanderFromList(lst []Rule) DepthExpander {
	rules := RulesFromList(lst)
	memory := make(map[State]RuneCounter)
	return DepthExpander{rules, memory}
}

func (e *DepthExpander) Process(pattern string, depth int) RuneCounter {
	counter := e.Step(pattern, depth)
	counter.Add(rune(pattern[len(pattern)-1]))
	return counter
}

func (e *DepthExpander) Step(pattern string, depth int) RuneCounter {
	state := State{pattern, depth}
	if result, ok := e.memory[state]; ok {
		return result
	}

	var result RuneCounter

	if depth == 0 {
		result = RuneCounterFromString(pattern[:len(pattern)-1])
	} else {
		counter := NewRuneCounter()
		for i := 0; i < len(pattern)-1; i++ {
			first := pattern[i:(i + 1)]
			second := pattern[(i + 1):(i + 2)]
			first = e.rules.MapPattern(first + second)

			partial := e.Step(first+second, depth-1)
			counter.Merge(partial)
		}
		result = counter
	}

	e.memory[state] = result
	return result
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

func (c *RuneCounter) Merge(other RuneCounter) {
	for key, val := range other.items {
		if _, ok := c.items[key]; ok {
			c.items[key] += val
		} else {
			c.items[key] = val
		}
	}
}

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

func applySteps(template string, rules []Rule, steps int, memoize bool) int {
	var counter RuneCounter
	if !memoize {
		r := RulesFromList(rules)
		counter = r.Process(template, steps)
	} else {
		e := DepthExpanderFromList(rules)
		counter = e.Process(template, steps)
	}
	_, mostCommonCount := counter.MostCommon()
	_, leastCommonCount := counter.LeastCommon()
	return mostCommonCount - leastCommonCount
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	template, rules, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1 := applySteps(template, rules, 10, false)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := applySteps(template, rules, 40, true)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
