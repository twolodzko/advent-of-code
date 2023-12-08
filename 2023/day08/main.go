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
	current string
	step    int
	Map
}

func NewNavigator(start string, network_map Map) Navigator {
	return Navigator{start, 0, network_map}
}

func (n *Navigator) HasNext() bool {
	return n.current[2] != 'Z'
}

func (n *Navigator) Next() {
	current_node, ok := n.nodes[n.current]
	if !ok {
		panic(fmt.Sprintf("'%s' node does not exist", n.current))
	}

	idx := n.step % len(n.instructions)
	if n.instructions[idx] == 'L' {
		n.current = current_node.left
	} else {
		n.current = current_node.right
	}
	n.step++
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	network_map := parse(file)
	if _, ok := network_map.nodes["AAA"]; ok {
		part1(network_map)
	}
	part2(network_map)
}
