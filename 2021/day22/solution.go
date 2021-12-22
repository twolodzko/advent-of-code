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

func readFile(filename string) ([]Step, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var steps []Step
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("invalid input: %v", fields)
		}

		state := fields[0] == "on"
		var xyz []Range
		for _, field := range strings.Split(fields[1], ",") {
			values := regexp.MustCompile(`-?\d+`).FindAllString(field, -1)
			if len(values) != 2 {
				return nil, fmt.Errorf("invalid input: %v", field)
			}

			min, err := strconv.Atoi(values[0])
			if err != nil {
				return nil, err
			}
			max, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, err
			}
			xyz = append(xyz, Range{min, max})
		}
		steps = append(steps, Step{Cube{xyz[0], xyz[1], xyz[2]}, state})
	}
	err = scanner.Err()
	return steps, err
}

type Range struct {
	min, max int
}

func (r Range) Size() int64 {
	return int64(r.max-r.min) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func (x Range) Intersect(y Range) (Range, bool) {
	low := max(x.min, y.min)
	high := min(x.max, y.max)
	return Range{low, high}, low <= high
}

func (x Range) Equal(y Range) bool {
	return x.min == y.min && x.max == y.max
}

func (r Range) Iterate() <-chan int {
	c := make(chan int)
	go func() {
		for i := r.min; i <= r.max; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

type Cube struct {
	x, y, z Range
}

func (a Cube) Intersect(b Cube) (Cube, bool) {
	var xyz [3]Range
	for i, r := range [][]Range{{a.x, b.x}, {a.y, b.y}, {a.z, b.z}} {
		value, ok := r[0].Intersect(r[1])
		if !ok {
			return Cube{}, false
		}
		xyz[i] = value
	}
	return Cube{xyz[0], xyz[1], xyz[2]}, true
}

type Step struct {
	Cube
	state bool
}

func (s Step) GetCube() Cube {
	return Cube{s.x, s.y, s.z}
}

func (a Cube) Equal(b Cube) bool {
	for _, r := range [][]Range{{a.x, b.x}, {a.y, b.y}, {a.z, b.z}} {
		if r[0] != r[1] {
			return false
		}
	}
	return true
}

func (s Step) OutsideOfInit() bool {
	for _, r := range []Range{s.x, s.y, s.z} {
		if r.min < -50 || r.min > 50 || r.max < -50 || r.max > 50 {
			return true
		}
	}
	return false
}

func (c Cube) Size() int64 {
	return c.x.Size() * c.y.Size() * c.z.Size()
}

type Point struct {
	x, y, z int
}

type Cuboid struct {
	cubes map[Point]bool
}

func NewCuboid() Cuboid {
	return Cuboid{make(map[Point]bool)}
}

func (c *Cuboid) Set(point Point, value bool) {
	if value == true {
		c.cubes[point] = value
	} else {
		// value=false -> delete key
		if _, ok := c.cubes[point]; ok {
			delete(c.cubes, point)
		}
	}
}

func (c *Cuboid) Run(step Step) {
	for x := range step.x.Iterate() {
		for y := range step.y.Iterate() {
			for z := range step.z.Iterate() {
				c.Set(Point{x, y, z}, step.state)
			}
		}
	}
}

func (c *Cuboid) SimpleInit(steps []Step) int {
	for _, step := range steps {
		if step.OutsideOfInit() {
			continue
		}
		// fmt.Printf("Running step %d: %v (size=%d)\n", i+1, step, step.Size())
		c.Run(step)
	}
	return c.CountActive()
}

func (c Cuboid) CountActive() int {
	return len(c.cubes)
}

func (c Cuboid) FullInit(steps []Step) int64 {

	var (
		areas     []Cube
		holes     []Cube
		antiholes []Cube
	)

	for _, step := range steps {
		cube := step.GetCube()
		fmt.Println(cube)

		if step.state {

			// if intersects with hole, fill the hole not to subtract it twice
			for _, hole := range holes {
				intersection, ok := cube.Intersect(hole)
				if ok {
					fmt.Printf(" antihole => %v\n", intersection)
					antiholes = append(antiholes, intersection)
				}
			}
			// if intersects with other area, cut off the intersection
			for _, area := range areas {
				intersection, ok := cube.Intersect(area)
				if ok {
					fmt.Printf(" hole => %v\n", intersection)
					holes = append(holes, intersection)
				}
			}
			// add the area
			areas = append(areas, cube)

		} else {

			// if intersects with hole, fill the hole not to subtract it twice
			for _, hole := range holes {
				intersection, ok := cube.Intersect(hole)
				fmt.Printf(" antihole => %v\n", intersection)
				if ok {
					antiholes = append(antiholes, intersection)
				}
			}
			// if intersects with area, punch a hole in the area
			for _, area := range areas {
				intersection, ok := cube.Intersect(area)
				if ok {
					fmt.Printf(" hole => %v\n", intersection)
					holes = append(holes, intersection)
				}
			}

		}
		fmt.Println()
	}

	var total int64 = 0
	for _, area := range areas {
		total += area.Size()
	}
	for _, hole := range holes {
		total -= hole.Size()
	}
	for _, antihole := range antiholes {
		total += antihole.Size()
	}
	return total
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	steps, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cuboid := NewCuboid()
	result1 := cuboid.SimpleInit(steps)
	fmt.Printf("Puzzle 1: %v\n", result1)

	cuboid = NewCuboid()
	result2 := cuboid.FullInit(steps)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
