package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day7() {
	file, err := os.Open("input/day7.txt")
	if err != nil {
		fmt.Println("unable to open file", err)
		return
	}

	scanner := bufio.NewScanner(file)
	puzzle := make([][]float64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, " ")
		row := make([]float64, 0)
		for i := range len(chars) {
			var charVal string
			if i == 0 {
				val := strings.Split(chars[i], ":")
				charVal = val[0]
			} else {
				charVal = chars[i]
			}
			floatVal, err := strconv.ParseFloat(charVal, 64)
			if err != nil {
				fmt.Println("err converting char to int", charVal, err)
				return
			}
			row = append(row, floatVal)
		}
		puzzle = append(puzzle, row)
	}

	day7part1(puzzle)
}

func day7part1(puzzle [][]float64) {
	sum := 0.0
	for _, row := range puzzle {
		if canEquate(row[0], row[1:]) {
			fmt.Println("good row:", row)
			sum += row[0]
		} else {

		}
	}

	fmt.Println("part1:", int(sum))
}

func canEquate(target float64, values []float64) bool {
	if len(values) == 0 {
		return target == 0
	}

	add := canEquate(target-values[len(values)-1], values[:len(values)-1])
	multiply := canEquate(target/values[len(values)-1], values[:len(values)-1])

	return add || multiply
}
