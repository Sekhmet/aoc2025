package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

const (
	START           = "S"
	SPLITTER        = "^"
	SPLITTER_ACTIVE = "%"
	BEAM            = "|"
	BLANK           = "."
)

func solve(input [][]string) int {
	splits := 0

	for r := range input {
		for c := range input[r] {
			if r == 0 && input[r][c] == START {
				input[r+1][c] = BEAM
			}
			if input[r][c] == BLANK && r > 0 && input[r-1][c] == BEAM {
				input[r][c] = BEAM
			}

			if input[r][c] == SPLITTER && input[r-1][c] == BEAM {
				splits += 1
				input[r][c-1] = BEAM
				input[r][c+1] = BEAM
			}
		}

	}

	return splits
}

func solveWithAlt(input [][]string) int {
	values := make([][]int, len(input))

	for r := range input {
		values[r] = make([]int, len(input[r]))
	}

	for r := range input {
		for c := range input[r] {
			if r == 0 && input[r][c] == START {
				values[r][c] = 1
			} else if input[r][c] == BLANK && r > 0 {
				values[r][c] += values[r-1][c]
			} else if input[r][c] == SPLITTER {
				values[r][c-1] += values[r-1][c]
				values[r][c+1] += values[r-1][c]
			}
		}

	}

	total := 0
	for _, v := range values[len(values)-1] {
		total += v
	}

	return total
}

func readInputFile(path string) ([][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	input := make([][]string, len(lines))

	for i, line := range lines {
		input[i] = strings.Split(line, "")
	}

	return input, nil
}

func main() {
	var level = flag.Int("level", 1, "level to solve")
	var inputFile = flag.String("input", "input.txt", "input file path")

	flag.Parse()

	input, err := readInputFile(*inputFile)
	if err != nil {
		log.Println("Error reading input file:", err)
		os.Exit(1)
	}

	switch *level {
	case 1:
		start := time.Now()
		count := solve(input)
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	case 2:
		start := time.Now()
		count := solveWithAlt(input)
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}
}
