package main

import (
	"aoc24/utils"
	"fmt"
)

// BFS
// Djikstra's
// A*

type Maze struct {
	reindeer Reindeer
	start    point
	end      point
	maze     [][]string
}

type Reindeer struct {
	position  point
	direction [2]int
}

func InitMaze(filepath string) *Maze {
	m := new(Maze)
	lines := utils.ReadFile(filepath)
	maze := make([][]string, len(lines))
	for i, line := range lines {
		maze[i] = make([]string, len(line))
		for j, char := range line {
			if char == 'S' {
				m.reindeer.position = point{i, j}
				m.reindeer.direction = [2]int{0, 1}
				m.start = point{i, j}
			} else if char == 'E' {
				m.end = point{i, j}
			}

			maze[i][j] = string(char)
		}
	}

	m.maze = maze

	return m
}

func (m *Maze) Complete() int {
	return dfsMaze(*m, m.reindeer.position, m.reindeer.direction, 0)
}

func dfsMaze(maze Maze, pos point, direction [2]int, score int) int {
	if pos == maze.end {
		return score
	}

	maze.maze[pos.x][pos.y] = "X"

	nextX := pos.x + direction[0]
	nextY := pos.y + direction[1]

	fmt.Println(nextX, nextY)

	if maze.maze[nextX][nextY] == "X" {
		return score * 100000000000000000
	}

	// var leftScore int
	if nextX < 0 || nextX >= len(maze.maze) || nextY < 0 || nextY >= len(maze.maze[0]) || maze.maze[nextX][nextY] == "#" {
		return dfsMaze(maze, pos, rotateCounterClockwise(direction), score+1000)
	}

	// var rightScore int
	// if nextX < 0 || nextX >= len(maze.maze) || nextY < 0 || nextY >= len(maze.maze[0]) || maze.maze[nextX][nextY] == "#" {
	// 	rightScore = dfsMaze(maze, pos, rotateClockwise(direction), score+1000)
	// }

	// if leftScore < rightScore {
	// 	score += leftScore
	// } else {
	// 	score += rightScore
	// }

	if maze.maze[nextX][nextY] == "." {
		return dfsMaze(maze, point{nextX, nextY}, direction, score+1)
	}

	return score
}

func rotateClockwise(direction [2]int) [2]int {
	x := direction[0]
	y := direction[1]

	if x == 0 && y == 1 {
		return [2]int{1, 0}
	}

	if x == 1 && y == 0 {
		return [2]int{0, -1}
	}

	if x == 0 && y == -1 {
		return [2]int{-1, 0}
	}

	return [2]int{0, 1}
}

func rotateCounterClockwise(direction [2]int) [2]int {
	x := direction[0]
	y := direction[1]

	if x == 0 && y == 1 {
		return [2]int{-1, 0}
	}

	if x == 1 && y == 0 {
		return [2]int{0, 1}
	}

	if x == 0 && y == -1 {
		return [2]int{1, 0}
	}

	return [2]int{0, -1}
}

func day16() {
	maze := InitMaze("input/day16.txt")
	fmt.Println(maze.Complete())
}
