package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var (
	ErrNotDigit = errors.New("not a digit")
	ErrLuhnSum  = errors.New("invalid Luhn sum")
)

func validate(number string) error {
	parity, luhnSum := len(number)%2, 0
	for i, r := range number {
		if !unicode.IsDigit(r) {
			return fmt.Errorf("%#U: %w", r, ErrNotDigit)
		}
		digit := int(r - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		luhnSum += digit
	}
	if luhnSum%10 != 0 {
		return ErrLuhnSum
	}
	return nil
}

func Task2(scanner *bufio.Scanner) (string, error) {
	fmt.Print("Enter card number:")
	scanner.Scan()
	data := scanner.Text()
	number := strings.ReplaceAll(data, " ", "")
	if err := validate(number); err != nil {
		return "", err
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return number[:len(number)-4], nil
}
