package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := openFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	allCal := []int{}
	currentCal := 0
	maxCal := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if currentCal > maxCal {
				maxCal = currentCal
				fmt.Println(maxCal)
				allCal = append(allCal, maxCal)
			}
			currentCal = 0
		} else {
			n, _ := strconv.Atoi(scanner.Text())
			currentCal += n
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(allCal)))
	fmt.Println(allCal[0] + allCal[1] + allCal[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}
