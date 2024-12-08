package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day8() {
	file, err := os.Open("input/day8.txt")
	if err != nil {
		fmt.Println("unable to open input file", err)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	freqToLoc := make(map[string][][]int)
	locToFreq := make(map[string]string)
	var jMax int
	puzzle := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		puzzle = append(puzzle, chars)
		for j, char := range chars {
			if char != "." {
				if _, ok := freqToLoc[char]; !ok {
					freqToLoc[char] = make([][]int, 0)
				}

				freqToLoc[char] = append(freqToLoc[char], []int{i, j})
				locToFreq[fmt.Sprintf("%d,%d", i, j)] = char

			}
			jMax = j
		}

		i++
	}

	day8part1(i, jMax, freqToLoc, locToFreq, puzzle)
}

func day8part1(iMax int, jMax int, freqToLoc map[string][][]int, locToFreq map[string]string, puzzle [][]string) {
	antinodes := make(map[string]bool, 0)
	for frequency, locations := range freqToLoc {
		for _, loc := range locations {

			// get distance to remaining locations
			distances := getDistances(loc[0], loc[1], locations)

			for i, distances := range distances {
				x := locations[i][0]
				y := locations[i][1]

				// check far side for antinode
				dx := distances[0]
				dy := distances[1]
				fmt.Println("x", x, "y", y, distances)
				fmt.Println("2x", x+2*dx)
				fmt.Println("2y", y+2*dy)
				if isValidAntionde(x, y, 2*dx, 2*dy, iMax, jMax, frequency, locToFreq) {
					antinodes[fmt.Sprintf("%d,%d", x+2*dx, y*2+dy)] = true
					puzzle[x+2*dx][y+2*dy] = "#"
				}

			}
		}
	}

	ans := 0
	for _, val := range antinodes {
		if val {
			ans++
		}
	}

	fmt.Println("ans:", ans)
	printPuzzle(puzzle)
}

func getDistances(i int, j int, locations [][]int) [][]int {
	fmt.Println("i,j,locations", i, j, locations)
	distances := make([][]int, len(locations))
	for k, loc := range locations {
		i2 := loc[0]
		j2 := loc[1]

		distances[k] = []int{i - i2, j - j2}
	}

	return distances
}

func isValidAntionde(x int, y int, dx int, dy int, xMax int, yMax int, frequency string, locToFreq map[string]string) bool {
	newX := x + dx
	newY := y + dy
	// newChar := locToFreq[fmt.Sprintf("%d,%d", newX, newY)]
	if newX >= 0 && newY >= 0 && newX < xMax && newY < yMax {
		return true
	}

	return false
}
