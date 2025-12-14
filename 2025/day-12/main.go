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
	Index  int
	Coords []Coord
	Width  int
	Height int
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

func rotate90(coords []Coord, height int) []Coord {
	rotated := make([]Coord, len(coords))
	for i, c := range coords {
		rotated[i] = Coord{
			Row: c.Col,
			Col: height - 1 - c.Row,
		}
	}
	return rotated
}

func mirrorX(coords []Coord, width int) []Coord {
	mirrored := make([]Coord, len(coords))
	for i, c := range coords {
		mirrored[i] = Coord{
			Row: c.Row,
			Col: width - 1 - c.Col,
		}
	}
	return mirrored
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

			width := 0
			if len(pattern) > 0 {
				width = len(pattern[0])
			}
			height := len(pattern)

			// Extract coordinates of '#' cells
			coords := []Coord{}
			for i, row := range pattern {
				for j, cell := range row {
					if cell == '#' {
						coords = append(coords, Coord{Row: i, Col: j})
					}
				}
			}

			input.Shapes = append(input.Shapes, Shape{
				Index:  index,
				Coords: coords,
				Width:  width,
				Height: height,
			})
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
