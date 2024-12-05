package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2024/day/4
func main() {
	firstPart := firstPart()
	fmt.Printf("Result of fist part: %d\n", firstPart)

	secondPart := secondPart()
	fmt.Printf("Result of second part: %d\n", secondPart)
}

// readFile reads puzzle file and parses it into a 2D array of integers
// each row in the array represents a single line from the file.
// The function returns an empty array and an error if the file can't be read.
func readFile() ([][]string, error) {
	result := [][]string{}

	inputFile, err := os.Open("../input")
	if err != nil {
		return result, errors.New("unable to read puzzle file")
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "")
		row := []string{}

		for j := 0; j < len(nums); j++ {
			row = append(row, nums[j])
		}

		result = append(result, row)
	}
	inputFile.Close()

	return result, err
}

// First part.
func firstPart() int {
	input, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	const SEARCH string = "XMAS"
	var result int

	maxY := len(input) - 1
	for y, row := range input {
		maxX := len(row) - 1
		for x := range row {

			// Start of word
			if input[y][x] != "X" {
				continue
			}

			// IDEA: make univerasl method, give index of search char, row (or only few ?), search word and all other indexes calculate based on len()

			// to right
			if x+3 <= maxX {
				word := input[y][x] + input[y][x+1] + input[y][x+2] + input[y][x+3]

				if word == SEARCH {
					result += 1
				}
			}

			// to left
			if x-3 >= 0 {
				word := input[y][x] + input[y][x-1] + input[y][x-2] + input[y][x-3]

				if word == SEARCH {
					result += 1
				}
			}

			// to down
			if y+3 <= maxY {
				word := input[y][x] + input[y+1][x] + input[y+2][x] + input[y+3][x]

				if word == SEARCH {
					result += 1
				}
			}

			// to up
			if y-3 >= 0 {
				word := input[y][x] + input[y-1][x] + input[y-2][x] + input[y-3][x]

				if word == SEARCH {
					result += 1
				}
			}

			// to down left
			if x-3 >= 0 && y+3 <= maxY {
				word := input[y][x] + input[y+1][x-1] + input[y+2][x-2] + input[y+3][x-3]

				if word == SEARCH {
					result += 1
				}
			}

			// to down right
			if x+3 <= maxX && y+3 <= maxY {
				word := input[y][x] + input[y+1][x+1] + input[y+2][x+2] + input[y+3][x+3]

				if word == SEARCH {
					result += 1
				}
			}

			// to up left
			if x-3 >= 0 && y-3 >= 0 {
				word := input[y][x] + input[y-1][x-1] + input[y-2][x-2] + input[y-3][x-3]

				if word == SEARCH {
					result += 1
				}
			}

			// to up right
			if x+3 <= maxX && y-3 >= 0 {
				word := input[y][x] + input[y-1][x+1] + input[y-2][x+2] + input[y-3][x+3]

				if word == SEARCH {
					result += 1
				}
			}
		}
	}

	return result
}

// Second part.
func secondPart() int {
	input, err := readFile()
	if err != nil {
		panic(err.Error())
	}

	const SEARCH string = "MAS"
	var result int

	maxY := len(input) - 1
	for y, row := range input {
		maxX := len(row) - 1
		for x := range row {

			// Start of word
			if input[y][x] != "A" {
				continue
			}

			found := false

			// to down right
			if x-1 >= 0 && y-1 >= 0 && y+1 <= maxY && x+1 <= maxX {
				word := input[y-1][x-1] + input[y][x] + input[y+1][x+1]
				reversed := input[y+1][x+1] + input[y][x] + input[y-1][x-1]

				if word == SEARCH || reversed == SEARCH {
					found = true
				}
			}

			// to down left
			if x+1 <= maxX && y-1 >= 0 && y+1 <= maxY && x-1 >= 0 {
				word := input[y-1][x+1] + input[y][x] + input[y+1][x-1]
				reversed := input[y+1][x-1] + input[y][x] + input[y-1][x+1]

				if (word == SEARCH || reversed == SEARCH) && found {
					result += 1
				}
			}
		}
	}

	return result
}
