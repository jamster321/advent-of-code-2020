package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(file), "\n\n")

	part1(groups)
	part2(groups)
}

func part1(groups []string) {
	questionCount := 0

	for _, group := range groups {
		questions := map[string]bool{}

		group = strings.Replace(group, "\n", "", -1)

		for i := 0; i < len(group); i++ {
			if !questions[string(group[i])] {
				questionCount++
			}
			questions[string(group[i])] = true
		}
	}
	fmt.Println(questionCount)
}

func part2(groups []string) {
	questionCount := 0

	for _, group := range groups {
		questions := map[string]int{}
		people := strings.Split(group, "\n")

		for _, person := range people {
			for i := 0; i < len(person); i++ {
				questions[string(person[i])]++
			}
		}

		for _, count := range questions {
			if count == len(people) {
				questionCount++
			}
		}
	}
	fmt.Println(questionCount)
}
