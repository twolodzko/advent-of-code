package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x, y int
}

type Rectangle struct {
	xmin, xmax, ymin, ymax int
}

func (r Rectangle) Contains(p Point) bool {
	return p.x >= r.xmin && p.x <= r.xmax && p.y >= r.ymin && p.y <= r.ymax
}

func parse(s string) (Rectangle, error) {
	fields := regexp.MustCompile(`-?\d+`).FindAllString(s, -1)
	var arr []int
	for _, r := range fields {
		x, err := strconv.Atoi(r)
		if err != nil {
			return Rectangle{}, err
		}
		arr = append(arr, x)
	}
	return Rectangle{arr[0], arr[1], arr[2], arr[3]}, nil
}

//  - The probe's x position increases by its x velocity.
//
//  - The probe's y position increases by its y velocity.
//
//  - Due to drag, the probe's x velocity changes by 1 toward the value 0; that is,
//  it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0,
//  or does not change if it is already 0.
//  - Due to gravity, the probe's y velocity decreases by 1.
//
//   dy = dy - n*(n+1)/2
//   y = n*dy - n*(n+1)/2
//   y = 0  at  (n-1)/2
//
//   dx = dx - n*(n+1)/2   x > 0
//   dx = 0                x = 0
//   dx = dx + n*(n+1)/2   x < 0
//
type Velocity struct {
	x, y int
}

func (v *Velocity) Next() (int, int) {
	if v.x > 0 {
		v.x--
	} else if v.x < 0 {
		v.x++
	}
	v.y--
	return v.x, v.y
}

type Probe struct {
	pos  Point
	vel  Velocity
	path []Point
}

func NewProbe(v Velocity) Probe {
	return Probe{Point{0, 0}, v, nil}
}

func (p *Probe) Next() {
	p.pos.x += p.vel.x
	p.pos.y += p.vel.y
	p.vel.Next()
	p.path = append(p.path, p.pos)
}

func (p *Probe) MaxY() int {
	max := math.MinInt
	for _, point := range p.path {
		if point.y > max {
			max = point.y
		}
	}
	return max
}

func (p *Probe) Fire(r Rectangle) bool {
	for {
		// fmt.Printf("x=%4d  y=%4d  vx=%4d  vy%4d\n", p.pos.x, p.pos.y, p.vel.x, p.vel.y)
		if p.pos.x > r.xmax || p.pos.y < r.ymin {
			// overshoot
			return false
		}
		if r.Contains(p.pos) {
			// hit
			return true
		}
		p.Next()
	}
}

func simulateProbes(xmin, xmax, ymin, ymax int, target Rectangle) map[Velocity]int {
	results := make(map[Velocity]int)
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			vel := Velocity{x, y}
			probe := NewProbe(vel)
			ok := probe.Fire(target)
			if ok {
				results[vel] = probe.MaxY()
			}
		}
	}
	return results
}

func findHighestFlyingProbe(results map[Velocity]int) int {
	ybest := math.MinInt
	for _, ymax := range results {
		if ymax > ybest {
			ybest = ymax
		}
	}
	return ybest
}

func howManyInitsExist(results map[Velocity]int) int {
	return len(results)
}

func example() {
	targetArea, err := parse("target area: x=20..30, y=-10..-5")
	if err != nil {
		log.Fatal(err)
	}

	for _, point := range [][]int{{7, 2}, {6, 3}, {9, 0}, {17, -4}} {
		probe := NewProbe(Velocity{point[0], point[1]})
		ok := probe.Fire(targetArea)
		max := probe.MaxY()

		fmt.Printf("Probe: %v\n", probe)
		fmt.Printf("Success=%v  Highest=%d\n\n", ok, max)
	}

	results := simulateProbes(-500, 500, -500, 500, targetArea)

	result1 := findHighestFlyingProbe(results)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := howManyInitsExist(results)
	fmt.Printf("Puzzle 2: %v\n", result2)
}

func main() {
	example()

	targetArea, err := parse("target area: x=60..94, y=-171..-136")
	if err != nil {
		log.Fatal(err)
	}

	results := simulateProbes(-500, 500, -500, 500, targetArea)

	result1 := findHighestFlyingProbe(results)
	fmt.Printf("Puzzle 1: %v\n", result1)

	result2 := howManyInitsExist(results)
	fmt.Printf("Puzzle 2: %v\n", result2)
}
