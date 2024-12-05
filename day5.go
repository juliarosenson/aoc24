package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day5() {
	file, err := os.Open("input/day5.txt")
	if err != nil {
		fmt.Printf("impossible to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	order := make(map[string][]string)
	isOrder := true
	puzzle := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isOrder = false
			continue
		}
		if isOrder {
			words := strings.Split(line, "|")
			if _, ok := order[words[1]]; !ok {
				order[words[1]] = make([]string, 0)
			}
			order[words[1]] = append(order[words[1]], words[0])
		} else {
			puzzle = append(puzzle, strings.Split(line, ","))
		}
	}
	day5part2(order, puzzle)
}

func day5part1(order map[string][]string, puzzle [][]string) {
	correctIndices := make([]int, 0)
	for i, row := range puzzle {
		correct := true
		seen := make(map[string]bool)
		for _, col := range row {
			seen[col] = true
			if val, ok := order[col]; ok {
				for _, precursor := range val {
					if !slices.Contains(row, precursor) {
						continue
					}

					if _, ok := seen[precursor]; !ok {
						correct = false
						break
					}
				}
			}

		}
		if correct {
			correctIndices = append(correctIndices, i)
		}
	}

	sum := 0
	for _, i := range correctIndices {
		middleIdx := len(puzzle[i]) / 2
		num, _ := strconv.Atoi(puzzle[i][middleIdx])
		sum += num
	}

	fmt.Println(sum)
}

func day5part2(order map[string][]string, puzzle [][]string) {
	incorrectIndices := make(map[int]bool)
	correctedRow := make(map[int][]string, 0)

	for i, row := range puzzle {
		seen := make(map[string]bool)
		for _, col := range row {
			seen[col] = true
			if val, ok := order[col]; ok {
				for _, precursor := range val {
					if !slices.Contains(row, precursor) {
						continue
					}

					if _, ok := seen[precursor]; !ok {
						incorrectIndices[i] = true
						row = redorder(row, order)
					}
				}
			}

		}
		correctedRow[i] = row
	}

	sum := 0

	for i, _ := range incorrectIndices {
		row := correctedRow[i]
		middleIdx := len(row) / 2
		num, _ := strconv.Atoi(row[middleIdx])
		sum += num
	}

	print(sum)

}

func redorder(row []string, order map[string][]string) []string {
	seen := make(map[string]bool)

	for _, col := range row {
		seen[col] = true
		if val, ok := order[col]; ok {
			for _, precursor := range val {
				if !slices.Contains(row, precursor) {
					continue
				}

				if _, ok := seen[precursor]; !ok {
					row = moveBehind(row, col, precursor)
				}
			}
		}

	}

	return row
}

func moveBehind(row []string, col string, precursor string) []string {
	newRow := make([]string, 0)
	for _, c := range row {
		if c == col {
			continue
		} else if c == precursor {
			newRow = append(newRow, c)
			newRow = append(newRow, col)
		} else {
			newRow = append(newRow, c)
		}
	}

	return newRow
}
