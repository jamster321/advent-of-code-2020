package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	boardingPasses := strings.Split(string(file), "\n")

	takenSeats := map[int]bool{}
	highestSeatID := 0
	lowestSeatID := -1

	for _, bp := range boardingPasses {
		seatID := calculateSeatID(bp)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		if seatID < lowestSeatID || lowestSeatID == -1 {
			lowestSeatID = seatID
		}
		takenSeats[seatID] = true
	}

	fmt.Println(highestSeatID)

	for i := lowestSeatID; i <= highestSeatID; i++ {
		if takenSeats[i] != true {
			fmt.Println(i)
		}
	}
}

func calculateSeatID(seat string) int {
	row := parseSeatPosition(seat[:7], "F", "B")
	column := parseSeatPosition(seat[7:], "L", "R")
	return row*8 + column
}

func parseSeatPosition(rowAddr, zero, one string) int {
	rowAddr = strings.Replace(rowAddr, zero, "0", -1)
	rowAddr = strings.Replace(rowAddr, one, "1", -1)
	i, err := strconv.ParseInt(rowAddr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}
