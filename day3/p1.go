package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// filename := "test.txt"
	filename := "day3.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	counter := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	length := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for i, c := range scanner.Text() {
			if c == 49 {
				counter[i] += 1
			}
		}
		length += 1
	}

	gamma_bits := ""
	epsilon_bits := ""
	half_of_file := length / 2
	for _, v := range counter {
		if v > half_of_file {
			gamma_bits += "1"
			epsilon_bits += "0"
		} else {
			gamma_bits += "0"
			epsilon_bits += "1"
		}
	}

	// fmt.Println(gamma_bits, epsilon_bits)

	gamma, err := strconv.ParseInt(gamma_bits, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	epsilon, err := strconv.ParseInt(epsilon_bits, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(gamma * epsilon)

	// fmt.Println(counter)
	// fmt.Println(length)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
