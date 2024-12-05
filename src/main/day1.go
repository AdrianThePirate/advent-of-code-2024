package main

import (
	"adventutils"
	"bufio"
	"fmt"
	"os"
	"sort"
)

// var left = []int{3, 4, 2, 1, 3, 3}
// var right = []int{4, 3, 5, 3, 9, 3}
var left, right []int

func main() {
	populate_list()
	part1()
	part2()
}

func part1() {
	sort.Ints(left)
	sort.Ints(right)

	var result int
	for index, value := range left {
		result += adventutils.Absolute(value - right[index])
	}

	fmt.Println("Result part 1:", result)
}

func part2() {
	right_map := repeat_map(right)

	var result int
	for _, value := range left {
		result += (value * right_map[value])
	}

	fmt.Println("Result part 2:", result)
}

func populate_list() {
	file, err := os.Open("tasks/day1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var num1, num2 int

		_, err := fmt.Sscanf(scanner.Text(), "%d   %d", &num1, &num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		left = append(left, num1)
		right = append(right, num2)
	}
}

func repeat_map(n []int) map[int]int {
	mapped := map[int]int{}
	for _, value := range n {
		_, exists := mapped[value]
		if exists {
			mapped[value] += 1
		} else {
			mapped[value] = 1
		}
	}

	return mapped
}