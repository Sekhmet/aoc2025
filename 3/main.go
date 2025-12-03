package main

import (
	"flag"
	"github.com/Sekhmet/aoc2025/toolkit"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func getJoltage(input string, num_batteries int) int {
	digits := make([]int, len(input))
	for i, v := range input {
		digit, _ := strconv.Atoi(string(v))
		digits[i] = digit
	}

	joltage := 0

	start_index := 0
	for i := range num_batteries {
		available := digits[start_index : len(input)-(num_batteries-1-i)]

		max_index := toolkit.FindMaxIndex(available)
		max_value := available[max_index]

		joltage += max_value * int(math.Pow10(num_batteries-1-i))

		start_index += max_index + 1
	}

	return joltage
}

func solve(input []string, num_batteries int) {
	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		total += getJoltage(line, num_batteries)
	}

	log.Println("Total joltage", total)

}

func readInputFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n"), nil
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
		solve(input, 2)
	case 2:
		solve(input, 12)
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
