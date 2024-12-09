package main

import (
	"aoc24/utils"
	"fmt"
	"strconv"
)

func day9() {
	lines := utils.ReadFile("input/day9.txt")
	day9part1(lines[0])
}

// 6398608069280
func day9part1(line string) {
	diskMap := ""
	fileNum := 0
	isFile := true
	for _, r := range line {
		var charToPrint string
		// if i%2 == 0 {
		if isFile {
			charToPrint = fmt.Sprintf("%d", fileNum)
			fileNum++
		} else {
			charToPrint = "."
		}

		isFile = !isFile

		num, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println(err)
		}
		for range num {
			fmt.Println("num:", num, "charToPrint:", charToPrint)
			diskMap += charToPrint
		}

	}

	utils.WriteAnswer([]any{string(diskMap)}, "day9-unswapped.txt")
	swappedSlice := swapEmpty(diskMap)

	//0099811188827773336446555566
	//0099811188827773336446555566
	utils.WriteAnswer([]any{string(swappedSlice)}, "day9-swapped.txt")
	sum := sumOfDigits(swappedSlice)

	fmt.Println("sum:", sum)
}

func swapEmpty(diskMap string) []rune {
	diskMapSlice := []rune(diskMap)
	i := 0
	j := len(diskMapSlice) - 1

	for {
		for i < j && diskMapSlice[i] != '.' {
			i += 1
		}

		for i < j && diskMapSlice[j] == '.' {
			j -= 1
		}

		if i >= j {
			break
		}

		diskMapSlice[i], diskMapSlice[j] = diskMapSlice[j], diskMapSlice[i]
		i += 1
		j -= 1
	}

	return diskMapSlice
}

// func swapEmpty(diskMap string) []rune {
// 	diskMapSlice := []rune(diskMap)
// 	i := 0
// 	j := len(diskMapSlice) - 1

// 	for i < j {
// 		if diskMapSlice[i] != '.' {
// 			i++
// 			continue
// 		}

// 		if diskMapSlice[j] == '.' {
// 			j--
// 			continue
// 		}

// 		if diskMapSlice[i] == '.' && diskMapSlice[j] != '.' {
// 			diskMapSlice[i], diskMapSlice[j] = diskMapSlice[j], diskMapSlice[i]
// 		}
// 	}

// 	return diskMapSlice
// }

func sumOfDigits(diskMap []rune) int {
	sum := 0
	for i, r := range diskMap {
		num, err := strconv.Atoi(string(r))
		if err != nil {
			fmt.Println(err)
		} else {
			sum += i * num
		}
	}

	return sum
}
