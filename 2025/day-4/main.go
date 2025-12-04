package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Roll struct {
	i, j int
}

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countAdjacentRolls(grid [][]rune, i, j int) int {
	rolls := 0
	for _, direction := range directions {
		newI, newJ := i+direction[0], j+direction[1]
		if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[i]) {
			continue
		}
		if grid[newI][newJ] == '@' {
			rolls++
		}
	}
	return rolls
}

func buildRollsQueue(grid [][]rune) []Roll {
	queue := []Roll{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && countAdjacentRolls(grid, i, j) < 4 {
				queue = append(queue, Roll{i, j})
			}
		}
	}
	return queue
}

func bfs(grid [][]rune, queue []Roll) int {
	totalRemoved := 0

	for len(queue) > 0 {
		roll := queue[0]
		queue = queue[1:]

		if grid[roll.i][roll.j] != '@' {
			continue
		}

		if countAdjacentRolls(grid, roll.i, roll.j) >= 4 {
			continue
		}

		grid[roll.i][roll.j] = '#'
		totalRemoved++

		for _, direction := range directions {
			newI, newJ := roll.i+direction[0], roll.j+direction[1]
			if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[roll.i]) {
				continue
			}

			if grid[newI][newJ] == '@' && countAdjacentRolls(grid, newI, newJ) < 4 {
				queue = append(queue, Roll{newI, newJ})
			}
		}
	}

	return totalRemoved
}

func main() {
	var grid [][]rune

	file, err := os.Open("2025/day-4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	queue := buildRollsQueue(grid)

	count := len(queue)
	totalRemoved := bfs(grid, queue)

	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", totalRemoved)
}
