package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var strategies = map[string]func([]int) int64{
	"*": func(a []int) int64 {
		result := int64(1)
		for _, num := range a {
			result *= int64(num)
		}
		return result
	},
	"+": func(a []int) int64 {
		result := int64(0)
		for _, num := range a {
			result += int64(num)
		}
		return result
	},
}

func solvePart1(operations [][]string) int64 {
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

	operationsIdx := len(operations) - 1
	for i := 0; i < len(operations[0]); i++ {
		operation := string(operations[operationsIdx][i])
		if operation != "*" && operation != "+" {
			continue
		}

		currentIdx := i
		numbers := []int{}
		for currentIdx < len(operations[0]) {
			newNumber := ""
			isAllBlank := true
			for j := 0; j < len(operations)-1; j++ {
				if operations[j][currentIdx] == ' ' {
					continue
				}
				isAllBlank = false
				newNumber += string(operations[j][currentIdx])
			}

			if isAllBlank {
				break
			}

			number, _ := strconv.Atoi(newNumber)
			numbers = append(numbers, number)
			currentIdx++
		}

		sum += strategies[operation](numbers)
	}

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
