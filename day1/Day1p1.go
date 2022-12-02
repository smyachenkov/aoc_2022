package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	max := 0
	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if current > max {
				max = current
			}
			current = 0
		} else {
			item, _ := strconv.Atoi(line)
			current += item
		}
	}

	fmt.Printf("Result %d\n", max)
}
