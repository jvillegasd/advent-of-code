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
	answer_part_2 := 0

	file, err := os.Open("2025/day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		movement, distance := parseInput(line)

		full_rotations := distance / MOD
		remainingDistance := mod(distance, MOD)
		answer_part_2 += full_rotations

		switch movement {
		case "L":
			delta := dial - remainingDistance
			dial = mod(delta, MOD)

			if delta < 0 {
				// Edge case: Dial was at 0 before moving left, so that is not a real rotation
				if dial+remainingDistance != MOD {
					answer_part_2++
				}
			} else if dial == 0 {
				answer_part_2++
			}
		case "R":
			delta := dial + remainingDistance
			dial = mod(delta, MOD)

			if delta > MOD {
				// Edge case: Dial was at MOD before moving right, so that is not a real rotation
				if dial-remainingDistance != 0 {
					answer_part_2++
				}
			} else if dial == 0 {
				answer_part_2++
			}
		}

		if dial == 0 {
			answer++
		}
	}

	fmt.Println("Part 1: ", answer)
	fmt.Println("Part 2: ", answer_part_2)
}
