package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func readFile(filename string) (map[string][]string, error) {
	paths := make(map[string][]string)

	file, err := os.Open(filename)
	if err != nil {
		return paths, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return paths, err
		}
		fields := strings.Split(line, "-")
		if len(fields) != 2 {
			return paths, fmt.Errorf("invalid input: %s", line)
		}

		paths = add(paths, fields[0], fields[1])
		paths = add(paths, fields[1], fields[0])
	}
	err = scanner.Err()
	return paths, err
}

func add(paths map[string][]string, from, to string) map[string][]string {
	if from == "end" || to == "start" {
		return paths
	}
	if _, ok := paths[from]; !ok {
		paths[from] = []string{}
	}
	paths[from] = append(paths[from], to)
	return paths
}

type Explorer struct {
	path      string
	explored  []string
	corridors map[string][]string
}

func (e *Explorer) Done() bool {
	return e.path == "end"
}

func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func (e *Explorer) NextSteps(shouldSkip func([]string, string) bool) []Explorer {
	var unexplored []Explorer
	explored := append(e.explored, e.path)
	for _, path := range e.corridors[e.path] {
		if IsLower(path) && shouldSkip(explored, path) {
			continue
		}
		unexplored = append(unexplored, Explorer{path, explored, e.corridors})
	}
	return unexplored
}

func wasVisited(explored []string, path string) bool {
	for _, x := range explored {
		if path == x {
			return true
		}
	}
	return false
}

func invalidPath(explored []string, path string) bool {
	counter := make(map[string]int)
	for _, p := range explored {
		if IsLower(p) {
			if _, ok := counter[p]; !ok {
				counter[p] = 1
			} else {
				counter[p] += 1
			}
		}
	}
	if count, ok := counter[path]; ok && count > 1 {
		return true
	}
	if count, ok := counter[path]; ok && count == 1 {
		for _, count := range counter {
			if count > 1 {
				return true
			}
		}
	}
	return false
}

type Queue struct {
	items []Explorer
}

func (q *Queue) Push(e []Explorer) {
	q.items = append(q.items, e...)
}

func (q *Queue) Pop() Explorer {
	if len(q.items) == 0 {
		log.Fatal("empty queue")
	}
	p := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return p
}

func (q *Queue) HasNext() bool {
	return len(q.items) > 0
}

func explore(paths map[string][]string, once bool) int {
	var skipFn func([]string, string) bool
	if once {
		skipFn = wasVisited
	} else {
		skipFn = invalidPath
	}

	var queue Queue
	for _, path := range paths["start"] {
		queue.Push([]Explorer{{path, []string{"start"}, paths}})
	}

	var found []Explorer
	for queue.HasNext() {
		explorer := queue.Pop()
		if explorer.Done() {
			found = append(found, explorer)
		}
		queue.Push(explorer.NextSteps(skipFn))
	}
	return len(found)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	paths, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1 := explore(paths, true)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := explore(paths, false)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
