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
	for scanner.Scan() {
		line := scanner.Text()

		// split line
		l := len(line)
		left := toMap(line[0 : l/2])
		right := toMap(line[l/2 : l])

		result = result + intersectionItem(left, right)
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

func intersectionItem(a, b map[int32]int) int32 {
	for k, _ := range a {
		if _, ok := b[k]; ok {
			return k
		}
	}
	return 0
}

//
// a b c x y z          A  B  C  X  Y  Z
// 97 98 99 120 121 122 65 66 67 88 89 90
