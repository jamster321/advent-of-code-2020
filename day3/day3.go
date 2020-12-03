package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type slope struct {
	right int
	down  int
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	mountain := strings.Split(string(file), "\n")

	if len(mountain) == 0 {
		log.Fatal("empty input")
	}

	width := len(mountain[0])

	slopes := []slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	result := 1

	for _, s := range slopes {
		trees := 0
		for row, col := 0, 0; row < len(mountain); row += s.down {
			if string(mountain[row][col]) == "#" {
				trees++
			}
			col += s.right
			if col >= width {
				col -= width
			}
		}
		result *= trees
	}

	fmt.Println(result)
}
