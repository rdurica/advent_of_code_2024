package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	firstPart()
	secondPart()
}

// firstPart reads pairs of integers from a file, sorts them, and computes the
// sum of the absolute differences between corresponding elements from two lists.
func firstPart() {
	inputFile, err := os.Open("input")
	if err != nil {
		fmt.Printf("Unable to read puzzle file: %s\n", err)
		return
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	first := []int{}
	second := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		// skip errors, inputs must be correct
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		first = append(first, x)
		second = append(second, y)
	}
	inputFile.Close()

	slices.Sort(first)
	slices.Sort(second)

	var sum int = 0
	for i := 0; i < len(first); i++ {
		distance := first[i] - second[i]
		if distance < 0 {
			sum += distance * -1
		} else {
			sum += distance
		}
	}

	fmt.Printf("Result of fist part: %d\n", sum)
}

// secondPart reads pairs of integers from a file, and computes the sum of the
// occurances of each number in the first column multiplied by the number itself.
func secondPart() {
	inputFile, err := os.Open("input")
	if err != nil {
		fmt.Printf("Unable to read puzzle file: %s", err)
		return
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	first := []int{}
	occurances := map[int]int{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		// skip errors, inputs must be correct
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		first = append(first, x)
		occurances[y] = occurances[y] + 1
	}
	inputFile.Close()

	var sum int = 0
	for i := 0; i < len(first); i++ {
		x := first[i]
		count := occurances[x]

		sum += x * count
	}

	fmt.Printf("Result of second part: %d\n", sum)
}
