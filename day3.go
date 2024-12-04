package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	re  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	re2 = regexp.MustCompile(`(\d{1,3})`)

	re3 = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
)

func day3() {
	day3part1()
	day3part2()
}

func day3part1() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Printf("unable to open file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		muls := re.FindAllString(line, -1)
		for _, s := range muls {
			nums := re2.FindAllString(s, -1)

			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])

			sum += n1 * n2
		}

	}
	err = scanner.Err()
	if err != nil {
		fmt.Printf("scanner encountered an err: %s", err)
	}

	fmt.Println(sum)
}

func day3part2() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Printf("unable to open file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	do := true
	for scanner.Scan() {
		line := scanner.Text()
		vals := re3.FindAllString(line, -1)
		for _, val := range vals {
			if val == "do()" {
				do = true
			}

			if val == "don't()" {
				do = false
			}
			nums := re2.FindAllString(val, -1)
			if do && len(nums) == 2 {
				n1, _ := strconv.Atoi(nums[0])
				n2, _ := strconv.Atoi(nums[1])

				sum += n1 * n2
			}
		}
	}
	fmt.Println(sum)
}
