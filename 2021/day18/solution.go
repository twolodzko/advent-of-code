package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Action int

const (
	EXPLODE Action = iota
	SPLIT
	NOOP
)

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return arr, err
		}
		arr = append(arr, line)
	}
	err = scanner.Err()
	return arr, err
}

type Pair [2]interface{}

func IsPair(obj interface{}) bool {
	switch obj.(type) {
	case Pair:
		return true
	default:
		return false
	}
}

func (p Pair) String() string {
	return fmt.Sprintf("[%v,%v]", p[0], p[1])
}

func (x Pair) Add(y Pair) Pair {
	return Pair{x, y}
}

func getAction(obj interface{}, depth int) Action {

	// pair, isPair := obj.(Pair)
	// if !isPair {
	// 	if num, ok := obj.(int); ok && num >= 10 {
	// 		return SPLIT
	// 	}
	// 	return NOOP
	// }

	// if depth == 0 {
	// 	leftmost := pair[0]
	// 	getAction(pair[1])
	// }

	// switch obj := obj.(type) {
	// case Pair:
	// 	right := getAction(obj[0])
	// 	left := getAction(obj[1])
	// default:
	// 	return NOOP
	// }

	// switch depth {
	// case 0:
	// case 2:
	// 	return EXPLODE
	// default:
	// 	if !IsPair(obj) {
	// 		return NOOP
	// 	} else {
	// 		return getAction()
	// 	}
	// }

	return NOOP
}

func Explode(pair Pair) Pair {
	return Pair{}
}

func Split(num int) Pair {
	left := int(math.Floor(float64(num) / 2))
	right := int(math.Ceil(float64(num) / 2))
	return Pair{left, right}
}

type Reader struct {
	*strings.Reader
	Head rune
}

func (cr *Reader) NextRune() error {
	r, _, err := cr.ReadRune()
	if err != nil {
		cr.Head = rune(0)
		return err
	}
	cr.Head = r
	return nil
}

func (r *Reader) ReadPair() (Pair, error) {
	if r.Head != '[' {
		return Pair{}, fmt.Errorf("not a pair start character: %v", r.Head)
	}
	err := r.NextRune()
	if err != nil {
		return Pair{}, err
	}

	var fields [2]interface{}
	for i := 0; i < 2; i++ {
		field, err := r.ReadField()
		if err != nil {
			return Pair{}, err
		}
		fields[i] = field
		err = r.NextRune()
		if err != nil && err != io.EOF {
			return Pair{}, err
		}
	}
	return Pair{fields[0], fields[1]}, nil
}

func (r *Reader) ReadField() (interface{}, error) {
	var field []rune
	for {
		switch r.Head {
		case ',', ']':
			return strconv.Atoi(string(field))
		case '[':
			return r.ReadPair()
		default:
			field = append(field, r.Head)
		}

		err := r.NextRune()
		if err != nil {
			return nil, err
		}
	}
}

func ParsePair(str string) (Pair, error) {
	reader := Reader{strings.NewReader(str), rune(0)}
	err := reader.NextRune()
	if err != nil {
		return Pair{}, err
	}
	return reader.ReadPair()
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	lines, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		pair, err := ParsePair(line)
		if err != nil {
			log.Fatal(err)
		}
		if pair.String() != line {
			log.Fatalf("Parsing error: %v != %v", pair, line)
		}
	}

	// result1 := packet.VersionNumberSum()
	// fmt.Printf("Puzzle 1: %v\n", result1)

	// result2 := packet.Value()
	// fmt.Printf("Puzzle 2: %v\n", result2)

}
