package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile rune

const (
	Space  Tile = '.'
	Galaxy Tile = '#'
)

func (tile Tile) String() string {
	return string(tile)
}

type SpaceMap struct {
	image                [][]Tile
	space_row, space_col []bool
	galaxies             []Point
}

func (this SpaceMap) Print() {
	for _, row := range this.image {
		fmt.Println(row)
	}
}

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

func (this Point) Add(other Point) Point {
	return Point{this.x + other.x, this.y + other.y}
}

func (this Point) Sub(other Point) Point {
	return Point{this.x - other.x, this.y - other.y}
}

func (this Point) Abs() Point {
	return Point{Abs(this.x), Abs(this.y)}
}

func (this SpaceMap) ExpandedSpaceBetween(a, b Point) (int, int) {
	var low, high int

	low, high = a.x, b.x
	if low > high {
		low, high = high, low
	}
	row_count := 0
	if low+1 < high {
		for _, expandable := range this.space_row[low+1 : high] {
			if expandable {
				row_count++
			}
		}
	}

	low, high = a.y, b.y
	if low > high {
		low, high = high, low
	}
	col_count := 0
	if low+1 < high {
		for _, expandable := range this.space_col[low+1 : high] {
			if expandable {
				col_count++
			}
		}
	}

	return row_count, col_count
}

// Taxicab distance between two points
func (this Point) Distance(other Point) int {
	return Abs(this.x-other.x) + Abs(this.y-other.y)
}

// Distance corrected for the expanded space
func (this SpaceMap) WeightedDistance(a, b Point, weight int) int {
	row_space, col_space := this.ExpandedSpaceBetween(a, b)
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

func Galaxies(image [][]Tile) []Point {
	var galaxies []Point
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			this := image[i][j]
			if this == Galaxy {
				galaxies = append(galaxies, Point{i, j})
			}
		}
	}
	return galaxies
}

func ExpandableSpace(image [][]Tile) ([]bool, []bool) {
	space_row := make([]bool, len(image))
	for i := 0; i < len(image); i++ {
		all_space := true
		for j := 0; j < len(image[i]); j++ {
			if image[i][j] != '.' {
				all_space = false
			}
		}
		if all_space {
			space_row[i] = true
		}
	}

	space_col := make([]bool, len(image[0]))
	for j := 0; j < len(image[0]); j++ {
		all_space := true
		for i := 0; i < len(image); i++ {
			if image[i][j] != '.' {
				all_space = false
			}
		}
		if all_space {
			space_col[j] = true
		}
	}

	return space_row, space_col
}

func parseLine(line string) []Tile {
	var row []Tile
	for _, r := range line {
		tile := Tile(r)
		row = append(row, tile)
	}
	return row
}

func parse(file *os.File) SpaceMap {
	scanner := bufio.NewScanner(file)

	var image [][]Tile

	for scanner.Scan() {
		line := scanner.Text()
		row := parseLine(line)
		image = append(image, row)
	}

	galaxies := Galaxies(image)
	space_row, space_col := ExpandableSpace(image)

	return SpaceMap{image, space_row, space_col, galaxies}
}

func TotalDistance(space_map SpaceMap, weight int) int {
	result := 0
	for i := 0; i < len(space_map.galaxies); i++ {
		for j := i + 1; j < len(space_map.galaxies); j++ {
			a := space_map.galaxies[i]
			b := space_map.galaxies[j]
			result += space_map.WeightedDistance(a, b, weight)
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

	space_map := parse(file)
	// space_map.Print()

	fmt.Println(TotalDistance(space_map, 2))
	fmt.Println(TotalDistance(space_map, 10))
	fmt.Println(TotalDistance(space_map, 100))
	fmt.Println(TotalDistance(space_map, 1000000))
}
