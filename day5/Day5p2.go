package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	/*
		[H]                 [Z]         [J]
		[L]     [W] [B]     [G]         [R]
		[R]     [G] [S]     [J] [H]     [Q]
		[F]     [N] [T] [J] [P] [R]     [F]
		[B]     [C] [M] [R] [Q] [F] [G] [P]
		[C] [D] [F] [D] [D] [D] [T] [M] [G]
		[J] [C] [J] [J] [C] [L] [Z] [V] [B]
		[M] [Z] [H] [P] [N] [W] [P] [L] [C]
	*/

	state := [9]Stack{
		Stack{"M", "J", "C", "B", "F", "R", "L", "H"},
		Stack{"Z", "C", "D"},
		Stack{"H", "J", "F", "C", "N", "G", "W"},
		Stack{"P", "J", "D", "M", "T", "S", "B"},
		Stack{"N", "C", "D", "R", "J"},
		Stack{"W", "L", "D", "Q", "P", "J", "G", "Z"},
		Stack{"P", "Z", "T", "F", "R", "H"},
		Stack{"L", "V", "M", "G"},
		Stack{"C", "B", "G", "P", "F", "Q", "R", "J"},
	}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])
		from, _ := strconv.Atoi(split[3])
		to, _ := strconv.Atoi(split[5])

		if amount < 2 {
			state[to-1].Push(state[from-1].Pop())
		} else {
			stack := Stack{}
			for i := 0; i < amount; i++ {
				stack.Push(state[from-1].Pop())
			}
			for i := 0; i < amount; i++ {
				state[to-1].Push(stack.Pop())
			}
		}
	}

	for i := range state {
		fmt.Printf(state[i].Pop())
	}
}
