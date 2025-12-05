package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

func main() {
	countFreshIds := 0
	ranges := []Range{}
	availableIds := []int64{}
	afterBlankLine := false

	file, err := os.Open("2025/day-5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			afterBlankLine = true
		}

		if !afterBlankLine {
			parts := strings.Split(line, "-")
			start, _ := strconv.ParseInt(parts[0], 10, 64)
			end, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, Range{start, end})
		} else {
			id, _ := strconv.ParseInt(line, 10, 64)
			availableIds = append(availableIds, id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, r := range ranges {
		idsToRemove := []int64{}
		for _, id := range availableIds {
			if id >= r.start && id <= r.end {
				countFreshIds++
				idsToRemove = append(idsToRemove, id)
			}
		}
		for _, idToRemove := range idsToRemove {
			availableIds = slices.DeleteFunc(availableIds, func(currentId int64) bool {
				return currentId == idToRemove
			})
		}
	}

	fmt.Println("Part 1: ", countFreshIds)
}
