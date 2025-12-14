package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	Row int
	Col int
}

type Shape struct {
	Index    int
	Variants [][]Coord
}

type Region struct {
	Width      int
	Height     int
	Quantities []int
}

type Input struct {
	Shapes  []Shape
	Regions []Region
}

func rotate90(coords []Coord) []Coord {
	if len(coords) == 0 {
		return coords
	}

	maxRow := coords[0].Row
	for _, c := range coords {
		if c.Row > maxRow {
			maxRow = c.Row
		}
	}

	rotated := make([]Coord, len(coords))
	for i, c := range coords {
		rotated[i] = Coord{
			Row: c.Col,
			Col: maxRow - c.Row,
		}
	}
	return rotated
}

func mirrorX(coords []Coord) []Coord {
	if len(coords) == 0 {
		return coords
	}

	maxCol := coords[0].Col
	for _, c := range coords {
		if c.Col > maxCol {
			maxCol = c.Col
		}
	}

	mirrored := make([]Coord, len(coords))
	for i, c := range coords {
		mirrored[i] = Coord{
			Row: c.Row,
			Col: maxCol - c.Col,
		}
	}
	return mirrored
}

func normalize(coords []Coord) []Coord {
	if len(coords) == 0 {
		return coords
	}

	minRow := coords[0].Row
	minCol := coords[0].Col
	for _, c := range coords {
		if c.Row < minRow {
			minRow = c.Row
		}
		if c.Col < minCol {
			minCol = c.Col
		}
	}

	// Translate coordinates so minimum is at (0, 0)
	normalized := make([]Coord, len(coords))
	for i, c := range coords {
		normalized[i] = Coord{
			Row: c.Row - minRow,
			Col: c.Col - minCol,
		}
	}

	return normalized
}

func generateVariants(coords []Coord) [][]Coord {
	variants := [][]Coord{}

	bases := [][]Coord{coords, mirrorX(coords)}
	for _, baseCoords := range bases {
		current := baseCoords
		for i := 0; i < 4; i++ {
			normalized := normalize(current)
			variants = append(variants, normalized)
			current = rotate90(current)
		}
	}

	return variants
}

func parseInput(filename string) *Input {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := &Input{
		Shapes:  []Shape{},
		Regions: []Region{},
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, ":") && !strings.Contains(line, "x") {
			parts := strings.Split(line, ":")
			index, _ := strconv.Atoi(parts[0])

			pattern := [][]rune{}
			for scanner.Scan() {
				patternLine := scanner.Text()
				if patternLine == "" {
					break
				}
				pattern = append(pattern, []rune(patternLine))
			}

			// Extract coordinates of '#' cells
			coords := []Coord{}
			for i, row := range pattern {
				for j, cell := range row {
					if cell == '#' {
						coords = append(coords, Coord{Row: i, Col: j})
					}
				}
			}

			variants := generateVariants(coords)
			shape := Shape{
				Index:    index,
				Variants: variants,
			}
			input.Shapes = append(input.Shapes, shape)
		} else if strings.Contains(line, "x") {
			parts := strings.Split(line, ": ")
			dims := strings.Split(parts[0], "x")
			width, _ := strconv.Atoi(dims[0])
			height, _ := strconv.Atoi(dims[1])

			quantityStrs := strings.Fields(parts[1])
			quantities := make([]int, len(quantityStrs))
			for i, qStr := range quantityStrs {
				q, _ := strconv.Atoi(qStr)
				quantities[i] = q
			}

			input.Regions = append(input.Regions, Region{
				Width:      width,
				Height:     height,
				Quantities: quantities,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func main() {
	input := parseInput("2025/day-12/input.txt")
	fmt.Println(input)

	fmt.Println("Part 1:", 0)
	fmt.Println("Part 2:", 0)
}
