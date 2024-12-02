package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1() {
	myFile, err := os.Open("input/day1.txt")
	if err != nil {
		fmt.Printf("impossible to open file: %s", err)
	}

	scanner := bufio.NewScanner(myFile)
	left := make([]int, 0)
	right := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, "   ")
		i, _ := strconv.Atoi(vals[0])
		j, _ := strconv.Atoi(vals[1])
		left = append(left, i)
		right = append(right, j)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Printf("scanner encountered an err: %s", err)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// part1(left, right)
	part2(left, right)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(left []int, right []int) {
	distance := 0
	for i := range len(left) {
		distance += absInt(right[i] - left[i])
	}

	fmt.Println(distance)
}

func part2(left []int, right []int) {
	counts := make([]int, len(left))
	for i, val := range left {
		num := 0
		for _, rightVal := range right {
			if rightVal == val {
				num += 1
			}

			if rightVal > val {
				break
			}
		}

		counts[i] = num
	}

	ans := 0
	for i, val := range left {
		ans += val * counts[i]
	}

	fmt.Println(ans)
}
