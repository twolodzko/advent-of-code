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

		if _, ok := paths[fields[0]]; !ok {
			paths[fields[0]] = []string{}
		}
		paths[fields[0]] = append(paths[fields[0]], fields[1])

		if fields[0] != "start" && fields[1] != "end" {
			if _, ok := paths[fields[1]]; !ok {
				paths[fields[1]] = []string{}
			}
			paths[fields[1]] = append(paths[fields[1]], fields[0])
		}

	}
	err = scanner.Err()
	return paths, err
}

type Explorer struct {
	path      string
	explored  []string
	corridors map[string][]string
}

func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func newExplorer(path string, explored []string, corridors map[string][]string) Explorer {
	return Explorer{path, explored, corridors}
}

func (e *Explorer) NextSteps() []Explorer {
	var unexplored []Explorer
	explored := append(e.explored, e.path)
	for _, path := range e.corridors[e.path] {
		if e.WasVisited(path) {
			continue
		}
		unexplored = append(unexplored, newExplorer(path, explored, e.corridors))
	}
	return unexplored
}

func (e *Explorer) Done() bool {
	return e.path == "end"
}

func (e *Explorer) WasVisited(path string) bool {
	if !IsLower(path) {
		return false
	}
	for _, x := range e.explored {
		if path == x {
			return true
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

func explore(paths map[string][]string) int {
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
		queue.Push(explorer.NextSteps())
	}
	// for _, e := range found {
	// 	fmt.Printf("%v -> %v (%v)\n", e.explored, e.path, e.blocked)
	// }
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

	result1 := explore(paths)
	fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := firstSynchronization(deepcopy(arr))
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
