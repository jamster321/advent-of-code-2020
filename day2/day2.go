package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(\d*)-(\d*) (\w): (.*)`)
	lines := re.FindAllStringSubmatch(string(file), -1)

	fmt.Printf("Total passwords: %d\n", len(lines))

	policy1, policy2 := 0, 0
	for _, line := range lines {
		digit1, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		digit2, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal(err)
		}

		character, password := line[3], line[4]

		count := strings.Count(password, character)
		if count >= digit1 && count <= digit2 {
			policy1++
		}

		result1 := checkPassword(password, character, digit1)
		result2 := checkPassword(password, character, digit2)
		if (result1 || result2) && !(result1 && result2) {
			policy2++
		}
	}

	fmt.Printf("Good Passwords: policy 1: %d, policy 2: %d\n", policy1, policy2)
}

func checkPassword(password, character string, position int) bool {
	return position <= len(password) && string(password[position-1]) == character
}
