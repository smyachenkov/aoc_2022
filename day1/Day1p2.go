package main

import (
	"bufio"
	"container/heap"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	h := &IntHeap{}
	heap.Init(h)

	current := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if h.Len() < 3 {
				heap.Push(h, current)
			} else {
				peek := (heap.Pop(h)).(int)
				if current > peek {
					heap.Push(h, current)
				} else {
					heap.Push(h, peek)
				}
			}
			current = 0
		} else {
			item, _ := strconv.Atoi(line)
			current += item
		}
	}

	result := 0
	for h.Len() > 0 {
		top := (heap.Pop(h)).(int)
		result = result + top
	}

	fmt.Printf("Result %d\n", result)
}
