package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row int
	col int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	walls := map[Coord]bool{} // row - col -val

	floor := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " -> ")
		for i := 0; i < len(split)-1; i++ {
			fromS := strings.Split(split[i], ",")
			from := Coord{
				row: uatoi(fromS[1]),
				col: uatoi(fromS[0]),
			}

			toS := strings.Split(split[i+1], ",")
			to := Coord{
				row: uatoi(toS[1]),
				col: uatoi(toS[0]),
			}
			floor = max(floor, to.row)

			if from.row > to.row {
				// up
				steps := from.row - to.row + 1
				for s := 0; s < steps; s++ {
					walls[Coord{row: from.row - s, col: from.col}] = true
				}
			} else if from.col < to.col {
				// right
				steps := to.col - from.col + 1
				for s := 0; s < steps; s++ {
					walls[Coord{row: from.row, col: from.col + s}] = true
				}
			} else if from.row < to.row {
				// bottom
				steps := to.row - from.row + 1
				for s := 0; s < steps; s++ {
					walls[Coord{row: from.row + s, col: from.col}] = true
				}
			} else if from.col > to.col {
				// left
				steps := from.col - to.col + 1
				for s := 0; s < steps; s++ {
					walls[Coord{row: from.row, col: from.col - s}] = true
				}
			}
		}
	}

	floor += 2
	for i := 0; i <= floor*2; i++ {
		c := Coord{
			row: floor,
			col: 500 - floor + i,
		}
		walls[c] = true
	}

	sand := Coord{
		row: 0,
		col: 500,
	}

	fixedSands := 0

	for true {
		_, blockedDown := walls[Coord{row: sand.row + 1, col: sand.col}]
		_, blockedLeftDown := walls[Coord{row: sand.row + 1, col: sand.col - 1}]
		_, blockedRightDown := walls[Coord{row: sand.row + 1, col: sand.col + 1}]

		if sand.row == 0 && sand.col == 500 && blockedDown && blockedLeftDown && blockedRightDown {
			fixedSands++
			break
		}
		if !blockedDown {
			sand.row++
		} else if !blockedLeftDown {
			sand.row++
			sand.col--
		} else if !blockedRightDown {
			sand.row++
			sand.col++
		} else {
			walls[sand] = true
			sand = Coord{
				row: 0,
				col: 500,
			}
			fixedSands++
		}
	}
	fmt.Printf("Result: %d\n", fixedSands)
}

func uatoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
