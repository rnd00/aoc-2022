package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// test()
	// challengeFirst()
	challengeSecond()
}

func test() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Println(string(i), " -> ", calculateScorePerRune(i))
	}

	for i := 'A'; i <= 'Z'; i++ {
		fmt.Println(string(i), " -> ", calculateScorePerRune(i))
	}
}

func challengeFirst() {
	f, err := openFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalPrios := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(totalPrios)

		totalPrios += calculateScorePerRune(findSameRune(scanner.Text()))
	}

	fmt.Println("total: ", totalPrios)
}

func challengeSecond() {
	f, err := openFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalPrios := 0

	v := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v = append(v, scanner.Text())
		if len(v) == 3 {
			r, err := findSameRuneThreeBag(v)
			if err != nil {
				log.Fatal(err)
			}
			totalPrios += calculateScorePerRune(r)

			fmt.Println(totalPrios)
			v = []string{}
		}
	}

	fmt.Println("total: ", totalPrios)
}

func openFile(filename string) (*os.File, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func findSameRune(s string) rune {
	checkedCharacters := ""
	// split string into two equal parts (compartments)
	first := s[0 : len(s)/2]
	later := s[len(s)/2:]
	// run strings.IndexRune()
	for _, v := range first {
		if strings.ContainsRune(checkedCharacters, v) {
			continue
		}
		checkedCharacters += string(v)
		if strings.IndexRune(later, v) > -1 {
			return v
		}
	}
	return '-'
}

func findSameRuneThreeBag(s []string) (rune, error) {
	if len(s) != 3 {
		return ' ', fmt.Errorf("length is not three")
	}

	checkedCharacters := ""

	for _, v := range s[0] {
		if strings.ContainsRune(checkedCharacters, v) {
			continue
		}
		checkedCharacters += string(v)
		if strings.IndexRune(s[1], v) > -1 {
			if strings.IndexRune(s[2], v) > -1 {
				return v, nil
			}
		}
	}

	return ' ', fmt.Errorf("did not find any")
}

func calculateScorePerRune(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r) - int('A') + 27
	}
	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}
	return -1
}
