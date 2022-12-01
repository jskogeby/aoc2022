package utils

import (
	"bufio"
	"log"
	"os"
)

// read from input/ folder and return a slice of strings
func ReadInput(filename string) []string {
	// open file
	file, err := os.Open("input/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// divide array of strings at empty lines
func DivideInput(input []string) [][]string {
	var divided [][]string
	var current []string
	for _, line := range input {
		if line == "" {
			divided = append(divided, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}
	divided = append(divided, current)

	return divided
}
