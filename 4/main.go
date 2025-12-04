package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

const (
	ROLL  = "@"
	BLANK = "."
)

type Position struct {
	x, y int
}

func (p Position) IsOccupied(input [][]string) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}

	if p.x >= len(input) || p.y >= len(input[0]) {
		return false
	}

	return input[p.x][p.y] == ROLL
}

func countNeighbours(input [][]string, x, y int) int {
	count := 0

	positions := []Position{
		// lt
		{x - 1, y - 1},
		// t
		{x - 1, y},
		// rt
		{x - 1, y + 1},
		// l
		{x, y - 1},
		// r
		{x, y + 1},
		// lb
		{x + 1, y - 1},
		// b
		{x + 1, y},
		// rb
		{x + 1, y + 1},
	}

	for _, pos := range positions {
		if pos.IsOccupied(input) {
			count += 1
		}
	}

	return count
}

func solve(input [][]string) []Position {
	positions := make([]Position, 0)

	for x := range input {
		for y := range input[x] {
			if input[x][y] != ROLL {
				continue
			}

			neighbours := countNeighbours(input, x, y)
			if neighbours < 4 {
				positions = append(positions, Position{x, y})
			}
		}
	}

	return positions

}

func solveWithCleanup(input [][]string) int {
	count := 0

	var current [][]string
	current = append(current, input...)

	for {
		positions := solve(input)
		if len(positions) == 0 {
			break
		}

		count += len(positions)

		for _, pos := range positions {
			current[pos.x][pos.y] = BLANK
		}
	}

	return count
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
		positions := solve(input)
		log.Printf("Answer: %d\n", len(positions))
		log.Printf("Elapsed time: %v\n", time.Since(start))
	case 2:
		start := time.Now()
		count := solveWithCleanup(input)
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
