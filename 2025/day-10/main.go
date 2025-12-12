package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Button struct {
	indices []int
}

type Machine struct {
	lightDiagram string
	buttons      []Button
	joltage      []int
}

func parseInput(line string) Machine {
	parts := strings.Split(line, " ")

	// Get the light diagram
	lightDiagram := strings.Replace(parts[0], "[", "", 1)
	lightDiagram = strings.Replace(lightDiagram, "]", "", 1)

	// Get the joltage
	joltageString := strings.Replace(parts[len(parts)-1], "{", "", 1)
	joltageString = strings.Replace(joltageString, "}", "", 1)
	joltageParts := strings.Split(joltageString, ",")
	joltageInts := []int{}
	for _, part := range joltageParts {
		joltageInt, _ := strconv.Atoi(part)
		joltageInts = append(joltageInts, joltageInt)
	}

	// parts[1:-1] are the buttons
	buttons := []Button{}
	for _, part := range parts[1 : len(parts)-1] {
		buttonStr := strings.TrimPrefix(part, "(")
		buttonStr = strings.TrimSuffix(buttonStr, ")")
		buttonParts := strings.Split(buttonStr, ",")

		indices := []int{}
		for _, bp := range buttonParts {
			idx, _ := strconv.Atoi(bp)
			indices = append(indices, idx)
		}

		buttons = append(buttons, Button{indices: indices})
	}

	return Machine{lightDiagram, buttons, joltageInts}
}

func main() {
	file, err := os.Open("2025/day-10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		machine := parseInput(line)
		fmt.Println(machine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
