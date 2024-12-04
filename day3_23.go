package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	r, _ = regexp.Compile("[0-9.]")
)

func day3_23() {
	myFile, err := os.Open("input/day3.txt")
	if err != nil {
		// fmt.Printf("impossible to open file: %s", err)
	}

	scanner := bufio.NewScanner(myFile)
	puzzle := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, strings.Split(line, ""))
	}
	err = scanner.Err()
	if err != nil {
		// fmt.Printf("scanner encountered an err: %s", err)
	}

	day3_23part1(puzzle)
}

func day3_23part1(puzzle [][]string) {
	sum := 0
	for i, row := range puzzle {
		for j, col := range row {
			if !r.MatchString(col) {
				fmt.Println("not a number", i*10+j, puzzle[i][j])
				sum += getAdjacentNumsSum(i, j, len(puzzle), len(col), &puzzle)
			}
		}
	}

	fmt.Println(sum)

	file, err := os.Create("puzzle-out.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range puzzle {
		_, err := writer.WriteString(strings.Join(line, "") + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	writer.Flush()
	fmt.Println("File written successfully.")

}

func getAdjacentNumsSum(i int, j int, rowMax int, colMax int, puzzle *[][]string) (sum int) {
	sum = 0
	fmt.Println(("checking adjacent nums for"), i, j, (*puzzle)[i][j])
	sum += checkTopLeftCorner(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after top left", sum)
	sum += checkTop(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after top", sum)
	sum += checkTopRightCorner(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after top right", sum)
	sum += checkLeft(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after left", sum)
	sum += checkRight(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after right", sum)
	sum += checkBottomLeftCorner(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after bottom left", sum)
	sum += checkBottom(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after bottom", sum)
	sum += checkBottomRightCorner(i, j, rowMax, colMax, puzzle)
	fmt.Println("sum after bottom right", sum)

	return sum
}

func checkTopLeftCorner(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == 0 || j == 0 {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i-1][j-1]); ok == nil {
		return getNumber(i-1, j-1, puzzle)
	}

	return 0
}

func checkTop(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == 0 {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i-1][j]); ok == nil {
		return getNumber(i-1, j, puzzle)
	}

	return 0
}

func checkTopRightCorner(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == 0 || j == colMax {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i-1][j+1]); ok == nil {
		return getNumber(i-1, j+1, puzzle)
	}

	return 0
}

func checkLeft(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if j == 0 {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i][j-1]); ok == nil {
		return getNumber(i, j-1, puzzle)
	}

	return 0
}

func checkRight(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if j == colMax {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i][j+1]); ok == nil {
		return getNumber(i, j+1, puzzle)
	}

	return 0
}

func checkBottomLeftCorner(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == rowMax || j == 0 {
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i+1][j-1]); ok == nil {
		return getNumber(i+1, j-1, puzzle)
	}

	return 0
}

func checkBottom(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == rowMax {
		fmt.Println("bottom")
		return 0
	}

	if _, ok := strconv.Atoi((*puzzle)[i+1][j]); ok == nil {
		return getNumber(i+1, j, puzzle)
	}

	return 0
}
func checkBottomRightCorner(i int, j int, rowMax int, colMax int, puzzle *[][]string) int {
	if i == rowMax || j == colMax {
		return 0
	}

	// fmt.Println("checking bottom right corner", i, j, rowMax, colMax)
	if _, ok := strconv.Atoi((*puzzle)[i+1][j+1]); ok == nil {
		return getNumber(i+1, j+1, puzzle)
	}

	return 0
}

func getNumber(i int, j int, puzzle *[][]string) int {
	numStr := ""
	// travers left til we hit non number
	for j >= 0 {
		if _, ok := strconv.Atoi((*puzzle)[i][j]); ok == nil {
			j--
		} else {
			break
		}
	}
	j++
	// traverse right til we hit non number
	for j < len((*puzzle)[i]) {
		if _, ok := strconv.Atoi((*puzzle)[i][j]); ok == nil {
			numStr += (*puzzle)[i][j]
			(*puzzle)[i][j] = "."
			j++
		} else {
			break
		}
	}

	num, _ := strconv.Atoi(numStr)

	fmt.Println("found num", num)
	return num
}

//553073 too low
