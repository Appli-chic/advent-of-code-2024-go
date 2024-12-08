package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := getInput()
	result := calculate(reports)
	fmt.Println(result)
}

func getInput() [][]int {
	file, err := os.Open("day2/2/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	inputList := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		levels := make([]int, 0)
		for _, field := range fields {
			number, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				os.Exit(1)
			}
			levels = append(levels, number)
		}
		inputList = append(inputList, levels)
	}

	return inputList
}

func calculate(reports [][]int) int {
	numberSafeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			numberSafeReports++
		} else {
			if checkSafetyWithOneMistake(report) {
				numberSafeReports++
			}
		}
	}

	return numberSafeReports
}

func checkSafetyWithOneMistake(report []int) bool {
	for index, _ := range report {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:index])
		copy(newReport[index:], report[index+1:])
		if isReportSafe(newReport) {
			return true
		}
	}

	return false
}

func isReportSafe(report []int) bool {
	isIncreasing := report[0] < report[1]

	for index := 1; index < len(report); index++ {
		previousNumber := report[index-1]
		number := report[index]

		if isIncreasing && previousNumber > number {
			return false
		}

		if !isIncreasing && previousNumber < number {
			return false
		}

		distance := calculateDistance(previousNumber, number)
		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true
}

func calculateDistance(leftNumber int, rightNumber int) int {
	if leftNumber > rightNumber {
		return leftNumber - rightNumber
	}
	return rightNumber - leftNumber
}
