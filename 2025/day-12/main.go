package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Shape struct {
	Index   int
	Pattern [][]rune
	Width   int
	Height  int
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

func rotate90(pattern [][]rune) [][]rune {
	if len(pattern) == 0 {
		return pattern
	}

	height := len(pattern)
	width := len(pattern[0])

	// Create rotated pattern: width becomes height, height becomes width
	rotated := make([][]rune, width)
	for i := range rotated {
		rotated[i] = make([]rune, height)
	}

	// Rotate: element at [i][j] goes to [j][height-1-i]
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			rotated[j][height-1-i] = pattern[i][j]
		}
	}

	return rotated
}

func mirrorX(pattern [][]rune) [][]rune {
	if len(pattern) == 0 {
		return pattern
	}

	height := len(pattern)
	width := len(pattern[0])

	// Create mirrored pattern
	mirrored := make([][]rune, height)
	for i := range mirrored {
		mirrored[i] = make([]rune, width)
	}

	// Mirror: element at [i][j] goes to [i][width-1-j]
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			mirrored[i][width-1-j] = pattern[i][j]
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

			input.Shapes = append(input.Shapes, Shape{
				Index:   index,
				Pattern: pattern,
				Width:   width,
				Height:  height,
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
