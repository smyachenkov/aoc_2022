package main

import (
	"bufio"
	"fmt"
	_ "fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tree struct {
	Path     string
	Size     int
	Children map[string]*Tree
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	result := 0

	// all dirs
	dirs := map[string]*Tree{}
	currentDirPath := ""

	for scanner.Scan() {
		cmd := scanner.Text()
		// ls
		if cmd == "$ ls" {
			size := 0
			parentPath := toPrevLevel(currentDirPath)
			// ls out scanner
			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "$ ") {
					// ls ended, adding prev dir
					if _, ok := dirs[currentDirPath]; !ok {
						currentNode := Tree{
							Path:     currentDirPath,
							Size:     size,
							Children: map[string]*Tree{},
						}
						if currentDirPath == "/" {
							dirs[currentDirPath] = &currentNode
						} else {
							// update parent's children
							parentNode := dirs[parentPath]
							parentNode.Children[currentDirPath] = &currentNode
							dirs[currentDirPath] = &currentNode
						}
					}
					// ls ended, forwarding to cd check
					cmd = line
					break
				}
				if strings.HasPrefix(line, "dir ") {
				}
				// file
				split := strings.Split(line, " ")
				fileSize, _ := strconv.Atoi(split[0])
				size += fileSize
			}

			if _, ok := dirs[currentDirPath]; !ok {
				currentNode := Tree{
					Path:     currentDirPath,
					Size:     size,
					Children: map[string]*Tree{},
				}
				if currentDirPath == "/" {
					dirs[currentDirPath] = &currentNode
				} else {
					// update parent's children
					parentNode := dirs[parentPath]
					parentNode.Children[currentDirPath] = &currentNode
					dirs[currentDirPath] = &currentNode
				}
			}

		}

		// cd
		if strings.HasPrefix(cmd, "$ cd") {
			line := strings.Split(cmd, " ")
			if line[2] == ".." {
				currentDirPath = toPrevLevel(currentDirPath)
			} else {
				if line[2] == "/" {
					currentDirPath = "/"
				} else if strings.HasSuffix(currentDirPath, "/") {
					currentDirPath += line[2] + "/"
				} else {
					currentDirPath += "/" + line[2] + "/"
				}
			}
		}
	}

	sizesMap := map[string]int{}
	dirs["/"].findSizes(sizesMap)

	availableSpace := 70000000 - sizesMap["/"] //23447691
	target := 30000000 - availableSpace

	var sizes []int
	for _, v := range sizesMap {
		sizes = append(sizes, v)
	}
	sort.Ints(sizes)

	idx := sort.SearchInts(sizes, target)
	result = sizes[idx]

	fmt.Printf("Available Space %d\n", availableSpace)
	fmt.Printf("Result %d\n", result)
}

func toPrevLevel(path string) string {
	split := strings.Split(path, "/")
	newDir := "/"
	var splitFiltered []string
	for _, s := range split {
		if s != "" {
			splitFiltered = append(splitFiltered, s)
		}
	}
	for i := 0; i < len(splitFiltered)-1; i++ {
		newDir += splitFiltered[i] + "/"
	}
	return newDir
}

func (t *Tree) findSizes(cache map[string]int) int {
	totalSize := t.Size
	for _, c := range t.Children {
		if v, ok := cache[c.Path]; ok {
			totalSize += v
		} else {
			totalSize += c.findSizes(cache)
		}
	}

	cache[t.Path] = totalSize
	return totalSize
}
