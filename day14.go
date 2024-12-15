package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	Set[T comparable] map[T]struct{}
	Pt                struct{ x, y int }
	Guard             struct {
		p Pt
		v Pt
	}
	BathroomSecurity struct {
		height, width int
		guards        []Guard
		time          int
	}
)

func (s Set[T]) Add(g T) {
	s[g] = struct{}{}
}

func (s Set[T]) Contains(g T) bool {
	_, exists := s[g]
	return exists
}

func (g Guard) Parse(line string) Guard {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	matches := re.FindStringSubmatch(line)
	g.p.x, _ = strconv.Atoi(matches[1])
	g.p.y, _ = strconv.Atoi(matches[2])
	g.v.x, _ = strconv.Atoi(matches[3])
	g.v.y, _ = strconv.Atoi(matches[4])
	return g
}

func (b BathroomSecurity) Parse(input_data string) *BathroomSecurity {
	for _, line := range strings.Split(strings.TrimSpace(input_data), "\n") {
		b.guards = append(b.guards, Guard{}.Parse(line))
	}
	return &b
}

func (b *BathroomSecurity) String() string {
	m := make([][]rune, b.height)
	for y := range b.height {
		m[y] = make([]rune, b.width)
		for x := range b.width {
			m[y][x] = '.'
		}
	}
	for _, guard := range b.guards {
		if m[guard.p.y][guard.p.x] == '.' {
			m[guard.p.y][guard.p.x] = '1'
		} else {
			m[guard.p.y][guard.p.x]++
		}
	}
	var s string
	for _, row := range m {
		for _, r := range row {
			s += string(r)
		}
		s += "\n"
	}
	return s
}

func (b *BathroomSecurity) Move(time int) {
	b.time += time
	for i, guard := range b.guards {
		guard.p.x += time * guard.v.x
		guard.p.y += time * guard.v.y
		guard.p.x = (guard.p.x%b.width + b.width) % b.width
		guard.p.y = (guard.p.y%b.height + b.height) % b.height
		b.guards[i] = guard
	}
}

func (b *BathroomSecurity) SafetyFactor() int {
	h2, w2 := b.height/2, b.width/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, guard := range b.guards {
		if guard.p.y < h2 && guard.p.x < w2 {
			q1++
		} else if guard.p.y < h2 && guard.p.x > w2 {
			q2++
		} else if guard.p.y > h2 && guard.p.x < w2 {
			q3++
		} else if guard.p.y > h2 && guard.p.x > w2 {
			q4++
		} else {
			// Guard is in the line
		}
	}
	return q1 * q2 * q3 * q4
}

/*
	Distributing Guards Across Rows

Assume the number of guards forms a triangular shape. The number of guards in each row will typically follow a
pattern where each row has one less guard than the row beneath it. For example:

Row 1 (top row): 1 guard
Row 2: 3 guards
Row 3: 5 guards
Row 4: 7 guards
...
The number of guards increases by 2 for each subsequent row.
This pattern forms a sequence of odd numbers: 1, 3, 5, 7, ...

The total number of guards in a triangle of height ℎ is the sum
of the first ℎ odd numbers:

Total Guards=1+3+5+7+…

The sum of the first ℎ odd numbers is equal to ℎ^2

So, if we want to know how many rows (or how high the tree is)
to fit 200 guards, we can find the largest ℎ such that
ℎ is less than or equal to 200.

ℎ^2 ≤ 200

Taking the square root of 200:

ℎ ≈ 14.14

So, the tree can have 14 rows.

The total number of guards in a tree with 14 rows would be:

14^2 = 196 guards
This is very close to 200, so you would likely have 14 rows for the tree.

For a tree with 14 rows, the number of guards in the base (the 14th row) would be:

2×14−1=27 guards
This is because the bottom row of a triangular arrangement will have the greatest
number of guards, and the number follows the formula
2h−1, where h is the row number (height).
*/
func (b *BathroomSecurity) CheckEasterEgg() bool {
	pointsy := make([]Set[Pt], b.height)
	for y := range b.height {
		pointsy[y] = Set[Pt]{}
	}
	for _, guard := range b.guards {
		pointsy[guard.p.y].Add(guard.p)
	}

	for y := range b.height {
		if len(pointsy[y]) > 30 {
			seq := 0
			for py := range pointsy[y] {
				p := Pt{y: y, x: py.x + 1}
				if pointsy[y].Contains(p) {
					seq++
				}
			}
			if seq > 25 { // 14th row
				return true
			}
		}
	}
	return false
}

func solve_1(input_data string) (interface{}, error) {
	height, width := 103, 101
	bsecurity := BathroomSecurity{height: height, width: width}.Parse(input_data)
	bsecurity.Move(100)
	return bsecurity.SafetyFactor(), nil
}

func solve_2(input_data string) (interface{}, error) {
	height, width := 103, 101
	bsecurity := BathroomSecurity{height: height, width: width}.Parse(input_data)
	var egg int
	for egg = 1; egg < 11000; egg++ {
		bsecurity.Move(1)
		if bsecurity.CheckEasterEgg() {
			break
		}
	}
	return egg, nil
}

func day14() {
	input, _ := os.ReadFile("input/day14.txt")
	ans, _ := solve_1(string(input))
	fmt.Println(ans)
	ans, _ = solve_2(string(input))
	fmt.Println(ans)
}
