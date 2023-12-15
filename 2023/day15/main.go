package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Hash(input string) int {
	current := 0
	for _, r := range input {
		current += int(r)
		current *= 17
		current %= 256
	}
	return current
}

func part1(sequence []string) {
	result := 0
	for _, step := range sequence {
		result += Hash(step)
	}
	fmt.Println(result)
}

type Lens struct {
	label  string
	length int
}

var parser = regexp.MustCompile(`([a-z]+)(=|-)([0-9]+)?`)

func split(str string) (string, rune, int) {
	fields := parser.FindStringSubmatch(str)

	label := fields[1]
	operation := rune(fields[2][0])

	var (
		length int
		err    error
	)
	if operation == '=' {
		length, err = strconv.Atoi(fields[3])
		if err != nil {
			panic(err)
		}
	}
	return label, operation, length
}

type Initializer struct {
	boxes [256][]Lens
}

func (this *Initializer) Step(input string) {
	label, operation, length := split(input)
	box := Hash(label)

	if operation == '=' {
		var existed bool
		for i, old := range this.boxes[box] {
			if old.label == label {
				this.boxes[box][i].length = length
				existed = true
				break
			}
		}
		if !existed {
			this.boxes[box] = append(this.boxes[box], Lens{label, length})
		}
	} else {
		for i, old := range this.boxes[box] {
			if old.label == label {
				if i+1 < len(this.boxes[box]) {
					this.boxes[box] = append(this.boxes[box][:i], this.boxes[box][i+1:]...)
				} else {
					this.boxes[box] = this.boxes[box][:i]
				}
			}
		}
	}
}

func (this Initializer) FocusingPower() int {
	total := 0
	for i, box := range this.boxes {
		for j, lens := range box {
			total += (i + 1) * (j + 1) * lens.length
		}
	}
	return total
}

func part2(sequence []string) {
	initializer := Initializer{}
	for _, step := range sequence {
		initializer.Step(step)
		// fmt.Printf("After '%s':\n", step)
		// for i, box := range initializer.boxes {
		// 	if len(box) > 0 {
		// 		fmt.Printf("Box %d: %v\n", i, box)
		// 	}
		// }
	}
	fmt.Println(initializer.FocusingPower())
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := strings.Split(scanner.Text(), ",")

	part1(sequence)
	part2(sequence)
}
