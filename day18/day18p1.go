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

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	cubes := map[Coord]int{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		cubes[Coord{
			x: uatoi(split[0]),
			y: uatoi(split[1]),
			z: uatoi(split[2]),
		}] = 6
	}

	result := 0
	for cube := range cubes {
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
			if _, ok := cubes[neighbor]; ok {
				cubes[cube]--
			}
		}
		result += cubes[cube]
	}
	fmt.Println(result)
}

func uatoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}
