package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	ticketLength = 6

	ErrNegativeNumber = errors.New("number is negative")
	ErrTicketLength   = fmt.Errorf("number is't %d digit long", ticketLength)
)

func getTicketNumber(number string) (int, error) {
	n, err := strconv.Atoi(number)
	if err != nil {
		return 0, strconv.ErrSyntax
	}
	if len(number) != ticketLength {
		return 0, ErrTicketLength
	}

	if n < 0 {
		return 0, ErrNegativeNumber
	}
	return n, nil
}

func isEasyLucky(n int) bool {
	l, r := 0, 0
	for i := 0; i < ticketLength/2; i++ {
		l += n / int(math.Pow10(ticketLength-1-i)) % 10
		r += n / int(math.Pow10(i)) % 10
	}
	return l == r
}
func isHardLucky(n int) bool {
	var o, e, d int
	for i := 0; i < ticketLength; i++ {
		d = n / int(math.Pow10(i)) % 10
		if d%2 == 0 {
			e += d
			continue
		}
		o += d
	}
	return o == e
}

func countLuckyNumbers(scanner *bufio.Scanner) (int, int, error) {
	fmt.Print("Min: ")
	scanner.Scan()
	min, err := getTicketNumber(scanner.Text())
	if err != nil {
		return 0, 0, err
	}
	fmt.Print("Max: ")
	scanner.Scan()
	max, err := getTicketNumber(scanner.Text())
	if err != nil {
		return 0, 0, err
	}
	easyCounter, hardCounter := 0, 0
	for i := min; i <= max; i++ {
		if isEasyLucky(i) {
			easyCounter++
		}
		if isHardLucky(i) {
			hardCounter++
		}
	}
	return easyCounter, hardCounter, nil
}

func Task2(scanner *bufio.Scanner) (string, error) {
	easyCounter, hardCounter, err := countLuckyNumbers(scanner)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("\n--Result--\nEasyFormula: %d\nHardFormula: %d", easyCounter, hardCounter), nil
}
