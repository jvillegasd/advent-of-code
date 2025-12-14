package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph map[string][]string

func parseInput(line string) (string, []string) {
	parts := strings.Split(line, ":")
	node := parts[0]
	outputs := strings.Split(parts[1], " ")

	return node, outputs
}

func dfs(graph Graph, node string, visited map[string]bool) int {
	if node == "out" {
		return 1
	}

	paths := 0
	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			paths += dfs(graph, neighbor, visited)
		}
	}
	return paths
}

func main() {
	graph := Graph{}

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

	part1 := dfs(graph, "you", make(map[string]bool))
	fmt.Println("Part 1: ", part1)
}
