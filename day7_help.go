// part1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day7part2() {
	// PartOne()
	PartTwo()
}

type Row struct {
	value   int
	numbers []int
}

func PartOne() {

	data := loadData()
	ans := parseData(data)
	fmt.Println("The answer to part 1 is:", ans)
}

func loadData() []Row {
	fi, err := os.Open("input/day7.txt")
	if err != nil {
		panic("Error opening!")
	}

	defer fi.Close()
	scanner := bufio.NewScanner(fi)

	data := make([]Row, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, ": ")
		value, err := strconv.Atoi(arr[0])
		if err != nil {
			panic("erroor converting value!")
		}
		numberArr := strings.Split(arr[1], " ")
		numbers := make([]int, 0)
		for _, number := range numberArr {
			numberValue, err := strconv.Atoi(number)
			if err != nil {
				panic("error converting numbers!")
			}
			numbers = append(numbers, numberValue)
		}
		r := Row{
			value:   value,
			numbers: numbers,
		}
		data = append(data, r)
	}

	return data
}

func parseData(data []Row) int {
	ans := 0

	for _, v := range data {
		if dfs(v.value, 0, 0, v.numbers, "-") {
			ans += v.value
		}
	}

	return ans
}

// knapsack dp - at each step, we can either do 2 things: add or multiply
// so we try it all out : at each step, we take the current element, then we either add or multiply it
// with whatw e currently have
// if our initial value is empty, just add
// i.e if we have : 190: 10 19, curr = 0, curr index = 0
// 10 -> dfs(10,1,numbers,+) -> see +, so add 10 + 19
// 10 -> dfs(10,1,numbers,*) -> see *, so 10 * 19.
// when we go past the final value in numbers, just check if value == total

func dfs(value int, total int, currIndex int, numbers []int, operation string) bool {

	if currIndex >= len(numbers) {
		return value == total
	}

	currentElement := numbers[currIndex]

	currSum := total
	if operation == "-" || operation == "+" {
		currSum += currentElement
	} else {
		currSum *= currentElement
	}

	left := dfs(value, currSum, currIndex+1, numbers, "+")
	right := dfs(value, currSum, currIndex+1, numbers, "*")

	return left || right

}

// part2

func PartTwo() {
	data := loadData()
	ans := parseData2(data)
	fmt.Println("The answer to part 2 is:", ans)
}

func parseData2(data []Row) int {
	ans := 0

	for _, v := range data {
		if dfs2(v.value, 0, 0, v.numbers, "+") {
			ans += v.value
		}
	}

	return ans
}

// just another option of || - this means concat this with what we currently have
func dfs2(value int, total int, currIndex int, numbers []int, operation string) bool {

	if currIndex >= len(numbers) {
		return value == total
	}

	currentElement := numbers[currIndex]

	currSum := total
	if operation == "+" {
		currSum += currentElement
	} else if operation == "*" {
		currSum *= currentElement
	} else {
		currSum = concat(total, currentElement)
	}

	operations := []string{
		"+", "*", "||",
	}

	for _, operation_ := range operations {
		if dfs2(value, currSum, currIndex+1, numbers, operation_) {
			return true
		}
	}
	return false

}

func concat(numb1 int, numb2 int) int {

	str1 := strconv.Itoa(numb1)
	str2 := strconv.Itoa(numb2)

	result := str1 + str2

	r, err := strconv.Atoi(result)
	if err != nil {
		panic("error converting concat")
	}
	return r
}
