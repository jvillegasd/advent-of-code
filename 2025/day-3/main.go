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

func getMaxJoltage(line string, size int) int64 {
	idx := 0
	maxBatteryIdxs := make([]int, size)
	batteryBanks := parseInput(line)

	for i := 0; i < size; i++ {
		maxIdx := idx
		digitsLeft := size - 1 - i
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
	sum := int64(0)
	sum2 := int64(0)

	file, err := os.Open("2025/day-3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getMaxJoltage(line, 2)
		sum2 += getMaxJoltage(line, 12)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
