package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	globalCount := 0
	firstNum, secondNum := 0, 0
	lastSum := 0
	count := 0
	for scanner.Scan() {

		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if globalCount > 2 {
			//fmt.Println(firstNum, secondNum, num)
			currSum := firstNum + secondNum + num
			if currSum > lastSum {
				count += 1
				//fmt.Println(currSum, "increasing")
			} else {
				//fmt.Println(currSum, "decreasing")
			}
			lastSum = currSum
			firstNum = secondNum
			secondNum = num
		} else if globalCount < 2 {
			if globalCount == 0 {
				firstNum = num
			} else {
				secondNum = num
			}
		} else if globalCount == 2 {
			lastSum = firstNum + secondNum + num
			//fmt.Println("t", firstNum, secondNum, num)
			firstNum = secondNum
			secondNum = num
		}

		globalCount += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(count)
}
