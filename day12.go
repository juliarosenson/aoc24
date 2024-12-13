package main

import (
	"aoc24/utils"
	"fmt"
)

func day12() {
	file := utils.ReadFile("input/day12.txt")

	puzzle := make([][]rune, 0)
	for _, line := range file {
		puzzle = append(puzzle, []rune(line))
	}

	day12Part2(puzzle)
}

func day12Part1(puzzle [][]rune) {
	fmt.Print("puzzle: ", puzzle, "\n")
	price := 0
	visited := make(map[string]bool)
	var area, perimiter int
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			// if visited[fmt.Sprintf("%d,%d", i, j)] {
			// 	continue
			// }

			area, perimiter, _ = travelRegion(puzzle, i, j, visited, puzzle[i][j])
			fmt.Println("RESULT \tarea: ", area, " perimiter: ", perimiter, "char", puzzle[i][j])
			price += area * perimiter
		}
	}

	fmt.Println(price)
}

func travelRegion(puzzle [][]rune, i int, j int, visited map[string]bool, char rune) (int, int, map[string]bool) {
	area, perimiter := 0, 0
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) || visited[fmt.Sprintf("%d,%d", i, j)] {
		return area, perimiter, visited
	}

	if puzzle[i][j] != char {
		return area, perimiter, visited
	}

	visited[fmt.Sprintf("%d,%d", i, j)] = true

	areaUp, perimiterUp, visitedUp := travelRegion(puzzle, i-1, j, visited, char)
	// fmt.Println("areaUp: ", areaUp, " perimiterUp: ", perimiterUp, "char", char)
	areaLeft, perimiterLeft, visitedLeft := travelRegion(puzzle, i, j-1, visitedUp, char)
	// fmt.Println("areaLeft: ", areaLeft, " perimiterLeft: ", perimiterLeft, "char", char)
	areaRight, perimiterRight, visitedRight := travelRegion(puzzle, i, j+1, visitedLeft, char)
	// fmt.Println("areaRight: ", areaRight, " perimiterRight: ", perimiterRight, "char", char)
	areaDown, perimiterDown, visitedDown := travelRegion(puzzle, i+1, j, visitedRight, char)
	// fmt.Println("areaDown: ", areaDown, " perimiterDown: ", perimiterDown, "char", char)

	area += 1 + areaUp + areaLeft + areaRight + areaDown
	perimiter = getPerimeter(puzzle, i, j, char) + perimiterUp + perimiterLeft + perimiterRight + perimiterDown

	return area, perimiter, visitedDown
}

func getPerimeter(puzzle [][]rune, i int, j int, char rune) int {
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) {
		return 0
	}

	perimiter := 0
	if i+1 >= len(puzzle) || puzzle[i+1][j] != char {
		perimiter++
	}
	if i-1 < 0 || puzzle[i-1][j] != char {
		perimiter++
	}

	if j+1 >= len(puzzle[i]) || puzzle[i][j+1] != char {
		perimiter++
	}

	if j-1 < 0 || puzzle[i][j-1] != char {
		perimiter++
	}

	return perimiter
}

func day12Part2(puzzle [][]rune) {
	// fmt.Print("puzzle: ", puzzle, "\n")
	price := 0
	visited := make(map[string]bool)
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			perimiterVisited := make(map[string][]string)
			area, _, pv := travelRegion2(puzzle, i, j, visited, puzzle[i][j], perimiterVisited)
			sides := calculateDistinctSides(pv)
			if area != 0 {
				// fmt.Println("perimiterVisited: ", string(puzzle[i][j]), perimiterVisited)
				fmt.Println("RESULT \tarea: ", area, "char", string(puzzle[i][j]), "sides", sides)
			}
			price += area * sides
		}
	}

	fmt.Println(price)
}

