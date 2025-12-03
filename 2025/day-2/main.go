package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseRange(rangeStr string) (string, string) {
	parts := strings.Split(rangeStr, "-")
	start := parts[0]
	end := parts[1]
	return start, end
}

func findInvalidIDsSum(start, end string) int64 {
	sum := int64(0)

	startInt, _ := strconv.Atoi(start)
	endInt, _ := strconv.Atoi(end)

	var halfStart int
	if len(start)/2 > 0 {
		halfStart, _ = strconv.Atoi(start[:len(start)/2])
	} else {
		halfStart, _ = strconv.Atoi(start)
	}

	var halfEnd int
	if (len(end)+1)/2 > 0 {
		halfEnd, _ = strconv.Atoi(end[:(len(end)+1)/2])
	} else {
		halfEnd, _ = strconv.Atoi(end)
	}

	for i := halfStart; i <= halfEnd; i++ {
		potentialID, _ := strconv.Atoi(fmt.Sprintf("%d%d", i, i))
		if potentialID >= startInt && potentialID <= endInt {
			sum += int64(potentialID)
		}
	}

	return sum
}

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

func areChunksEqual(chunks []string) bool {
	if len(chunks) <= 1 {
		return false
	}

	firstChunk := chunks[0]
	for i := 1; i < len(chunks); i++ {
		if chunks[i] != firstChunk {
			return false
		}
	}

	return true
}

func isInvalidID2(number int) bool {
	numberStr := strconv.Itoa(number)
	for i := 1; i <= len(numberStr)/2; i++ {
		chunks := chunkString(numberStr, i)
		if areChunksEqual(chunks) {
			return true
		}
	}

	return false
}

func main() {
	sum := int64(0)
	sum2 := int64(0)
	var rangeStrings []string

	file, err := os.Open("2025/day-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		rangeStrings = strings.Split(line, ",")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, rangeStr := range rangeStrings {
		start, end := parseRange(strings.TrimSpace(rangeStr))
		sum += findInvalidIDsSum(start, end)

		startInt, _ := strconv.Atoi(start)
		endInt, _ := strconv.Atoi(end)
		for i := startInt; i <= endInt; i++ {
			if isInvalidID2(i) {
				sum2 += int64(i)
			}
		}
	}

	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", sum2)
}
