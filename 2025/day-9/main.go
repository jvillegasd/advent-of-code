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

type Line struct {
	Start, End Point
}

type RectangleBounds struct {
	MinX, MaxX, MinY, MaxY int
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

func isPointOnBoundary(point Point, h_lines []Line, v_lines []Line) bool {
	// Check if the point is on a horizontal line
	for _, line := range h_lines {
		if point.Y == line.Start.Y && point.X >= line.Start.X && point.X <= line.End.X {
			return true
		}
	}

	// Check if the point is on a vertical line
	for _, line := range v_lines {
		if point.X == line.Start.X && point.Y >= line.Start.Y && point.Y <= line.End.Y {
			return true
		}
	}
	return false
}

func isPointInsidePolygon(point Point, polygon []Line) bool {
	// Ray casting algorithm
	n := len(polygon)
	intersections := 0

	for i := 0; i < n; i++ {
		line := polygon[i]

		// Check if the ray intersects the edge
		if (line.Start.Y <= point.Y && line.End.Y > point.Y) || (line.Start.Y > point.Y && line.End.Y <= point.Y) {
			// Calculate the x-coordinate where the ray crosses the edge's line
			intersectionX := line.Start.X + (point.Y-line.Start.Y)*(line.End.X-line.Start.X)/(line.End.Y-line.Start.Y)

			// If the intersection is to the right of the point, count it
			if intersectionX > point.X {
				intersections++
			}
		}
	}

	return intersections%2 != 0
}

func getRectangleBounds(point1 Point, point2 Point) RectangleBounds {
	minX := point1.X
	maxX := point2.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}

	minY := point1.Y
	maxY := point2.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	return RectangleBounds{minX, maxX, minY, maxY}
}

func polygonIntersectsRectangle(bounds RectangleBounds, h_lines []Line, v_lines []Line) bool {
	// Check if any horizontal line intersects the rectangle interior
	for _, line := range h_lines {
		// Normalize line endpoints (ensure Start.X <= End.X)
		lineMinX := line.Start.X
		lineMaxX := line.End.X
		if lineMinX > lineMaxX {
			lineMinX, lineMaxX = lineMaxX, lineMinX
		}

		// Horizontal line intersects rectangle interior if:
		// 1. Line's Y is strictly inside rectangle's Y range
		// 2. Line's X range overlaps with rectangle's X range
		if line.Start.Y > bounds.MinY && line.Start.Y < bounds.MaxY {
			if !(lineMaxX < bounds.MinX || lineMinX > bounds.MaxX) {
				return true
			}
		}
	}

	// Check if any vertical line intersects the rectangle interior
	for _, line := range v_lines {
		// Normalize line endpoints (ensure Start.Y <= End.Y)
		lineMinY := line.Start.Y
		lineMaxY := line.End.Y
		if lineMinY > lineMaxY {
			lineMinY, lineMaxY = lineMaxY, lineMinY
		}

		// Vertical line intersects rectangle interior if:
		// 1. Line's X is strictly inside rectangle's X range
		// 2. Line's Y range overlaps with rectangle's Y range
		if line.Start.X > bounds.MinX && line.Start.X < bounds.MaxX {
			if !(lineMaxY < bounds.MinY || lineMinY > bounds.MaxY) {
				return true
			}
		}
	}

	return false
}

func isValidPoint(point Point, polygon []Line, h_lines []Line, v_lines []Line) bool {
	return isPointInsidePolygon(point, polygon) || isPointOnBoundary(point, h_lines, v_lines)
}

func solvePart2(points []Point) int {
	maxArea := 0
	n := len(points)
	polygon := []Line{}
	v_lines := []Line{}
	h_lines := []Line{}

	// Build the polygon and track vertical and horizontal lines
	for i := 0; i < n; i++ {
		point1 := points[i]
		point2 := points[(i+1)%n] // Wrap around to the first point
		line := Line{point1, point2}
		polygon = append(polygon, line)

		if point1.X == point2.X {
			v_lines = append(v_lines, line)
		} else {
			h_lines = append(h_lines, line)
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			point1 := points[i]
			point2 := points[j]
			bounds := getRectangleBounds(point1, point2)

			area := (bounds.MaxX - bounds.MinX + 1) * (bounds.MaxY - bounds.MinY + 1)
			if area <= maxArea {
				continue
			}

			if polygonIntersectsRectangle(bounds, h_lines, v_lines) {
				continue
			}

			cornerPoint1 := Point{bounds.MinX, bounds.MaxY}
			cornerPoint2 := Point{bounds.MaxX, bounds.MinY}

			isValid := isValidPoint(cornerPoint1, polygon, h_lines, v_lines) && isValidPoint(cornerPoint2, polygon, h_lines, v_lines)
			if !isValid {
				continue
			}

			maxArea = area
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
	part2 := solvePart2(points)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
