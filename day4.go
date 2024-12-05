package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day4() {
	input, err := os.Open("input/day4.txt")
	if err != nil {
		fmt.Printf("Unable to open file")
		return
	}

	scanner := bufio.NewScanner(input)
	puzzle := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		puzzle = append(puzzle, line)
	}

	// day4part1(puzzle)
	day4part2(puzzle)

}

func day4part1(puzzle [][]string) {
	sum := 0
	for i := range len(puzzle) {
		for j := range len(puzzle[i]) {
			if puzzle[i][j] == "X" {
				temp := 0
				temp += topLeft(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += top(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += topRight(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += left(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += right(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += bottomLeft(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += bottom(i, j, puzzle, []string{"X", "M", "A", "S"})
				temp += bottomRight(i, j, puzzle, []string{"X", "M", "A", "S"})
				puzzle[i][j] = fmt.Sprintf("%d", temp)
				sum += temp
			}
		}
	}

	fmt.Println("sum", sum)

	for i := range len(puzzle) {
		fmt.Printf("%+v\n", puzzle[i])
	}
}

// Not used but would've if word search could've gone any direction like NYT strands
func findWord(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	topLeft := findWord(i-1, j-1, puzzle, target[1:])
	top := findWord(i-1, j, puzzle, target[1:])
	topRight := findWord(i-1, j+1, puzzle, target[1:])
	left := findWord(i, j-1, puzzle, target[1:])
	right := findWord(i, j+1, puzzle, target[1:])
	bottomLeft := findWord(i+1, j-1, puzzle, target[1:])
	bottom := findWord(i+1, j, puzzle, target[1:])
	bottomRight := findWord(i+1, j+1, puzzle, target[1:])

	return topLeft + top + topRight + left + right + bottomLeft + bottom + bottomRight
}

func topLeft(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return topLeft(i-1, j-1, puzzle, target[1:])
}

func topRight(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return topRight(i-1, j+1, puzzle, target[1:])
}

func top(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return top(i-1, j, puzzle, target[1:])
}

func left(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return left(i, j-1, puzzle, target[1:])
}

func right(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return right(i, j+1, puzzle, target[1:])
}

func bottomLeft(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return bottomLeft(i+1, j-1, puzzle, target[1:])
}

func bottom(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return bottom(i+1, j, puzzle, target[1:])
}

func bottomRight(i int, j int, puzzle [][]string, target []string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if target[0] != puzzle[i][j] {
		return 0
	}

	if len(target) == 1 {
		return 1
	}

	return bottomRight(i+1, j+1, puzzle, target[1:])
}

func day4part2(puzzle [][]string) {
	sum := 0
	for i := range len(puzzle) {
		for j := range len(puzzle[i]) {
			if puzzle[i][j] == "A" {
				sum += countXMAS(i, j, puzzle)
			}
		}
	}

	fmt.Println("sum", sum)
}

// A surrounded by 2 Ms and 2 Ss and Ms not diagonal from each other
func countXMAS(i int, j int, puzzle [][]string) int {
	if i == 0 || j == 0 || i == len(puzzle)-1 || j == len(puzzle[i])-1 {
		return 0
	}

	cornerM := countTarget(i-1, j-1, puzzle, "M") + countTarget(i-1, j+1, puzzle, "M") + countTarget(i+1, j-1, puzzle, "M") + countTarget(i+1, j+1, puzzle, "M")
	cornerS := countTarget(i-1, j-1, puzzle, "S") + countTarget(i-1, j+1, puzzle, "S") + countTarget(i+1, j-1, puzzle, "S") + countTarget(i+1, j+1, puzzle, "S")

	if cornerM == 2 && cornerS == 2 && puzzle[i-1][j-1] != puzzle[i+1][j+1] {
		return 1
	}

	return 0
}

func countTarget(i int, j int, puzzle [][]string, target string) int {
	if i < 0 || i >= len(puzzle) {
		return 0
	}

	if j < 0 || j >= len(puzzle[i]) {
		return 0
	}

	if puzzle[i][j] == target {
		return 1
	}

	return 0
}
