package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func applyOperation(operation string, a []int) int64 {
	if len(a) == 0 {
		return int64(0)
	}

	result := int64(a[0])
	for i := 1; i < len(a); i++ {
		switch operation {
		case "*":
			result *= int64(a[i])
		case "+":
			result += int64(a[i])
		}
	}
	return result
}

func main() {
	sum := int64(0)
	operations := [][]string{}

	file, err := os.Open("2025/day-6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Use Fields to split on whitespace and remove empty strings
		operations = append(operations, strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(operations[0]); i++ {
		a, _ := strconv.Atoi(operations[0][i])
		b, _ := strconv.Atoi(operations[1][i])
		c, _ := strconv.Atoi(operations[2][i])
		d, _ := strconv.Atoi(operations[3][i])
		numbers := []int{a, b, c, d}
		sum += int64(applyOperation(operations[4][i], numbers))
	}

	fmt.Println("Part 1: ", sum)
}
