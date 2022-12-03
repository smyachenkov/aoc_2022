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

	result := int32(0)

	cnt := 0
	a := map[int32]int{}
	b := map[int32]int{}
	c := map[int32]int{}
	for scanner.Scan() {
		line := scanner.Text()

		// split line
		switch cnt {
		case 0:
			a = toMap(line)
			break
		case 1:
			b = toMap(line)
			break
		case 2:
			c = toMap(line)
			break
		}

		if cnt == 2 {
			ab := mapIntersection(a, b)
			abc := mapIntersection(ab, c)
			for k, _ := range abc {
				result = result + k
			}
		}

		cnt++
		if cnt == 3 {
			cnt = 0
		}
	}

	fmt.Printf("Result %d\n", result)
}

func toMap(s string) map[int32]int {
	m := map[int32]int{}
	for _, l := range s {
		cost := int32(0)
		if 'a' > l {
			// upper
			cost = l - 'A' + 27
		} else {
			// lower
			cost = l - 'a' + 1
		}
		m[cost] = 0
	}

	return m
}

func mapIntersection(a, b map[int32]int) map[int32]int {
	result := map[int32]int{}
	for k, _ := range a {
		if _, ok := b[k]; ok {
			result[k] = 0
		}
	}
	return result
}
