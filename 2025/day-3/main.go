package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInput(line string) []int {
	var batteryBanks []int
	for _, char := range line {
		batteryBanks = append(batteryBanks, int(char-'0'))
	}
	return batteryBanks
}

func getMaxJoltage(line string) int {
	maxBatteryIdx := 0
	batteryBanks := parseInput(line)
	secondMaxBatteryIdx := len(batteryBanks) - 1

	for i := 1; i < len(batteryBanks)-1; i++ {
		if batteryBanks[i] > batteryBanks[maxBatteryIdx] {
			maxBatteryIdx = i
		}
	}

	for i := len(batteryBanks) - 2; i > maxBatteryIdx; i-- {
		if batteryBanks[i] > batteryBanks[secondMaxBatteryIdx] {
			secondMaxBatteryIdx = i
		}
	}

	maxJoltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", batteryBanks[maxBatteryIdx], batteryBanks[secondMaxBatteryIdx]))
	return maxJoltage
}

func main() {
	sum := 0

	file, err := os.Open("2025/day-3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getMaxJoltage(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", sum)
}
