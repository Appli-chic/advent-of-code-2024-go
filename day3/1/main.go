package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input := getInput()
	result := calculate(input)
	print(result)
}

func calculate(input string) int {
	result := 0
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		if len(match) == 3 {
			firstNumber, _ := strconv.Atoi(match[1])
			secondNumber, _ := strconv.Atoi(match[2])
			result += firstNumber * secondNumber
		}
	}

	return result
}

func getInput() string {
	content, err := os.ReadFile("day3/1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return string(content)
}
