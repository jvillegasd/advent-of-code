package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePart1(diagram []string) int {
	count := 0

	lastBeams := make([]bool, len(diagram[0]))
	for i := 0; i < len(diagram[0]); i++ {
		if diagram[0][i] == 'S' {
			lastBeams[i] = true
		}
	}

	currentBeams := lastBeams
	for i := 1; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] == '^' && lastBeams[j] {
				currentBeams[j] = false
				count++

				if j+1 < len(diagram[i]) {
					currentBeams[j+1] = true
				}
				if j-1 >= 0 {
					currentBeams[j-1] = true
				}
			}
		}
		lastBeams = currentBeams
	}

	return count
}

func main() {
	diagram := []string{}

	file, err := os.Open("2025/day-7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		diagram = append(diagram, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := solvePart1(diagram)

	fmt.Println("Part 1: ", part1)
}
