package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var window []int32
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		m := map[int32]int{}

		for idx, char := range line {
			if len(window) == 14 {
				prevFirst := window[0]
				if m[prevFirst] == 1 {
					delete(m, prevFirst)
				} else {
					m[prevFirst] = m[prevFirst] - 1
				}
				window = window[1:]
			}
			window = append(window, char)
			if v, ok := m[char]; !ok {
				m[char] = 1
			} else {
				m[char] = v + 1
			}
			if len(m) == 14 {
				result = idx + 1
				break
			}
		}
	}

	fmt.Printf("Result %d\n", result)
}
