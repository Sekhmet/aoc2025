package main

import (
	"flag"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	low, high int
}

func (r Range) IsIncluded(id int) bool {
	return r.low <= id && id <= r.high
}

func (r Range) Count() int {
	return r.high - r.low + 1
}

func solveLevel1(ranges []Range, ids []int) int {
	count := 0

	for _, id := range ids {
		for _, r := range ranges {
			if r.IsIncluded(id) {
				count += 1
				break
			}
		}
	}

	return count
}

// Intervals are merged together before counting.
// https://www.rishabhxchoudhary.com/blog/overlapping-interval-problems
func solveLevel2(ranges []Range) int {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.low - b.low
	})

	unique_ranges := make([]Range, 0)

	for _, r := range ranges {
		if len(unique_ranges) == 0 || r.low > unique_ranges[len(unique_ranges)-1].high {
			unique_ranges = append(unique_ranges, r)
		} else {
			unique_ranges[len(unique_ranges)-1].high = max(unique_ranges[len(unique_ranges)-1].high, r.high)
		}
	}

	count := 0

	for _, r := range unique_ranges {
		count += r.Count()
	}

	return count
}

func readInputFile(path string) ([]Range, []int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	ranges_list := strings.Split(lines[0], "\n")
	ids_list := strings.Split(lines[1], "\n")

	ranges := make([]Range, len(ranges_list))
	for i, r := range ranges_list {
		split := strings.Split(r, "-")
		low, _ := strconv.Atoi(split[0])
		high, _ := strconv.Atoi(split[1])

		ranges[i] = Range{low, high}
	}

	ids := make([]int, len(ids_list))
	for i, id := range ids_list {
		parsed, _ := strconv.Atoi(id)
		ids[i] = parsed
	}

	return ranges, ids, nil
}

func main() {
	var level = flag.Int("level", 1, "level to solve")
	var inputFile = flag.String("input", "input.txt", "input file path")

	flag.Parse()

	ranges, ids, err := readInputFile(*inputFile)
	if err != nil {
		log.Println("Error reading input file:", err)
		os.Exit(1)
	}

	switch *level {
	case 1:
		start := time.Now()
		count := solveLevel1(ranges, ids)
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	case 2:
		start := time.Now()
		count := solveLevel2(ranges)
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
