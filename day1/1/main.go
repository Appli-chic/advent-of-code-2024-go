package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList := getInput()
	result := calculate(leftList, rightList)
	print(result)
}

func getInput() ([]int, []int) {
	file, err := os.Open("day1/1/input.txt")
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
	if len(leftList) != len(rightList) {
		os.Exit(1)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	result := 0
	for listIndex := 0; listIndex < len(leftList); listIndex++ {
		result += calculateDistance(leftList[listIndex], rightList[listIndex])
	}

	return result
}

func calculateDistance(leftNumber int, rightNumber int) int {
	if leftNumber > rightNumber {
		return leftNumber - rightNumber
	}
	return rightNumber - leftNumber
}
