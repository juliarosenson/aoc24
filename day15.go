package main

import (
	"aoc24/utils"
	"fmt"
	"slices"
	"strings"
)

type Warehouse struct {
	grid  [][]string
	moves []string
	robot Point
}

func Parse(filename string) *Warehouse {
	w := &Warehouse{}
	file := utils.ReadFile(filename)

	afterBreak := false
	puzzle := make([][]string, 0)
	moves := make([]string, 0)
	for i, line := range file {
		if !afterBreak {
			if line == "" {
				afterBreak = true
			} else {
				row := strings.Split(line, "")
				for j, val := range row {
					if val == "@" {
						w.robot = Point{x: i, y: j}
					}
				}
				puzzle = append(puzzle, row)
			}
		} else {
			row := strings.Split(line, "")
			moves = append(moves, row...)
		}
	}

	w.grid = puzzle
	w.moves = moves

	return w
}

func ParseWide(filename string) *Warehouse {
	w := &Warehouse{}
	file := utils.ReadFile(filename)

	afterBreak := false
	puzzle := make([][]string, 0)
	moves := make([]string, 0)
	for i, line := range file {
		if !afterBreak {
			if line == "" {
				afterBreak = true
			} else {
				row := strings.Split(line, "")
				for j, val := range row {
					if val == "@" {
						w.robot = Point{x: i, y: j}
					}
				}
				wideRow := make([]string, 0)
				for _, char := range row {
					if char == "#" || char == "." {
						wideRow = append(wideRow, char, char)
					} else if char == "@" {
						wideRow = append(wideRow, char, ".")
					} else {
						wideRow = append(wideRow, "[", "]")
					}
				}
				puzzle = append(puzzle, wideRow)
			}
		} else {
			row := strings.Split(line, "")
			moves = append(moves, row...)
		}
	}

	w.grid = puzzle
	w.moves = moves

	return w
}

