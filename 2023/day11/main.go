package main

import (
	"bufio"
	"fmt"
	"os"
)

// Point defined in terms od 2D coordinates
type Point struct {
	x, y int
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (this Point) Sub(other Point) Point {
	return Point{this.x - other.x, this.y - other.y}
}

func (this Point) Abs() Point {
	return Point{Abs(this.x), Abs(this.y)}
}

type SpaceMap struct {
	galaxies               []Point
	galaxy_row, galaxy_col []bool
}

func Parse(file *os.File) SpaceMap {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	var (
		galaxies   []Point
		galaxy_row []bool
		i          int
	)
	galaxy_col := make([]bool, len(line))

	for {
		galaxy_row = append(galaxy_row, false)

		for j, r := range line {
			if r == '#' {
				galaxies = append(galaxies, Point{i, j})
				galaxy_row[i] = true
				galaxy_col[j] = true
			}
		}

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		i++
	}

	return SpaceMap{galaxies, galaxy_row, galaxy_col}
}

func space_between(low, high int, any_galaxies []bool) int {
	if low > high {
		low, high = high, low
	}
	// don't include lower bound
	low++

	count := 0
	if low < high {
		for _, has_galaxy := range any_galaxies[low:high] {
			if !has_galaxy {
				count++
			}
		}
	}
	return count
}

// Distance corrected for the expanded space
func (m SpaceMap) WeightedDistance(a, b Point, weight int) int {
	row_space := space_between(a.x, b.x, m.galaxy_row)
	col_space := space_between(a.y, b.y, m.galaxy_col)

	// Taxicab distance is:
	// dist(A,B) = |A_1 - B_1| + |A_2 + B_2|

	// We create D = |A - B|, where A, B, D are points. D is like the
	// distance before summing the dimensions.
	d := b.Sub(a).Abs()

	// The space expansion correction means that the absolute difference
	// |A_i - B_i| is adjusted by some expansion factor equal to w times
	// the number of empty space fields between them. It becomes:
	//
	//   #non_empty_points + w * #empty_space_points.
	//
	// The same is applied to both dimensions.
	// Since for the distance we care is only about the absolute differences,
	// we can create the D = |A - B| point, re-scale it's coordinates
	// accordingly, and sum them.
	return (d.x - row_space) + row_space*weight + (d.y - col_space) + col_space*weight
}

func (m SpaceMap) TotalDistance(weight int) int {
	result := 0
	for i := 0; i < len(m.galaxies); i++ {
		for j := i + 1; j < len(m.galaxies); j++ {
			a := m.galaxies[i]
			b := m.galaxies[j]
			result += m.WeightedDistance(a, b, weight)
		}
	}
	return result
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	space_map := Parse(file)

	fmt.Println(space_map.TotalDistance(2))
	fmt.Println(space_map.TotalDistance(10))
	fmt.Println(space_map.TotalDistance(100))
	fmt.Println(space_map.TotalDistance(1000000))
}
