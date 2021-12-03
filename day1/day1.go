package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/advent/utils"
)

func main() {
	lines, err := utils.GetFileLines("day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var nums []int
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nums = append(nums, n)
	}

	part1Answer, err := part1Answer(nums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part1Answer)

	part2Answer, err := part2Answer(nums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part2Answer)
}

func part1Answer(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("there are no numbers")
	}

	currVal := nums[0]
	increases := 0

	for _, val := range nums[1:] {
		if val > currVal {
			increases++
		}
		currVal = val
	}
	return increases, nil
}

func part2Answer(nums []int) (int, error) {
	if len(nums) <= 2 {
		return 0, errors.New("there are not enough numbers to assess window size 3")
	}

	currWindowSum := nums[0] + nums[1] + nums[2]
	increases := 0

	for idx, val := range nums[3:] {
		newWindowSum := (currWindowSum - nums[idx]) + val
		if newWindowSum > currWindowSum {
			increases++
		}
		currWindowSum = newWindowSum
	}

	return increases, nil
}
