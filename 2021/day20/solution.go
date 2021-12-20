package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) (ImageEnchancement, error) {
	file, err := os.Open(filename)
	if err != nil {
		return ImageEnchancement{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	if err != nil {
		return ImageEnchancement{}, err
	}
	algo, err := parsePixels(line)
	if err != nil {
		return ImageEnchancement{}, err
	}

	// skip blank line
	scanner.Scan()

	var pixels []Pixels
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			return ImageEnchancement{}, err
		}
		row, err := parsePixels(line)
		if err != nil {
			return ImageEnchancement{}, err
		}
		pixels = append(pixels, row)
	}

	err = scanner.Err()

	image := NewImage()
	image.FromSlice(pixels)

	return ImageEnchancement{algo, image}, err
}

func parsePixels(line string) ([]Pixel, error) {
	var arr []Pixel
	for _, r := range line {
		switch r {
		case '.':
			arr = append(arr, 0)
		case '#':
			arr = append(arr, 1)
		default:
			return nil, fmt.Errorf("wrong input: %v", r)
		}
	}
	return arr, nil
}

// dark pixel '.' = 0
// light pixel '#' = 1
type Pixel byte

func (p Pixel) String() string {
	if p == 0 {
		return "."
	} else {
		return "#"
	}
}

type Pixels []Pixel

func (p Pixels) String() string {
	s := ""
	for _, x := range p {
		s += fmt.Sprintf("%v", x)
	}
	return s
}

func pow(x, m int) int {
	y := 1
	for i := 0; i < m; i++ {
		y *= x
	}
	return y
}

func (p Pixels) ToInt() int {
	dec := 0
	for m, x := range p {
		dec += int(x) * pow(2, len(p)-m-1)
	}
	return dec
}

type Position struct {
	i, j int
}

type Image struct {
	dark                   map[Position]bool
	mini, maxi, minj, maxj int
	outside                Pixel
}

func NewImage() Image {
	dark := make(map[Position]bool)
	return Image{dark, 0, 0, 0, 0, 0}
}

func (image *Image) FromSlice(pixels []Pixels) {
	for i, row := range pixels {
		for j, x := range row {
			image.Add(x, i, j)
		}
	}
}

func (image *Image) Add(p Pixel, i, j int) {
	if p == 1 {
		if i < image.mini {
			image.mini = i
		} else if i > image.maxi {
			image.maxi = i
		}
		if j < image.minj {
			image.minj = j
		} else if j > image.maxj {
			image.maxj = j
		}
		image.dark[Position{i, j}] = true
	}
}

func (image Image) String() string {
	margin := 2
	s := ""
	for i := image.mini - margin; i <= image.maxi+margin; i++ {
		for j := image.minj - margin; j <= image.maxj+margin; j++ {
			s += fmt.Sprintf("%v", image.Get(i, j))
		}
		s += "\n"
	}
	return s
}

func (image Image) OutsideBounds(i, j int) bool {
	return i <= image.mini || i >= image.maxi || j <= image.minj || j >= image.maxj
}

func (image Image) Get(i, j int) Pixel {
	if _, ok := image.dark[Position{i, j}]; ok {
		return 1
	} else if image.OutsideBounds(i, j) {
		return image.outside
	} else {
		return 0
	}
}

func (image Image) Window(i, j int) int {
	var pixels Pixels
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			pixels = append(pixels, image.Get(i+x, j+y))
		}
	}
	return pixels.ToInt()
}

func (image Image) NumDark() int {
	return len(image.dark)
}

type ImageEnchancement struct {
	algo Pixels
	Image
}

func (e *ImageEnchancement) Step() {
	image := NewImage()
	margin := 3
	for i := e.mini - margin; i < e.maxi+margin; i++ {
		for j := e.minj - margin; j < e.maxj+margin; j++ {
			pos := e.Window(i, j)
			image.Add(e.algo[pos], i-e.mini, j-e.minj)
		}
	}
	// The catch: when algorithm starts with '#'
	// the outer border flips at each step
	pos := 511 * int(e.outside)
	e.outside = e.algo[pos]

	e.dark = image.dark
	e.mini = image.mini
	e.maxi = image.maxi
	e.minj = image.minj
	e.maxj = image.maxj
}

func (e *ImageEnchancement) Run(steps int) {
	for i := 0; i < steps; i++ {
		e.Step()
	}
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	filename := os.Args[1]
	image, err := readFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%v\n", image)
	// fmt.Println(image.NumDark())
	// image.Step()
	// fmt.Printf("%v\n", image)
	// fmt.Println(image.NumDark())
	// image.Step()
	// fmt.Printf("%v\n", image)
	// fmt.Println(image.NumDark())

	image.Run(2)
	result1 := image.NumDark()
	fmt.Printf("Puzzle 1: %v\n", result1)

	image.Run(48)
	result2 := image.NumDark()
	fmt.Printf("Puzzle 2: %v\n", result2)

}
