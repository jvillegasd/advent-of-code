package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(line string) (int, int) {
	parts := strings.Split(line, "-")
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return start, end
}

func main() {
	file, err := os.Open("2025/day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		start, end := parseInput(line)
	}
}
