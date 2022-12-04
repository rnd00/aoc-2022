package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	timesContainingOthers := 0
	timesOverlapOthers := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		textLine := scanner.Text()
		holdFirst, holdSecond := 0, 0
		for i, v := range breakdownToTwoStrings(textLine) {
			numbers := breakdownToTwoNumbers(v)
			if i == 0 {
				holdFirst = numbers[0]
				holdSecond = numbers[1]
				continue
			}

			// first challenge
			if (numbers[0] >= holdFirst && numbers[1] <= holdSecond) || (numbers[0] <= holdFirst && numbers[1] >= holdSecond) {
				timesContainingOthers++
			}

			// second challenge
			if (holdFirst <= numbers[0] && holdSecond >= numbers[0]) || (numbers[0] <= holdFirst && numbers[1] >= holdFirst) {
				timesOverlapOthers++
			}
			holdFirst, holdSecond = 0, 0
		}
	}
	fmt.Println("[C1] times contains: ", timesContainingOthers)
	fmt.Println("[C2] times overlaps: ", timesOverlapOthers)
}

func breakdownToTwoStrings(data string) []string {
	split := strings.Split(data, ",")
	return split
}

func breakdownToTwoNumbers(data string) []int {
	// expecting string in "X-Y" format
	hold := []int{}
	split := strings.Split(data, "-")

	for _, v := range split {
		n, _ := strconv.Atoi(v)
		hold = append(hold, n)
	}

	return hold
}
