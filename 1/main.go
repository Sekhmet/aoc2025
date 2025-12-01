package main

import (
	"flag"
	"fmt"
	"github.com/Sekhmet/aoc2025/toolkit"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	STARTING_POSITION = 50
	CLICKS            = 100
)

func solveLevel1(input []string) {
	log.Println("Solving level 1")

	p := STARTING_POSITION
	c := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		direction, distanceStr := line[:1], line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			continue
		}

		switch direction {
		case "L":
			p -= distance
		case "R":
			p += distance
		default:
			log.Printf("Unknown direction: %s\n", direction)
		}

		p = toolkit.Pmod(p, CLICKS)

		if p == 0 {
			c += 1
		}
	}

	log.Printf("Final position: %d, Number of 0 positions: %d\n", p, c)
}

func solveLevel2(input []string) {
	log.Println("Solving level 2")

	p := STARTING_POSITION
	c := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		direction, distanceStr := line[:1], line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			continue
		}

		// Good thing Go is fast, I gave up doing math
		switch direction {
		case "L":
			for range distance {
				p = toolkit.Pmod(p-1, CLICKS)

				if p == 0 {
					c += 1
				}
			}
		case "R":
			for range distance {
				p = toolkit.Pmod(p+1, CLICKS)

				if p == 0 {
					c += 1
				}
			}
		default:
			log.Printf("Unknown direction: %s\n", direction)
		}
	}

	log.Printf("Final position: %d, Number of 0 positions: %d\n", p, c)
}

func readInputFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
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
		solveLevel1(input)
	case 2:
		solveLevel2(input)
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
