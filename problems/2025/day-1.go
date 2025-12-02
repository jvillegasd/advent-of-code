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

	file, err := os.Open("inputs/2025/day-1.txt")
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
			dial = mod(dial-distance, MOD)
		case "R":
			dial = mod(dial+distance, MOD)
		}

		if dial == 0 {
			answer = answer + 1
		}
	}

	fmt.Println(answer)
}
