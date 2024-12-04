package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2
func main() {
	firstPart := firstPart()
	fmt.Printf("Result of fist part: %d\n", firstPart)

	secondPart := secondPart()
	fmt.Printf("Result of second part: %d\n", secondPart)
}

// readFile reads puzzle file and parses it into a 2D array of integers
// each row in the array represents a single line from the file.
// The function returns an empty array and an error if the file can't be read.
func readFile() ([][]int, error) {
	result := [][]int{}

	inputFile, err := os.Open("../input")
	if err != nil {
		return result, errors.New("unable to read puzzle file")
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		row := []int{}

		for j := 0; j < len(nums); j++ {
			level, _ := strconv.Atoi(nums[j])

			row = append(row, level)
		}

		result = append(result, row)
	}
	inputFile.Close()

	return result, err
}

// First part.
func firstPart() int {
	levels, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	var result int
	for i := 0; i < len(levels); i++ {
		err := validateDifficulties(levels[i])

		if err != nil {
			// Incorrect difficulties
			continue
		}

		result += 1
	}

	return result
}

// Second part.
func secondPart() int {
	levels, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	var result int
	for i := 0; i < len(levels); i++ {
		err := validateDifficulties(levels[i])

		if err == nil {
			result += 1
			continue
		}

		// Incorrect difficulties, try to remove elements one by one and reuse already created functions
		for skippedIndex := range levels[i] {

			// first skip
			if skippedIndex == 0 {
				tmp := levels[i][1:]
				err := validateDifficulties(tmp)
				if err == nil {
					result += 1
					break
				} else {
					continue
				}
			}

			// Check all other combinations
			tmpFirst := levels[i][:skippedIndex]
			tmpLast := levels[i][skippedIndex+1:]

			var tmp []int
			tmp = append(tmp, tmpFirst...)
			tmp = append(tmp, tmpLast...)

			err := validateDifficulties(tmp)
			if err == nil {
				result += 1
				break
			}
		}
	}

	return result
}

// Validates if the levels are in same order (increasing or decreasing),
// and if the difference between the levels is in range of 1 to 3.
// If the levels are valid, the function returns nil, otherwise it returns an error.
func validateDifficulties(levels []int) error {
	difficultyIncreases := []int{}
	diff := 0
	lastIndex := len(levels) - 1

	for i, val := range levels {

		if i == lastIndex {
			break
		}

		diff = val - levels[i+1]
		difficultyIncreases = append(difficultyIncreases, diff)
	}

	increasesCount, decreasesCount, misstakes := 0, 0, 0
	for _, val := range difficultyIncreases {

		if val == 0 || val < -3 || val > 3 {
			return errors.New("elements out of range")
		}

		if val < 0 {
			increasesCount++
		} else if val > 0 {
			decreasesCount++
		} else {
			misstakes++
		}
	}

	// If valid sequence then no error, i use lastindex becouse the last element is not checked
	if increasesCount == lastIndex || decreasesCount == lastIndex {
		return nil
	}

	return errors.New("incorrect input, not all elements in same order")
}
