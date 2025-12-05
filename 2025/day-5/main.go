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

func mergeRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		if a.start != b.start {
			return int(a.start - b.start)
		}
		return int(a.end - b.end)
	})

	mergedRanges := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &mergedRanges[len(mergedRanges)-1]
		current := ranges[i]

		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}

	return mergedRanges
}

func countFreshIds(ranges []Range, availableIds []int64) int64 {
	count := int64(0)
	for _, id := range availableIds {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				count++
				break
			}
		}
	}
	return count
}

func countContiguousIds(ranges []Range) int64 {
	count := int64(0)
	for _, r := range ranges {
		count += r.end - r.start + 1
	}
	return count
}

func main() {
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

	mergedRanges := mergeRanges(ranges)
	countFreshIds := countFreshIds(mergedRanges, availableIds)
	countContiguousIds := countContiguousIds(mergedRanges)

	fmt.Println("Part 1: ", countFreshIds)
	fmt.Println("Part 2: ", countContiguousIds)
}
