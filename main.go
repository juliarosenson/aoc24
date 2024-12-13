package main

import (
	"fmt"
	"sort"
)

func getNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	sideCount := 0

	for _, s := range sides {
		getCombinations := getNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

// Example usage:
func main() {
	// Example shape (E shape from the problem)
	// EEEE
	// E
	// EEEE
	// E
	// EEEE
	points := [][3]int{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
		{1, 0},
		{2, 0}, {2, 1}, {2, 2}, {2, 3},
		{3, 0},
		{4, 0}, {4, 1}, {4, 2}, {4, 3},
	}

	sides := getSideCount(points)
	fmt.Printf("Number of sides: %d\n", sides)
}
