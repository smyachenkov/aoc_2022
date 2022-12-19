package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y, z int
}

const (
	typeEmpty = 0
	typeCube  = 1
	typeWater = 2
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	cubes := map[Coord]int{}

	maxDimension := []int{-1, -1, -1}          // xyz
	minDimension := []int{99999, 99999, 99999} // xyz

	// same for min
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		x := uatoi(split[0])
		y := uatoi(split[1])
		z := uatoi(split[2])
		cubes[Coord{x: x, y: y, z: z}] = typeCube
		maxDimension[0] = max(maxDimension[0], x)
		maxDimension[1] = max(maxDimension[1], y)
		maxDimension[2] = max(maxDimension[2], z)

		minDimension[0] = min(minDimension[0], x)
		minDimension[1] = min(minDimension[1], y)
		minDimension[2] = min(minDimension[2], z)
	}

	maxDimension[0] += 1
	maxDimension[1] += 1
	maxDimension[2] += 1

	minDimension[0] -= 1
	minDimension[1] -= 1
	minDimension[2] -= 1

	for x := minDimension[0]; x < maxDimension[0]; x++ {
		for y := minDimension[1]; y < maxDimension[1]; y++ {
			for z := minDimension[2]; z < maxDimension[2]; z++ {
				coord := Coord{x: x, y: y, z: z}
				if _, ok := cubes[coord]; !ok {
					cubes[coord] = typeEmpty
				}
			}
		}
	}

	// fill with water
	queue := []Coord{{x: minDimension[0], y: minDimension[1], z: minDimension[2]}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, dir := range [][]int{
			{0, 0, 1},
			{0, 0, -1},
			{0, 1, 0},
			{0, -1, 0},
			{1, 0, 0},
			{-1, 0, 0},
		} {
			neighbor := Coord{
				x: curr.x + dir[0],
				y: curr.y + dir[1],
				z: curr.z + dir[2],
			}
			if v, ok := cubes[neighbor]; ok {
				if v == typeEmpty {
					cubes[neighbor] = typeWater
					queue = append(queue, neighbor)
				}
			}
		}

	}

	contactsWater := 0
	for cube := range cubes {
		if cubes[cube] != typeCube {
			continue
		}
		exposed := 6
		for _, dir := range [][]int{
			{0, 0, 1},
			{0, 0, -1},
			{0, 1, 0},
			{0, -1, 0},
			{1, 0, 0},
			{-1, 0, 0},
		} {
			neighbor := Coord{
				x: cube.x + dir[0],
				y: cube.y + dir[1],
				z: cube.z + dir[2],
			}
			if cubetype, ok := cubes[neighbor]; ok {
				if cubetype != typeWater {
					exposed--
				}
			}
		}
		contactsWater += exposed
	}

	fmt.Println(contactsWater)
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func uatoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}
