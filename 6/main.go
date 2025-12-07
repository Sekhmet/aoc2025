package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	SYMBOL_ADD      = "+"
	SYMBOL_MULTIPLY = "*"
)

type Entry []string

func getStartingValue(op string) int {
	if op == SYMBOL_MULTIPLY {
		return 1
	}

	return 0
}

func _(input [][]int, operands []string) int {
	total := 0

	for j, op := range operands {
		count := getStartingValue(op)

		for i := range len(input) {
			if op == SYMBOL_ADD {
				count += input[i][j]
			}

			if op == SYMBOL_MULTIPLY {
				count *= input[i][j]
			}
		}

		total += count
	}

	return total
}

func readInputFile(path string) ([][]string, []string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	operands := make([]string, 0)
	edges := make([]int, 0)
	for i, f := range lines[len(lines)-1] {
		if f != ' ' {
			operands = append(operands, string(f))
			edges = append(edges, i)
		}
	}

	log.Printf("Operands: %v\n", operands)
	log.Printf("Edges: %v\n", edges)

	input := make([][]string, len(lines)-1)

	total := 0

	for i, op := range operands {
		count := getStartingValue(op)
		start := edges[i]

		end := len(lines[0])
		if i != len(edges)-1 {
			end = edges[i+1] - 1
		}

		fragments := make([]string, 0)
		for j := range len(lines) - 1 {
			fragment := string(lines[j][start:end])
			fragments = append(fragments, fragment)
		}

		for j := range len(fragments[0]) {
			v := ""
			for _, f := range fragments {
				v += string(f[j])
			}

			vv, _ := strconv.Atoi(strings.Trim(v, " "))
			if op == SYMBOL_ADD {
				count += vv
			}

			if op == SYMBOL_MULTIPLY {
				count *= vv
			}
		}

		total += count

	}

	log.Printf("Total: %d\n", total)

	return input, operands, nil
}

func main() {
	var level = flag.Int("level", 1, "level to solve")
	var inputFile = flag.String("input", "input.txt", "input file path")

	flag.Parse()

	start := time.Now()
	_, _, err := readInputFile(*inputFile)
	if err != nil {
		log.Println("Error reading input file:", err)
		os.Exit(1)
	}

	switch *level {
	case 1:
		count := 0
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	case 2:
		count := 0
		log.Printf("Answer: %d\n", count)
		log.Printf("Elapsed time: %v\n", time.Since(start))
	default:
		log.Println("Unknown level")
		os.Exit(1)
	}

}