func travelRegion2(puzzle [][]rune, i int, j int, visited map[string]bool, char rune, perimeters map[string][]string) (int, map[string]bool, map[string][]string) {
	area := 0
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) || visited[fmt.Sprintf("%d,%d", i, j)] {
		return area, visited, perimeters
	}

	if puzzle[i][j] != char {
		return area, visited, perimeters
	}

	visited[fmt.Sprintf("%d,%d", i, j)] = true

	areaUp, visitedUp, perimeters := travelRegion2(puzzle, i-1, j, visited, char, perimeters)
	areaLeft, visitedLeft, perimeters := travelRegion2(puzzle, i, j-1, visitedUp, char, perimeters)
	areaRight, visitedRight, perimeters := travelRegion2(puzzle, i, j+1, visitedLeft, char, perimeters)
	areaDown, visitedDown, perimeters := travelRegion2(puzzle, i+1, j, visitedRight, char, perimeters)

	area += 1 + areaUp + areaLeft + areaRight + areaDown
	perimeters = getPerimeter2(puzzle, i, j, char, perimeters)

	return area, visitedDown, perimeters
}

func getPerimeter2(puzzle [][]rune, i int, j int, char rune, perimiters map[string][]string) map[string][]string {
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) {
		return perimiters
	}

	if i+1 >= len(puzzle) || puzzle[i+1][j] != char {
		perimiters["down"] = append(perimiters["down"], fmt.Sprintf("%d,%d", i, j))
	}
	if i-1 < 0 || puzzle[i-1][j] != char {
		perimiters["up"] = append(perimiters["down"], fmt.Sprintf("%d,%d", i, j))
	}

	if j+1 >= len(puzzle[i]) || puzzle[i][j+1] != char {
		perimiters["right"] = append(perimiters["down"], fmt.Sprintf("%d,%d", i, j))
	}

	if j-1 < 0 || puzzle[i][j-1] != char {
		perimiters["left"] = append(perimiters["down"], fmt.Sprintf("%d,%d", i, j))
	}

	return perimiters
}

func calculateDistinctSides(grid map[string][]string) int {
	sides := 0
	seen := make(map[string][]int) // Keeps track of the borders we've already counted

	for key := range grid {
		// Parse the key into i, j, and direction
		var i, j int
		var direction string
		_, err := fmt.Sscanf(key, "%d,%d,%s", &i, &j, &direction)
		if err != nil {
			fmt.Println("Error parsing key:", err)
			continue
		}

		// Check if this border is part of an already seen side
		// If the side has already been seen in the same direction for adjacent cells, skip it
		if direction == "up" {
			up, ok := seen[fmt.Sprintf("%d,%d, up", i, j)]
			if !ok {
				seen[fmt.Sprintf("%d,%d, up", i, j)] = []int{j}
				sides++
			} else {
				// if j-1 or j+1 is in the up slice, then we've already counted this side
				if !(utils.Contains(up, j-1) || utils.Contains(up, j+1)) {
					sides++
				}
				seen[fmt.Sprintf("%d,%d, up", i, j)] = append(up, j)
			}

		}

		if direction == "down" {
			down, ok := seen[fmt.Sprintf("%d,%d, down", i, j)]
			if !ok {
				seen[fmt.Sprintf("%d,%d, down", i, j)] = []int{j}
				sides++
			} else {
				// if j-1 or j+1 is in the down slice, then we've already counted this side
				if !(utils.Contains(down, j-1) || utils.Contains(down, j+1)) {
					sides++
				}
				seen[fmt.Sprintf("%d,%d, down", i, j)] = append(down, j)
			}
		}

		if direction == "left" {
			left, ok := seen[fmt.Sprintf("%d,%d, left", i, j)]
			if !ok {
				seen[fmt.Sprintf("%d,%d, left", i, j)] = []int{i}
				sides++
			} else {
				// if i-1 or i+1 is in the left slice, then we've already counted this side
				if !(utils.Contains(left, i-1) || utils.Contains(left, i+1)) {
					sides++
				}
				seen[fmt.Sprintf("%d,%d, left", i, j)] = append(left, i)
			}
		}

		if direction == "right" {
			right, ok := seen[fmt.Sprintf("%d,%d, right", i, j)]
			if !ok {
				seen[fmt.Sprintf("%d,%d, right", i, j)] = []int{i}
				sides++
			} else {
				// if i-1 or i+1 is in the right slice, then we've already counted this side
				if !(utils.Contains(right, i-1) || utils.Contains(right, i+1)) {
					sides++
				}
				seen[fmt.Sprintf("%d,%d, right", i, j)] = append(right, i)
			}
		}
	}

	return sides
}
