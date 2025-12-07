package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePart1(operations [][]string) int64 {
	strategies := map[string]func([]int) int64{
		"*": func(a []int) int64 {
			return int64(a[0]) * int64(a[1]) * int64(a[2]) * int64(a[3])
		},
		"+": func(a []int) int64 {
			return int64(a[0]) + int64(a[1]) + int64(a[2]) + int64(a[3])
		},
	}

	sum := int64(0)
	for i := 0; i < len(operations[0]); i++ {
		a, _ := strconv.Atoi(operations[0][i])
		b, _ := strconv.Atoi(operations[1][i])
		c, _ := strconv.Atoi(operations[2][i])
		d, _ := strconv.Atoi(operations[3][i])

		numbers := []int{a, b, c, d}
		sum += strategies[operations[4][i]](numbers)
	}
	return sum
}

func solvePart2(operations []string) int64 {
	sum := int64(0)
	return sum
}

func main() {
	operationsPart1 := [][]string{}
	operationsPart2 := []string{}

	file, err := os.Open("2025/day-6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Use Fields to split on whitespace and remove empty strings
		operationsPart1 = append(operationsPart1, strings.Fields(line))
		operationsPart2 = append(operationsPart2, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := solvePart1(operationsPart1)
	part2 := solvePart2(operationsPart2)

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
