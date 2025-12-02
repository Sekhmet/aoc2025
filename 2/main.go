package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

func validateLevel1(input int) bool {
	str := strconv.Itoa(input)
	part1, part2 := str[:len(str)/2], str[len(str)/2:]

	return part1 != part2
}

func validateLevel2(input int) bool {
	str := strconv.Itoa(input)

	for i := range len(str) / 2 {
		sample, remainder := str[:i+1], str[i+1:]

		if len(remainder)%len(sample) != 0 {
			continue
		}

		target := strings.Repeat(sample, len(remainder)/len(sample))

		if remainder == target {
			return false
		}
	}

	return true
}

func solve(input []string, validate func(int) bool) {
	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		start, end, found := strings.Cut(line, "-")
		if !found {
			log.Fatalln("Invalid range", line)
			return
		}

		startInt, err := strconv.Atoi(start)
		if err != nil {
			log.Fatalln("Error parsing start of range", err)
			continue
		}

		endInt, err := strconv.Atoi(end)
		if err != nil {
			log.Fatalln("Error parsing end of range", err)
			continue
		}

		for i := startInt; i <= endInt; i++ {
			if !validate(i) {
				total += i
			}
		}
	}

	log.Println("Total sum of matching numbers:", total)

}

func readInputFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(data)), ","), nil
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
		solve(input, validateLevel1)
	case 2:
		solve(input, validateLevel2)
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
