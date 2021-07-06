package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	emptySymbol = ' '

	ErrNotPositiveInteger = errors.New("not positive integer")
	ErrNotSingleSymbol    = errors.New("not single symbol")
)

type board struct {
	height  int
	width   int
	squares [][]rune
}

func (br board) String() string {
	var b strings.Builder
	for _, row := range br.squares {
		fmt.Fprintln(&b, string(row))
	}
	return b.String()
}

func newBoard(height, width int, symbol rune) *board {
	squares := make([][]rune, height)
	for i := range squares {
		row := make([]rune, width*2)
		for j := range row {
			if (i+j)%2 == 0 {
				row[j] = symbol
			} else {
				row[j] = emptySymbol
			}
		}
		squares[i] = row
	}
	return &board{height: height, width: width, squares: squares}
}

func Task0(scanner *bufio.Scanner) (string, error) {

	fmt.Print("Enter board width:")
	scanner.Scan()
	width, err := strconv.Atoi(scanner.Text())
	if err != nil || width < 0 {
		return "", fmt.Errorf("width: %w", ErrNotPositiveInteger)
	}

	fmt.Print("Enter board height:")
	scanner.Scan()
	height, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "", fmt.Errorf("height: %w", ErrNotPositiveInteger)
	}

	fmt.Print("Enter board symbol:")
	scanner.Scan()
	symbols := []rune(scanner.Text())
	if len(symbols) != 1 {
		return "", fmt.Errorf("symbol: %w", ErrNotSingleSymbol)
	}
	board := newBoard(height, width, symbols[0])
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return board.String(), nil
}
