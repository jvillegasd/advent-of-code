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

func getMaxJoltage2(line string) int64 {
	idx := 0
	maxBatteryIdxs := make([]int, 12)
	batteryBanks := parseInput(line)

	for i := 0; i < len(maxBatteryIdxs); i++ {
		maxIdx := idx
		digitsLeft := len(maxBatteryIdxs) - 1 - i
		for j := idx + 1; j < len(batteryBanks)-digitsLeft; j++ {
			if batteryBanks[j] > batteryBanks[maxIdx] {
				maxIdx = j
			}
		}
		maxBatteryIdxs[i] = maxIdx
		idx = maxIdx + 1
	}

	var resultStr string
	for _, idx := range maxBatteryIdxs {
		resultStr += strconv.Itoa(batteryBanks[idx])
	}
	resultInt, _ := strconv.ParseInt(resultStr, 10, 64)
	return resultInt
}

func main() {
	sum := 0
	sum2 := int64(0)

	file, err := os.Open("2025/day-3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getMaxJoltage(line)
		sum2 += getMaxJoltage2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
