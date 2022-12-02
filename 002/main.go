package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var m = map[rune]int{
	'A': -1,
	'B': -2,
	'C': -3,
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var s = map[int]int{
	-2: 6,
	-1: 0,
	0:  3,
	1:  6,
	2:  0,
}

var c = map[rune]int{
	'X': 0,
	'Y': 3,
	'Z': 6,
}

var z = map[rune]rune{
	'A': 'Y',
	'B': 'Z',
	'C': 'X',
}

var y = map[rune]rune{
	'A': 'X',
	'B': 'Y',
	'C': 'Z',
}

var x = map[rune]rune{
	'A': 'Z',
	'B': 'X',
	'C': 'Y',
}

func main() {
	f, err := openFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n := check(scanner.Text())
		fmt.Printf("%d\n", n)
		total += n
	}

	fmt.Println("total: ", total)
}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func calculate(data string) (result int) {
	total := 0
	for i, v := range data {
		if v == ' ' {
			continue
		}
		total += m[v]
		fmt.Printf("%d ", m[v])
		if i == 2 {
			result = s[total] + m[v]
		}
	}
	return
}

func check(data string) (result int) {
	keep := []rune{}
	for _, v := range data {
		if v == ' ' {
			continue
		}
		keep = append(keep, v)
	}

	// check win condition/score
	result += c[keep[1]]
	// check our hands' score
	switch keep[1] {
	case 'X':
		result += m[x[keep[0]]]
	case 'Y':
		result += m[y[keep[0]]]
	case 'Z':
		result += m[z[keep[0]]]
	}

	return
}