func (w *Warehouse) Predict() {
	for _, move := range w.moves {

		dx, dy := getDxDy(move)

		if w.grid[w.robot.x+dx][w.robot.y+dy] == "." {
			w.grid[w.robot.x][w.robot.y] = "."
			w.robot.x = w.robot.x + dx
			w.robot.y = w.robot.y + dy
			w.grid[w.robot.x][w.robot.y] = "@"
		} else if w.grid[w.robot.x+dx][w.robot.y+dy] == "O" {
			fmt.Println("next move is into box")
			nextX := w.robot.x + dx
			nextY := w.robot.y + dy
			canShift := false
			for nextX > 0 && nextX < len(w.grid) && nextY > 0 && nextY < len(w.grid[0]) {
				nextX += dx
				nextY += dy

				fmt.Println("shifting")
				if w.grid[nextX][nextY] == "#" {
					fmt.Println("hit wall")
					break
				}
				if w.grid[nextX][nextY] == "." {
					fmt.Println("found opening")
					canShift = true
					break
				}
			}

			if canShift {
				w.grid[w.robot.x][w.robot.y] = "."
				w.robot.x = w.robot.x + dx
				w.robot.y = w.robot.y + dy
				w.grid[w.robot.x][w.robot.y] = "@"

				w.grid[nextX][nextY] = "O"
			}
		}

		fmt.Println("moved", move, dx, dy)
		for _, row := range w.grid {
			for _, col := range row {
				fmt.Print(col)
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")

	}
}

func (w *Warehouse) PredictWide() {
	for _, move := range w.moves {

		dx, dy := getDxDy(move)

		if w.grid[w.robot.x+dx][w.robot.y+dy] == "." {
			w.grid[w.robot.x][w.robot.y] = "."
			w.robot.x = w.robot.x + dx
			w.robot.y = w.robot.y + dy
			w.grid[w.robot.x][w.robot.y] = "@"
		} else if w.grid[w.robot.x+dx][w.robot.y+dy] == "[" || w.grid[w.robot.x+dx][w.robot.y+dy] == "]" {

			// BFS to move boxes see day15.py

		}

		fmt.Println("moved", move, dx, dy)
		w.PrintGrid()

	}
}

func (w *Warehouse) PrintGrid() {
	for _, row := range w.grid {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Print("\n")
	}
}

func (w *Warehouse) GetBoxCoordinatesSum() int {
	sum := 0
	for i, row := range w.grid {
		for j := range row {
			if w.grid[i][j] == "O" {
				sum += 100*i + j
			}
		}
	}

	return sum
}

func getDxDy(char string) (int, int) {
	if char == "<" {
		return 0, -1
	}

	if char == ">" {
		return 0, 1
	}

	if char == "^" {
		return -1, 0
	}

	if char == "v" {
		return 1, 0
	}

	return 0, 0
}

func day15() {
	// w := Parse("input/day15.txt")
	// w.Predict()
	// fmt.Println(w.GetBoxCoordinatesSum())
	w := ParseWide("input/day15.txt")
	w.PrintGrid()
	fmt.Println()
	w.PredictWide()
	// fmt.Println(w.GetBoxCoordinatesSum())
}

func removeDuplicates(inp [][2]int) [][2]int {
	out := [][2]int{}
	for _, v := range inp {
		if !slices.Contains(out, v) {
			out = append(out, v)
		}
	}
	return out
}

func findAllConnected(curLoc [2]int, m [][]string, dir string) [][2]int {
	coords := [][2]int{}

	switch dir {
	case "^":
		coords = append(coords, [2]int{curLoc[0], curLoc[1] - 1})
		if m[curLoc[1]-1][curLoc[0]] == "[" {
			coords = append(coords, [2]int{curLoc[0] + 1, curLoc[1] - 1})
		} else {
			coords = append(coords, [2]int{curLoc[0] - 1, curLoc[1] - 1})
		}
		for i := curLoc[1] - 1; i >= 0; i-- {
			add_to_coords := [][2]int{}
			for j := 0; j < len(m[0]); j++ {
				for ind := range coords {
					if j == coords[ind][0] && coords[ind][1] == i+1 && (m[i][j] == "[" || m[i][j] == "]" || m[i][j] == "#") {
						add_to_coords = append(add_to_coords, [2]int{j, i})
					}
				}
			}
			others_to_add := [][2]int{}
			for j := 0; j < len(add_to_coords); j++ {
				if m[add_to_coords[j][1]][add_to_coords[j][0]] == "[" {
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0] + 1, add_to_coords[j][1]})
				} else if m[add_to_coords[j][1]][add_to_coords[j][0]] == "]" {
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0] - 1, add_to_coords[j][1]})
				}
			}
			add_to_coords = removeDuplicates(append(add_to_coords, others_to_add...))
			coords = append(coords, add_to_coords...)
		}
	case "v":
		coords = append(coords, [2]int{curLoc[0], curLoc[1] + 1})
		if m[curLoc[1]+1][curLoc[0]] == "[" {
			coords = append(coords, [2]int{curLoc[0] + 1, curLoc[1] + 1})
		} else {
			coords = append(coords, [2]int{curLoc[0] - 1, curLoc[1] + 1})
		}
		for i := curLoc[1] + 1; i < len(m); i++ {
			add_to_coords := [][2]int{}
			for j := 0; j < len(m[0]); j++ {
				for ind := range coords {
					if j == coords[ind][0] && coords[ind][1] == i-1 && (m[i][j] == "[" || m[i][j] == "]" || m[i][j] == "#") {
						add_to_coords = append(add_to_coords, [2]int{j, i})
					}
				}
			}
			others_to_add := [][2]int{}
			for j := 0; j < len(add_to_coords); j++ {
				if m[add_to_coords[j][1]][add_to_coords[j][0]] == "[" {
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0] + 1, add_to_coords[j][1]})
				} else if m[add_to_coords[j][1]][add_to_coords[j][0]] == "]" {
					others_to_add = append(others_to_add, [2]int{add_to_coords[j][0] - 1, add_to_coords[j][1]})
				}
			}
			add_to_coords = removeDuplicates(append(add_to_coords, others_to_add...))
			coords = append(coords, add_to_coords...)
		}
	default:
		// dont do anything
	}
	for _, c := range coords {
		if m[c[1]][c[0]] == "#" {
			coords = [][2]int{}
			break
		}
	}
	return coords
}
