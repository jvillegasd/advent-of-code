package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func parseInput(line string) (string, int) {
	movement := string(line[0])
	distance, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}

	return movement, int(distance)
}

func main() {
	dial := 50
	MOD := 100
	answer := 0
	rotation := 0

	file, err := os.Open("2025/day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		movement, distance := parseInput(line)
		switch movement {
		case "L":
			delta := dial - distance
			dial = mod(delta, MOD)

			if delta < 0 {
				rotation++
			}
		case "R":
			delta := dial + distance
			dial = mod(delta, MOD)

			if delta > MOD {
				rotation++
			}
		}

		if dial == 0 {
			answer++
		}

		div := distance / MOD
		rotation += div
	}

	fmt.Println("Part 1: ", answer)
	fmt.Println("Part 2: ", rotation)
}
