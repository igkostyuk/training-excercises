package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func runTask(scanner *bufio.Scanner) (string, error) {
	fmt.Print("Enter task number:")
	scanner.Scan()
	switch scanner.Text() {
	case "0":
		return Task0(scanner)
	case "1":
		return Task1(scanner)
	case "2":
		return Task2(scanner)
	case "3":
		return Task3(scanner)
	default:
		return "", errors.New("task number must be in range 0-3")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result, err := runTask(scanner)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
