package main

import (
	"aoc24/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting string to int: %s", s)
	}

	return i
}

func day11() {
	lines := utils.ReadFile("input/day11.txt")

	values := strings.Split(lines[0], " ")

	cache := make(map[int]int)

	for _, v := range values {
		cache[Atoi(v)]++
	}

	const blinks int = 75

	for i := 0; i < blinks; i++ {
		cache = blink(cache)
	}

	fmt.Println("after blinking", blinks, "times, the stones are:", countStones(cache))
}

func blink(currentCache map[int]int) map[int]int {
	newCache := make(map[int]int)

	for k, v := range currentCache {
		if k == 0 {
			newCache[1] += v
		} else if len(strconv.Itoa(k))%2 == 0 {
			strVal := strconv.Itoa(k)
			halfLen := len(strVal) / 2
			firstHalf := Atoi(strVal[:halfLen])
			secondHalf := Atoi(strVal[halfLen:])

			newCache[firstHalf] += v
			newCache[secondHalf] += v
		} else {
			newCache[k*2024] += v
		}
	}

	return newCache
}

func countStones(currentCache map[int]int) int {
	sum := 0
	for _, v := range currentCache {
		sum += v
	}

	return sum
}

// This does not scale lol
func day1part1(stones []string) {
	for idx := 0; idx < 75; idx++ {
		newStones := make([]string, 0)
		for i := 0; i < len(stones); i++ {
			if stones[i] == "0" {
				newStones = append(newStones, "1")
			} else if len(stones[i])%2 == 0 {
				stone1 := stones[i][:len(stones[i])/2]
				stone2, _ := strconv.Atoi(stones[i][len(stones[i])/2:])

				newStones = append(newStones, stone1, fmt.Sprintf("%d", stone2))
			} else {
				val, _ := strconv.Atoi(stones[i])
				newStones = append(newStones, fmt.Sprintf("%d", val*2024))
			}
		}
		stones = newStones
	}

	fmt.Println("stones", len(stones))
}
