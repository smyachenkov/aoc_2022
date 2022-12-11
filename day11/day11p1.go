package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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

type Monkey struct {
	items         []int
	operation     func(old int) int
	testDivision  int
	targetSuccess int
	targetFail    int
}

func main() {
	result := 1

	monkeys := []*Monkey{
		{
			items:         []int{61},
			operation:     func(old int) int { return old * 11 },
			testDivision:  5,
			targetSuccess: 7,
			targetFail:    4,
		}, {
			items:         []int{76, 92, 53, 93, 79, 86, 81},
			operation:     func(old int) int { return old + 4 },
			testDivision:  2,
			targetSuccess: 2,
			targetFail:    6,
		}, {
			items:         []int{91, 99},
			operation:     func(old int) int { return old * 19 },
			testDivision:  13,
			targetSuccess: 5,
			targetFail:    0,
		}, {
			items:         []int{58, 67, 66},
			operation:     func(old int) int { return old * old },
			testDivision:  7,
			targetSuccess: 6,
			targetFail:    1,
		}, {
			items:         []int{94, 54, 62, 73},
			operation:     func(old int) int { return old + 1 },
			testDivision:  19,
			targetSuccess: 3,
			targetFail:    7,
		}, {
			items:         []int{59, 95, 51, 58, 58},
			operation:     func(old int) int { return old + 3 },
			testDivision:  11,
			targetSuccess: 0,
			targetFail:    4,
		}, {
			items:         []int{87, 69, 92, 56, 91, 93, 88, 73},
			operation:     func(old int) int { return old + 8 },
			testDivision:  3,
			targetSuccess: 5,
			targetFail:    2,
		}, {
			items:         []int{71, 57, 86, 67, 96, 95},
			operation:     func(old int) int { return old + 7 },
			testDivision:  17,
			targetSuccess: 3,
			targetFail:    1,
		},
	}

	inspections := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}

	for round := 0; round < 20; round++ {
		for monkeyId, monkey := range monkeys {
			for _, item := range monkey.items {
				inspections[monkeyId]++
				worryLevel := monkey.operation(item)
				boredItem := worryLevel / 3

				var throwTarget int
				if boredItem%monkey.testDivision == 0 {
					throwTarget = monkey.targetSuccess
				} else {
					throwTarget = monkey.targetFail
				}
				fmt.Printf("Item with worry level %d is thrown to monkey  %d\n", boredItem, throwTarget)
				monkeys[throwTarget].items = append(monkeys[throwTarget].items, boredItem)
			}
			monkey.items = []int{}
		}
		fmt.Println()
	}

	for _, m := range monkeys {
		fmt.Println(m)
	}

	h := &IntHeap{}
	heap.Init(h)

	for _, v := range inspections {
		heap.Push(h, v)
	}
	for i := 0; i < 2; i++ {
		result *= heap.Pop(h).(int)
	}
	fmt.Printf("Result %d\n", result)
}
