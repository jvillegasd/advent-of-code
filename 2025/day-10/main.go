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

type Node struct {
	Diagram string
	Steps   int
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

func bfs(machine Machine) int {
	stack := []Node{}
	visited := make(map[string]bool)

	emptyDiagram := strings.Replace(machine.lightDiagram, "#", ".", -1)
	stack = append(stack, Node{Diagram: emptyDiagram, Steps: 0})
	visited[emptyDiagram] = true

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.Diagram == machine.lightDiagram {
			return current.Steps
		}

		for _, button := range machine.buttons {
			newDiagram := current.Diagram
			for _, idx := range button.indices {
				switch current.Diagram[idx] {
				case '.':
					newDiagram = newDiagram[:idx] + "#" + newDiagram[idx+1:]
				case '#':
					newDiagram = newDiagram[:idx] + "." + newDiagram[idx+1:]
				}
			}

			if !visited[newDiagram] {
				stack = append(stack, Node{Diagram: newDiagram, Steps: current.Steps + 1})
				visited[newDiagram] = true
			}
		}
	}

	return -1
}

func solvePart1(machines []Machine) int {
	total := 0
	machineSteps := []int{}

	for _, machine := range machines {
		steps := bfs(machine)
		if steps != -1 {
			machineSteps = append(machineSteps, steps)
		}
	}

	for _, steps := range machineSteps {
		total += steps
	}
	return total
}

func main() {
	machines := []Machine{}

	file, err := os.Open("2025/day-10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		machine := parseInput(line)
		machines = append(machines, machine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := solvePart1(machines)
	fmt.Println("Part 1: ", part1)
}
