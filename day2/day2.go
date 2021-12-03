package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/advent/utils"
)

func main() {
	lines, err := utils.GetFileLines("day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	part1Answer, err := part1Answer(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part1Answer)

	part2Answer, err := part2Answer(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part2Answer)
}

func part1Answer(lines []string) (int, error) {
	horizontal := 0
	depth := 0

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			return 0, errors.New("there should be a command and an amount for each line")
		}

		command := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		if err != nil {
			return 0, errors.New("cannot convert string to int")
		}

		switch command {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}
	return horizontal * depth, nil
}

func part2Answer(lines []string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			return 0, errors.New("there should be a command and an amount for each line")
		}

		command := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		if err != nil {
			return 0, errors.New("cannot convert string to int")
		}

		switch command {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}
	return horizontal * depth, nil
}
