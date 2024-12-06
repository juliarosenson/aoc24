package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	day6()
}

func printPuzzle(puzzle [][]string) {
	file, err := os.Create("puzzle-out.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, line := range puzzle {
		_, err := writer.WriteString(strings.Join(line, "") + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	writer.Flush()
	fmt.Println("File written successfully.")
}
