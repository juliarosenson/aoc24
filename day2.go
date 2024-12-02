package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2() {
	myFile, err := os.Open("input/day2.txt")
	if err != nil {
		fmt.Printf("error opening file: %s", err)
	}

	scanner := bufio.NewScanner(myFile)
	reports := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		report := make([]int, 0)
		for _, v := range vals {
			i, _ := strconv.Atoi(v)
			report = append(report, i)
		}
		reports = append(reports, report)
	}

	day2part1(reports)
	day2part2(reports)

}

func day2part1(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe += 1
		}
	}

	fmt.Println(numSafe)
}

func day2part2(reports [][]int) {
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe += 1
		} else {
			for i := 0; i < len(report); i++ {
				newReport := make([]int, 0, len(report)-1)
				newReport = append(newReport, report[:i]...)
				newReport = append(newReport, report[i+1:]...)
				if isSafe(newReport) {
					numSafe += 1
					break
				}
			}
		}
	}

	fmt.Println(numSafe)
}

func isSafe(row []int) bool {
	max := len(row)
	prevDirection := sigNum(row[1] - row[0])
	for i := 0; i < max-1; i++ {
		j := i + 1

		direction := sigNum(row[j] - row[i])

		if direction == 0 {
			return false
		}

		if direction != prevDirection {
			return false
		}

		if direction > 0 && row[j]-row[i] > 3 {
			return false
		}

		if direction < 0 && row[i]-row[j] > 3 {
			return false
		}

		prevDirection = direction

	}

	return true
}

func sigNum(x int) int {
	if x < 0 {
		return -1
	}

	if x > 0 {
		return 1
	}

	return 0
}
