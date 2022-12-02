package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// 1 for Rock, 2 for Paper, and 3 for Scissors
	// 0 if you lost, 3 if the round was a draw, and 6 if you won
	// X means you need to lose, Y draw, and Z win
	score := map[string]map[string]int{
		"A": {
			"X": 0 + 3,
			"Y": 3 + 1,
			"Z": 6 + 2,
		},
		"B": {
			"X": 0 + 1,
			"Y": 3 + 2,
			"Z": 6 + 3,
		},
		"C": {
			"X": 0 + 2,
			"Y": 3 + 3,
			"Z": 6 + 1,
		},
	}

	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		result += score[split[0]][split[1]]
	}

	fmt.Printf("Result %d\n", result)
}
