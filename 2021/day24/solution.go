package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
	var arr []string

	file, err := os.Open(filename)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	err = scanner.Err()
	return arr, err
}

type Program struct {
	variables    map[string]int
	instructions []Instruction
}

type Instruction struct {
	fn   string
	a, b interface{}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %v %v", i.fn, i.a, i.b)
}

func (i *Instruction) Call(env *map[string]int) (int, error) {
	var (
		a, b, result int
		ok           bool
	)

	switch v := i.a.(type) {
	case int:
		a = v
	case string:
		a, ok = (*env)[v]
		if !ok {
			return 0, fmt.Errorf("invalid variable: %v", v)
		}
	default:
		return 0, fmt.Errorf("invalid argument: %v (%T)", i.a, i.a)
	}

	switch v := i.b.(type) {
	case int:
		b = v
	case string:
		b, ok = (*env)[v]
		if !ok {
			return 0, fmt.Errorf("invalid variable: %v", v)
		}
	default:
		return 0, fmt.Errorf("invalid argument: %v (%T)", i.a, i.a)
	}

	switch i.fn {
	case "add":
		result = a + b
	case "mul":
		result = a * b
	case "div":
		result = a / b
	case "mod":
		result = a % b
	case "eql":
		if a == b {
			result = 1
		} else {
			result = 0
		}
	default:
		return 0, fmt.Errorf("invalid instruction: %v", i.fn)
	}

	return result, nil
}

func MaybeNumber(s string) (interface{}, error) {
	if IsNumber(s) {
		return strconv.Atoi(s)
	} else {
		return s, nil
	}
}

func NewInstruction(s string) (Instruction, error) {
	fields := strings.Fields(s)

	if fields[0] == "inp" {
		if len(fields) != 2 {
			return Instruction{}, fmt.Errorf("invalid instruction: %v", fields)
		}
		return Instruction{fields[0], fields[1], nil}, nil
	} else {
		if len(fields) != 3 {
			return Instruction{}, fmt.Errorf("invalid instruction: %v", fields)
		}
		a, err := MaybeNumber(fields[1])
		if err != nil {
			return Instruction{}, err
		}
		b, err := MaybeNumber(fields[2])
		if err != nil {
			return Instruction{}, err
		}
		return Instruction{fields[0], a, b}, nil
	}
}

func (p *Program) Parse(lines []string) error {
	for _, line := range lines {
		instruction, err := NewInstruction(line)
		if err != nil {
			return err
		}
		p.instructions = append(p.instructions, instruction)
	}
	return nil
}

func NewProgram(lines []string) (*Program, error) {
	vars := make(map[string]int)
	prog := Program{vars, nil}
	err := prog.Parse(lines)
	return &prog, err
}

func (p *Program) Init() {
	for _, v := range []string{"w", "x", "y", "z"} {
		p.variables[v] = 0
	}
}

func IsNumber(s string) bool {
	ok, _ := regexp.MatchString(`-?\d+`, s)
	return ok
}

func (p *Program) Run(input []int) (int, error) {
	p.Init()
	i := 0

	for _, instruction := range p.instructions {
		// fmt.Println(p.variables)
		// fmt.Printf("%v\n\n", instruction)
		fn := instruction.fn
		a := instruction.a.(string)

		if fn == "inp" {
			p.variables[a] = input[i]
			i++
		} else {
			result, err := instruction.Call(&p.variables)
			if err != nil {
				return 0, err
			}
			p.variables[a] = result
		}
	}

	return p.variables["z"], nil
}

func FindLargestModelNumber(p *Program) int {
	var input []int

	for i := 0; i < 14; i++ {
		input = append(input, 9)
	}

	for i := 0; i < 14; i++ {
		for j := 9; j > 0; j-- {
			input[i] = j
			fmt.Println(input)
			result, err := p.Run(input)
			if err != nil {
				log.Fatal(err)
			}
			if result == 0 {
				return toNumber(input)
			}
		}
	}

	return 0
}

func pow(x, m int) int {
	var result int = 1
	for i := 0; i < m; i++ {
		result *= x
	}
	return result
}

func toNumber(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		pos := len(arr) - (i + 1)
		result += arr[i] * pow(10, pos)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	arr, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	prog, err := NewProgram(arr)
	if err != nil {
		log.Fatal(err)
	}

	result1 := FindLargestModelNumber(prog)
	fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := basinSizesMultiplied(&heightmap)
	// fmt.Printf("Puzzle 2: %v\n", result2)
}
