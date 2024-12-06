package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func day6() {
	file, err := os.Open("input/day6.txt")
	if err != nil {
		fmt.Println("unable to read input")
		return
	}

	startI, startJ := -1, -1
	i := 0
	puzzle := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		positions := strings.Split(line, "")

		if startI == -1 {
			for j, position := range positions {
				if position == "^" || position == ">" || position == "v" || position == "<" {
					startI = i
					startJ = j
					break
				}
			}
		}

		puzzle = append(puzzle, positions)
		i++
	}

	day6part2(startI, startJ, puzzle)
}

func getMove(direction string) (int, int, error) {
	if direction == "^" {
		return -1, 0, nil
	}

	if direction == ">" {
		return 0, 1, nil
	}

	if direction == "v" {
		return 1, 0, nil
	}

	if direction == "<" {
		return 0, -1, nil
	}

	return 0, 0, fmt.Errorf("invalid direction")
}

func turnRight(direction string) (string, error) {
	if direction == "^" {
		return ">", nil
	}

	if direction == ">" {
		return "v", nil
	}

	if direction == "v" {
		return "<", nil
	}

	if direction == "<" {
		return "^", nil
	}

	return "", fmt.Errorf("invalid direction")
}

func day6part1(startI int, startJ int, puzzle [][]string) {
	i, j := startI, startJ
	direction := puzzle[i][j]
	for i >= 0 && i < len(puzzle) && j >= 0 && j < len(puzzle[0]) {
		nextI, nextJ, err := getMove(direction)
		if err != nil {
			fmt.Println("error getting next move", err)
			return
		}

		fmt.Println("nextI", i+nextI)
		fmt.Println("nextJ", j+nextJ)
		if nextI+i < 0 || nextI+i >= len(puzzle) || nextJ+j < 0 || nextJ+j >= len(puzzle[i]) {
			break
		}

		if puzzle[i+nextI][j+nextJ] == "#" {
			fmt.Println("hitting wall", i+nextI, j+nextJ)
			direction, err = turnRight(direction)
			fmt.Println("next direction", direction)
			if err != nil {
				fmt.Println("error turning", err)
				return
			}
		} else {
			puzzle[i][j] = "X"
			i += nextI
			j += nextJ
		}
	}

	countX := 0
	for _, line := range puzzle {
		for _, position := range line {
			if position == "X" {
				countX++
			}
		}
	}

	countX += 1 // last move out of the grid

	fmt.Println("part1:", countX)
}

// brute force :')
func day6part2(startI int, startJ int, puzzle [][]string) {
	obstacles := 0
	obs := make([][]int, 0)
	puzzleCopy := puzzle
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == "." {
				puzzleCopy[i][j] = "#"
				if hasCycle(startI, startJ, puzzleCopy) {
					obstacles++
					fmt.Println("obstacle", i, j)
					obs = append(obs, []int{i, j})
				}
				puzzleCopy[i][j] = "."
			}
		}
	}

	for _, ob := range obs {
		i, j := ob[0], ob[1]
		puzzle[i][j] = "0"
	}

	fmt.Println("part2:", obstacles)
}

func hasCycle(startI int, startJ int, puzzle [][]string) bool {
	i, j := startI, startJ
	directionsMap := make(map[string][]string) // instead of just dir string store i,j,dir in a struct to optimize
	direction := puzzle[i][j]
	for i >= 0 && i < len(puzzle) && j >= 0 && j < len(puzzle[i]) {
		nextI, nextJ, err := getMove(direction)
		if err != nil {
			fmt.Println("error getting next move", err, direction)
			return false
		}

		if nextI+i < 0 || nextI+i >= len(puzzle) || nextJ+j < 0 || nextJ+j >= len(puzzle[i]) {
			break
		}

		if dirs, ok := directionsMap[fmt.Sprintf("%d-%d", nextI+i, nextJ+j)]; ok && slices.Contains(dirs, direction) {
			return true
		} else {
			directionsMap[fmt.Sprintf("%d-%d", nextI+i, nextJ+j)] = append(directionsMap[fmt.Sprintf("%d-%d", nextI+i, nextJ+j)], direction)
		}
		if puzzle[i+nextI][j+nextJ] == "#" {
			direction, err = turnRight(direction)
			if err != nil {
				fmt.Println("error turning", err)
				return false
			}
		} else {
			directionsMap[fmt.Sprintf("%d-%d", i, j)] = append(directionsMap[fmt.Sprintf("%d-%d", i, j)], direction)
			i += nextI
			j += nextJ
		}
	}

	return false
}
