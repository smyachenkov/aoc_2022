package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	executeCycle int
	value        int
}
type OperationHeap []Operation

func (h OperationHeap) Len() int           { return len(h) }
func (h OperationHeap) Less(i, j int) bool { return h[i].executeCycle < h[j].executeCycle }
func (h OperationHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OperationHeap) Push(x any) {
	*h = append(*h, x.(Operation))
}

func (h *OperationHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *OperationHeap) Peek() Operation {
	return (*h)[0]
}
func (h *OperationHeap) PeekLast() Operation {
	return (*h)[len(*h)-1]
}

func (h *OperationHeap) Empty() bool {
	return len(*h) == 0
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	register := 1

	h := &OperationHeap{}
	heap.Init(h)

	executeCycle := 0

	for scanner.Scan() {
		cmd := scanner.Text()

		split := strings.Split(cmd, " ")
		var op Operation

		switch split[0] {
		case "noop":
			executeCycle += 1
			op = Operation{
				executeCycle: executeCycle,
				value:        0,
			}
		case "addx":
			v, _ := strconv.Atoi(split[1])
			executeCycle += 2
			op = Operation{
				executeCycle: executeCycle,
				value:        v,
			}
		}
		heap.Push(h, op)
	}

	cycle := 0
	row := 0
	col := 0
	for !h.Empty() {
		cycle++
		// draw
		if register >= col-1 && register <= col+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		// update
		for !h.Empty() && h.Peek().executeCycle == cycle {
			update := heap.Pop(h).(Operation).value
			register += update
		}
		col++
		if col == 40 {
			col = 0
			row++
			fmt.Printf("\n")
		}
	}
}
