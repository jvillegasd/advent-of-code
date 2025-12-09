package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

type Edge struct {
	From, To int
	Distance float64
}

type UnionFind struct {
	parent []int
	size   []int
}

func initUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y int) bool {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX == rootY {
		return false
	}

	// Union by size
	if uf.size[rootX] < uf.size[rootY] {
		rootX, rootY = rootY, rootX
	}

	uf.parent[rootY] = rootX
	uf.size[rootX] += uf.size[rootY]
	return true
}

func (uf *UnionFind) getSize(x int) int {
	return uf.size[uf.find(x)]
}

func distance(p1, p2 Point) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	dz := float64(p1.Z - p2.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parsePoint(line string) Point {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		log.Fatalf("Invalid point format: %s", line)
	}

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])
	return Point{x, y, z}
}

func solvePart1(n int, edges []Edge) int {
	totalConnections := 0
	uf := initUnionFind(n)
	circuitSizes := []int{}
	seenRoots := make(map[int]bool)

	for _, edge := range edges {
		if totalConnections >= 1000 {
			break
		}

		uf.union(edge.From, edge.To)
		totalConnections++
	}

	for i := 0; i < n; i++ {
		root := uf.find(i)
		if !seenRoots[root] {
			seenRoots[root] = true
			circuitSizes = append(circuitSizes, uf.getSize(root))
		}
	}

	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})

	return circuitSizes[0] * circuitSizes[1] * circuitSizes[2]
}

func solvePart2(n int, points []Point, edges []Edge) int {
	uf := initUnionFind(n)
	var lastMergingEdge Edge

	for _, edge := range edges {
		if uf.union(edge.From, edge.To) {
			lastMergingEdge = edge
			if uf.getSize(0) == n {
				break
			}
		}
	}

	return points[lastMergingEdge.From].X * points[lastMergingEdge.To].X
}

func main() {
	edges := []Edge{}
	points := []Point{}

	file, err := os.Open("2025/day-8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, parsePoint(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distance(points[i], points[j])
			edges = append(edges, Edge{i, j, dist})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Distance < edges[j].Distance
	})

	part1 := solvePart1(n, edges)
	part2 := solvePart2(n, points, edges)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
