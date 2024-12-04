package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2024/day/3
func main() {
	firstPart := firstPart()
	fmt.Printf("Result of fist part: %d\n", firstPart)

	secondPart := secondPart()
	fmt.Printf("Result of second part: %d\n", secondPart)
}

// Reads the puzzle file and returns its content as a string. If the file can't
// be read, the function returns an empty string and an error.
func readFile() (string, error) {
	var result string

	inputFile, err := os.ReadFile("../input")
	if err != nil {
		return result, errors.New("unable to read puzzle file")
	}

	return string(inputFile), err
}

// First part.
func firstPart() int {
	var result int
	input, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)

	for _, v := range matches {
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matchedNums := re.FindStringSubmatch(v)

		num1, _ := strconv.Atoi(matchedNums[1])
		num2, _ := strconv.Atoi(matchedNums[2])

		result += num1 * num2
	}

	return result
}

// Second part.
func secondPart() int {
	var result int
	_, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	return result
}
