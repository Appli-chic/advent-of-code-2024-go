package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList := getInput()
	result := calculate(leftList, rightList)
	print(result)
}

func getInput() ([]int, []int) {
	file, err := os.Open("day1/2/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	leftList := make([]int, 0)
	rightList := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		leftNumber, leftNumberError := strconv.Atoi(fields[0])
		rightNumber, rightNumberError := strconv.Atoi(fields[1])
		if leftNumberError != nil || rightNumberError != nil {
			fmt.Println("Left error:", leftNumberError, "Right error:", rightNumberError)
			os.Exit(1)
		}

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	return leftList, rightList
}

func calculate(leftList []int, rightList []int) int {
	result := 0

	for _, leftItem := range leftList {
		count := 0
		for _, rightItem := range rightList {
			if leftItem == rightItem {
				count += 1
			}
		}

		result += leftItem * count
	}

	return result
}
