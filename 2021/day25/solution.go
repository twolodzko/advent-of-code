package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	EMPTY = '.'
	EAST  = '>'
	SOUTH = 'v'
)

func readFile(filename string) (Region, error) {
	var arr [][]MaybeCucumber

	file, err := os.Open(filename)
	if err != nil {
		return Region{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []MaybeCucumber
		for _, r := range line {
			row = append(row, MaybeCucumber{r})
		}
		arr = append(arr, row)
	}
	err = scanner.Err()
	return Region{arr}, err
}

type MaybeCucumber struct {
	kind rune
}

func (c MaybeCucumber) String() string {
	return string(c.kind)
}

func (c MaybeCucumber) Occupied() bool {
	return c.kind == EAST || c.kind == SOUTH
}

func (x MaybeCucumber) Equal(y MaybeCucumber) bool {
	return x.kind == y.kind
}

type Region struct {
	region [][]MaybeCucumber
}

func (x Region) Equal(y Region) bool {
	n := len(x.region)
	k := len(x.region[n-1])

	if n != len(y.region) || k != len(y.region[n-1]) {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if x.region[i][j] != y.region[i][j] {
				return false
			}
		}
	}
	return true
}

func (r Region) String() string {
	s := ""
	for _, row := range r.region {
		for _, r := range row {
			s += r.String()
		}
		s += "\n"
	}
	return s
}

func (r Region) Copy() Region {
	var new [][]MaybeCucumber
	n := len(r.region)
	k := len(r.region[n-1])
	for i := 0; i < n; i++ {
		var row []MaybeCucumber
		for j := 0; j < k; j++ {
			row = append(row, MaybeCucumber{r.region[i][j].kind})
		}
		new = append(new, row)
	}
	return Region{new}
}

func (r *Region) Step() Region {
	new := r.Copy()
	for _, kind := range []rune{EAST, SOUTH} {
		new.region = new.Move(kind)
	}
	return new
}

func (r Region) Move(kind rune) [][]MaybeCucumber {
	n := len(r.region)
	k := len(r.region[n-1])

	new := r.Copy().region
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			this := r.region[i][j]
			if this.kind == kind {
				ni, nj := r.Next(i, j)
				if !r.region[ni][nj].Occupied() {
					new[ni][nj].kind = this.kind
					new[i][j].kind = EMPTY
				}
			}
		}
	}
	return new
}

func (r Region) Next(i, j int) (int, int) {
	n := len(r.region)
	k := len(r.region[n-1])

	switch r.region[i][j].kind {
	case SOUTH:
		return (i + 1) % n, j
	case EAST:
		return i, (j + 1) % k
	default:
		return i, j
	}
}

func (r *Region) Run() int {
	step := 0
	for {
		step++
		new := r.Step()
		if r.Equal(new) {
			break
		}
		r.region = new.region
	}
	return step
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	region, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1 := region.Run()
	fmt.Printf("Puzzle 1: %v\n", result1)

}
