package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func fibonacciSeq() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func Task3(scanner *bufio.Scanner) (string, error) {
	fmt.Print("Enter fibonacci number:")
	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text())
	if err != nil || number < 0 {
		return "", fmt.Errorf("number: %w", ErrNotPositiveInteger)
	}
	nextInt := fibonacciSeq()
	var b strings.Builder
	for i := 0; i < number; i++ {
		fmt.Fprint(&b, nextInt(), " ")
	}
	fmt.Fprintln(&b)
	newInts := fibonacciSeq()
	fmt.Fprintln(&b, newInts())
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return b.String(), nil
}
