package main

import (
	"aoc24/utils"
	"fmt"
	"strconv"
	"strings"
)

func day10() {
	lines := utils.ReadFile("input/day10.txt")

	puzzle := make([][]int, 0)
	for _, line := range lines {
		heights := strings.Split(line, "")
		heightInts := make([]int, 0)
		for _, height := range heights {
			heightInt, ok := strconv.Atoi(height)
			if ok != nil {
				heightInt = -1
			}
			heightInts = append(heightInts, heightInt)
		}

		puzzle = append(puzzle, heightInts)
	}

	// fmt.Println(day10part1(puzzle))
	fmt.Println(day10part2(puzzle))
}

func day10part1(trailMap [][]int) int {
	score := 0
	for i := 0; i < len(trailMap); i++ {
		for j := 0; j < len(trailMap[i]); j++ {
			if trailMap[i][j] == 0 {
				score += countTrails(copyPuzzle(trailMap), i, j, -1)
				fmt.Printf("%+v\n", trailMap)
			}
		}
	}

	return score
}

func copyPuzzle(puzzle [][]int) [][]int {
	puzzleCopy := make([][]int, len(puzzle))
	for i := 0; i < len(puzzle); i++ {
		puzzleCopy[i] = make([]int, len(puzzle[i]))
		copy(puzzleCopy[i], puzzle[i])
	}

	return puzzleCopy
}

func countTrails(trailMap [][]int, i int, j int, prev int) int {
	score := 0
	if i < 0 || j < 0 || i >= len(trailMap) || j >= len(trailMap[i]) {
		return 0
	}

	if trailMap[i][j] == 9 && prev == 8 {
		trailMap[i][j] = -2
		return 1
	}

	if trailMap[i][j] == prev+1 {
		score += countTrails(trailMap, i-1, j, trailMap[i][j])
		score += countTrails(trailMap, i+1, j, trailMap[i][j])
		score += countTrails(trailMap, i, j-1, trailMap[i][j])
		score += countTrails(trailMap, i, j+1, trailMap[i][j])
	}

	return score
}

func day10part2(trailMap [][]int) int {
	score := 0
	for i := 0; i < len(trailMap); i++ {
		for j := 0; j < len(trailMap[i]); j++ {
			if trailMap[i][j] == 0 {
				score += countRaiting(trailMap, i, j, -1)
			}
		}
	}

	return score
}

func countRaiting(trailMap [][]int, i int, j int, prev int) int {
	score := 0
	if i < 0 || j < 0 || i >= len(trailMap) || j >= len(trailMap[i]) {
		return 0
	}

	if trailMap[i][j] == 9 && prev == 8 {
		return 1
	}

	if trailMap[i][j] == prev+1 {
		score += countRaiting(trailMap, i-1, j, trailMap[i][j])
		score += countRaiting(trailMap, i+1, j, trailMap[i][j])
		score += countRaiting(trailMap, i, j-1, trailMap[i][j])
		score += countRaiting(trailMap, i, j+1, trailMap[i][j])
	}

	return score
}
