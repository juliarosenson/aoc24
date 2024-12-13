package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func WritePuzzle(puzzle [][]string) {
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

func WriteAnswer(ans []any, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, a := range ans {
		_, err := writer.WriteString(fmt.Sprintf("%+v", a))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	writer.Flush()
	fmt.Println("File written successfully.")
}

func Contains(arr []int, s int) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}
