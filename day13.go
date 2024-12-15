// https://github.com/luxedo/advent-of-code/blob/main/solutions/go/2024/13/main.go
// https://www.youtube.com/watch?v=-5J-DAsWuJc
package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Solution struct {
	A int
	B int
}

type ClawMachine struct {
	buttonA  Point
	buttonB  Point
	prize    Point
	solution *Solution
}

func day13() {
	input, _ := os.ReadFile("input/day13.txt")
	ans, _ := solve_pt2(string(input))
	fmt.Println(ans)

}

func (c ClawMachine) Parse(input string) ClawMachine {
	data := [3]Point{}
	re := regexp.MustCompile(`: X[+=](\d+), Y[+=](\d+)`)
	for i, row := range strings.Split(input, "\n") {
		matches := re.FindStringSubmatch(row)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		data[i] = Point{x, y}
	}
	c.buttonA = data[0]
	c.buttonB = data[1]
	c.prize = data[2]
	return c
}

func (c *ClawMachine) SolveDirectSubstitution() (*Solution, error) {
	x0, y0 := c.buttonA.x, c.buttonA.y
	x1, y1 := c.buttonB.x, c.buttonB.y
	cx, cy := c.prize.x, c.prize.y

	Bdividend := (cy*x0 - cx*y0)
	Bdivisor := (y1*x0 - y0*x1)
	if Bdivisor == 0 {
		return nil, errors.New("zero division for B. no solution")
	} else if Bdividend%Bdivisor != 0 {
		return nil, errors.New("non integer solution")
	}

	B := Bdividend / Bdivisor
	Adividend := (cx - B*x1)
	Adivisor := x0
	if Adivisor == 0 {
		return nil, errors.New("zero division for A. no solution")
	} else if Adividend%Adivisor != 0 {
		return nil, errors.New("non integer solution")
	}
	A := Adividend / Adivisor

	return &Solution{A, B}, nil
}

type Arcade struct {
	costA    int
	costB    int
	maxPress int
	machines []ClawMachine
}

func (a Arcade) Parse(input_data string) Arcade {
	for _, input := range strings.Split(strings.TrimSpace(input_data), "\n\n") {
		a.machines = append(a.machines, ClawMachine{}.Parse(input))
	}
	return a
}

func (a *Arcade) Solve() {
	for i, machine := range a.machines {
		if solution, err := machine.SolveDirectSubstitution(); err == nil {
			a.machines[i].solution = solution
		}
	}
}

func (a *Arcade) TotalTokens() int {
	tokens := 0
	for _, machine := range a.machines {
		if machine.solution == nil {
			continue
		} else if (a.maxPress > 0) && ((machine.solution.A > a.maxPress) || (machine.solution.B > a.maxPress)) {
			continue
		}
		tokens += machine.solution.A*a.costA + machine.solution.B*a.costB
	}
	return tokens
}

func solve_pt1(input_data string) (interface{}, error) {
	arcade := Arcade{maxPress: 100, costA: 3, costB: 1}.Parse(input_data)
	arcade.Solve()
	return arcade.TotalTokens(), nil
}

func solve_pt2(input_data string) (interface{}, error) {
	arcade := Arcade{maxPress: -1, costA: 3, costB: 1}.Parse(input_data)
	const OFFSET = 10000000000000
	for i := range arcade.machines {
		arcade.machines[i].prize.x += OFFSET
		arcade.machines[i].prize.y += OFFSET
	}
	arcade.Solve()
	return arcade.TotalTokens(), nil
}
