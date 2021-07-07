package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var numberLength = 10

func getPalindrom(n []rune, i, j int) []rune {
	for i >= 0 && j < len(n) && n[i] == n[j] {
		i, j = i-1, j+1
	}
	return n[i+1 : j]
}
func Task1(scanner *bufio.Scanner) (string, error) {
	fmt.Printf("Enter a number that contains more than %d digits: ", numberLength)
	scanner.Scan()
	number := []rune(scanner.Text())
	if len(number) < numberLength {
		return "", fmt.Errorf("number less then %d digit long: %w", numberLength, strconv.ErrRange)
	}
	var palindroms strings.Builder
	for i := 0; i < len(number); i++ {
		if palindrom := getPalindrom(number, i, i); len(palindrom) > 2 {
			fmt.Fprint(&palindroms, string(palindrom), " ")
		}
		if palindrom := getPalindrom(number, i, i+1); len(palindrom) > 2 {
			fmt.Fprint(&palindroms, string(palindrom), " ")
		}
	}
	return fmt.Sprintf("palindroms: %s", palindroms.String()), nil
}
