package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//filename := "day2_test.txt"
	filename := "day2.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontalPosition := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		instructions := strings.Split(scanner.Text(), " ")

		distance, err := strconv.Atoi(instructions[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if instructions[0] == "forward" {
			horizontalPosition += distance
			depth += distance * aim
		} else if instructions[0] == "down" {
			aim += distance
		} else if instructions[0] == "up" {
			aim -= distance
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(horizontalPosition, depth, horizontalPosition*depth)
}
