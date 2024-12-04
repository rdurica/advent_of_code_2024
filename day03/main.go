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
	input, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(input, -1)

	return calculateMultiplication(matches)
}

// Second part.
func secondPart() int {
	var result int
	input, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	re := regexp.MustCompile(`(?:mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := re.FindAllString(input, -1)
	do := true
	for _, val := range matches {

		if val == "don't()" {
			do = false
			continue
		} else if val == "do()" {
			do = true
			continue
		}

		if !do {
			continue
		}

		re := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := re.FindAllString(val, -1)

		result += calculateMultiplication(matches)
	}

	return result
}

// Takes a list of strings in the form of "mul(num1, num2)" and
// returns the sum of all the multiplications of the numbers in the strings.
func calculateMultiplication(data []string) int {
	var result int

	for _, val := range data {
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matchedNums := re.FindStringSubmatch(val)

		num1, _ := strconv.Atoi(matchedNums[1])
		num2, _ := strconv.Atoi(matchedNums[2])

		result += num1 * num2
	}

	return result
}
