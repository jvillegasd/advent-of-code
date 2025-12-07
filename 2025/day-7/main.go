package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	count1 := 0
	count2 := 1 // Base timeline 'S'
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

	lastBeams := make([]bool, len(diagram[0]))
	lastBeansTimelines := make([]int, len(diagram[0]))
	for i := 0; i < len(diagram[0]); i++ {
		if diagram[0][i] == 'S' {
			lastBeams[i] = true
			lastBeansTimelines[i] = 1
		}
	}

	currentBeams := lastBeams
	for i := 1; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			if diagram[i][j] == '^' && lastBeams[j] {
				count1++
				count2 += lastBeansTimelines[j]

				if j+1 < len(diagram[i]) {
					currentBeams[j+1] = true
					lastBeansTimelines[j+1] += lastBeansTimelines[j]
				}
				if j-1 >= 0 {
					currentBeams[j-1] = true
					lastBeansTimelines[j-1] += lastBeansTimelines[j]
				}

				currentBeams[j] = false
				lastBeansTimelines[j] = 0
			}
		}
		lastBeams = currentBeams
	}

	fmt.Println("Part 1: ", count1)
	fmt.Println("Part 2: ", count2)
}
