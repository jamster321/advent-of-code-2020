package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input = map[int]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("non-integer input - %v", err)
		}
		input[val] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	const target = 2020

	result := findAnswer(input, target)
	if result != 0 {
		fmt.Printf("%d\n", result)
	} else {
		fmt.Println("answer 1 could not be found")
	}

	for key := range input {
		result = findAnswer(input, target-key)
		if result != 0 {
			fmt.Printf("%d\n", result*key)
			break
		}
	}
	if result == 0 {
		fmt.Println("answer 2 could not be found")
	}
}

func findAnswer(input map[int]struct{}, target int) int {
	for k1 := range input {
		k2 := target - k1
		if _, ok := input[k2]; ok {
			return k1 * k2
		}
	}
	return 0
}
