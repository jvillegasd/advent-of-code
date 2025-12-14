package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph map[string][]string

type memoKey struct {
	node string
	dac  bool
	fft  bool
}

func parseInput(line string) (string, []string) {
	parts := strings.Split(line, ":")
	node := parts[0]
	outputs := strings.Split(parts[1], " ")

	return node, outputs
}

func dfs(graph Graph, node string) int {
	if node == "out" {
		return 1
	}

	paths := 0
	for _, neighbor := range graph[node] {
		paths += dfs(graph, neighbor)
	}
	return paths
}

func dfs2(graph Graph, node string, dac bool, fft bool, memo map[memoKey]int) int {
	key := memoKey{node, dac, fft}
	if val, exists := memo[key]; exists {
		return val
	}

	switch node {
	case "dac":
		dac = true
	case "fft":
		fft = true
	case "out":
		if dac && fft {
			memo[key] = 1
			return 1
		}
		memo[key] = 0
		return 0
	}

	paths := 0
	for _, neighbor := range graph[node] {
		paths += dfs2(graph, neighbor, dac, fft, memo)
	}

	memo[key] = paths
	return paths
}

func main() {
	graph := Graph{}
	memo := map[memoKey]int{}

	file, err := os.Open("2025/day-11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		node, outputs := parseInput(line)
		graph[node] = outputs
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := dfs(graph, "you")
	part2 := dfs2(graph, "svr", false, false, memo)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
