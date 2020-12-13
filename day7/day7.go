package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	re := regexp.MustCompile(`(\d)+\s([\w \s]*) bag[s]?`)

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	rules := map[string]map[string]int{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineSplit := strings.Split(line, " bags contain ")
		if len(lineSplit) != 2 {
			panic("invalid line")
		}

		if lineSplit[1] == "no other bags." {
			rules[lineSplit[0]] = map[string]int{}
			continue
		}

		conditionSplit := strings.Split(lineSplit[1], ",")
		conditions := map[string]int{}

		for _, lineCondition := range conditionSplit {
			bags := re.FindAllStringSubmatch(lineCondition, -1)

			for _, bag := range bags {
				limit, err := strconv.Atoi(bag[1])
				if err != nil {
					panic(err)
				}

				conditions[bag[2]] = limit
			}
		}

		rules[lineSplit[0]] = conditions
	}

	shinyGold := "shiny gold"

	count := 0
	for bagType := range rules {
		if canContain(&rules, bagType, shinyGold) {
			count++
		}
	}
	fmt.Println(count)

	count = countBags(&rules, shinyGold) - 1 // minus 1 since we don't want to count the shiny gold bag
	fmt.Println(count)
}

func canContain(rules *map[string]map[string]int, searchBag, targetBag string) bool {
	if _, ok := (*rules)[searchBag][targetBag]; ok {
		return true
	}
	for k := range (*rules)[searchBag] {
		if canContain(rules, k, targetBag) {
			return true
		}
	}
	return false
}

func countBags(rules *map[string]map[string]int, targetBag string) int {
	sum := 1
	for k, v := range (*rules)[targetBag] {
		sum += countBags(rules, k) * v
	}
	return sum
}
