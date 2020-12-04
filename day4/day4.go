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
	passports := strings.Split(string(file), "\n\n")

	validPassports := 0

	for _, p := range passports {
		valid := true
		fields := strings.Split(strings.Replace(p, "\n", " ", -1), " ")

		fieldCounter := map[string]int{}
		requiredFields := map[string]bool{
			"byr": false,
			"iyr": false,
			"eyr": false,
			"hgt": false,
			"hcl": false,
			"ecl": false,
			"pid": false,
		}

		for _, f := range fields {
			f2 := strings.Split(f, ":")
			fieldCounter[f2[0]]++

			switch f2[0] {
			default:
				valid = false
			case "byr":
				requiredFields[f2[0]] = validateRange(f2[1], 1920, 2002)
			case "iyr":
				requiredFields[f2[0]] = validateRange(f2[1], 2010, 2020)
			case "eyr":
				requiredFields[f2[0]] = validateRange(f2[1], 2020, 2030)
			case "hgt":
				requiredFields[f2[0]] = validateHeight(f2[1])
			case "hcl":
				requiredFields[f2[0]] = validateHairColor(f2[1])
			case "ecl":
				requiredFields[f2[0]] = validateEyeColor(f2[1])
			case "pid":
				requiredFields[f2[0]] = validatePID(f2[1])
			case "cid":
				// optional, no validation
			}
		}

		// Check for exactly 1 of each field
		for _, v := range fieldCounter {
			if v != 1 {
				valid = false
				break
			}
		}

		// Check we have all required fields
		for _, v := range requiredFields {
			if !v {
				valid = false
				break
			}
		}

		if valid {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func validateHeight(input string) bool {
	if strings.HasSuffix(input, "cm") {
		return validateRange(input[:len(input)-2], 150, 193)
	}
	if strings.HasSuffix(input, "in") {
		return validateRange(input[:len(input)-2], 59, 76)
	}
	return false
}

func validateHairColor(input string) bool {
	return regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(input)
}

func validateEyeColor(input string) bool {
	validEyeColors := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	_, ok := validEyeColors[input]
	return ok
}

func validateRange(input string, lower, upper int) bool {
	i, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return i >= lower && i <= upper
}

func validatePID(input string) bool {
	re1 := regexp.MustCompile(`^0*`)
	re2 := regexp.MustCompile(`^\d{9}$`)
	return re1.MatchString(input) && re2.MatchString(input)
}
