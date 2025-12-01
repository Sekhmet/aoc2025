package main

import (
	"flag"
	"log"
	"os"
	"strings"
)

func solveLevel1(_ []string) {
	log.Println("Solving level 1")
}

func solveLevel2(_ []string) {
	log.Println("Solving level 2")
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
