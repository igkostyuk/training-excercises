package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var separator = ","

func countPositiveNumbers(data []string) (int, error) {
	res := 0
	for _, str := range data {
		num, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return 0, err
		}
		if num%2 == 0 && num > 0 {
			res += num
		}
	}
	return res, nil
}

func Task1(scanner *bufio.Scanner) (string, error) {
	fmt.Printf("Enter numbers separated by \"%s\" :", separator)
	scanner.Scan()
	data := scanner.Text()
	res, err := countPositiveNumbers(strings.Split(data, separator))
	if err != nil {
		return "", err
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strconv.Itoa(res), nil
}
