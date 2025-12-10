package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func parseInput(line string) Point {
	parts := strings.Split(line, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return Point{x, y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solvePart1(points []Point) int {
	maxArea := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i].X == points[j].X || points[i].Y == points[j].Y {
				continue
			}

			width := abs(points[j].X-points[i].X) + 1
			height := abs(points[j].Y-points[i].Y) + 1
			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	points := []Point{}
	file, err := os.Open("2025/day-9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, parseInput(line))
	}

	part1 := solvePart1(points)
	fmt.Println("Part 1: ", part1)
}
