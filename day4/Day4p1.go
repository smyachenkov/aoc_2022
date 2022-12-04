package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	result := int32(0)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		a := pair[0]
		b := pair[1]

		aSplit := strings.Split(a, "-")
		aFrom, _ := strconv.Atoi(aSplit[0])
		aTo, _ := strconv.Atoi(aSplit[1])

		bSplit := strings.Split(b, "-")
		bFrom, _ := strconv.Atoi(bSplit[0])
		bTo, _ := strconv.Atoi(bSplit[1])

		if contains(aFrom, aTo, bFrom, bTo) || contains(bFrom, bTo, aFrom, aTo) {
			result++
		}
	}
	fmt.Printf("Result %d\n", result)
}

func contains(aFrom, aTo, bFrom, bTo int) bool {
	return aFrom <= bFrom && aTo >= bTo
}
