package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	left, right string
}

func ParseNode(line string) (string, Node) {
	re := regexp.MustCompile(`([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)`)
	matches := re.FindStringSubmatch(line)
	return matches[1], Node{matches[2], matches[3]}
}

type Map struct {
	instructions []rune
	nodes        map[string]Node
}

func parse(file *os.File) Map {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := []rune(scanner.Text())

	nodes := make(map[string]Node)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		name, node := ParseNode(line)
		nodes[name] = node
	}
	return Map{instructions, nodes}
}

type Navigator struct {
	node        string
	step        int
	network_map Map
}

func NewNavigator(start string, network_map Map) Navigator {
	return Navigator{start, 0, network_map}
}

func (n *Navigator) Next() bool {
	if n.node[2] == 'Z' {
		return false
	}

	current, ok := n.network_map.nodes[n.node]
	if !ok {
		panic(fmt.Sprintf("'%s' node does not exist", n.node))
	}

	idx := n.step % len(n.network_map.instructions)
	if n.network_map.instructions[idx] == 'L' {
		n.node = current.left
	} else {
		n.node = current.right
	}
	n.step++

	return true
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	network_map := parse(file)
	part1(network_map)
	part2(network_map)
}
