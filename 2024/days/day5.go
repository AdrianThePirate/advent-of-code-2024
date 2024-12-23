package main

import (
	"advent/adventutils"
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rules map[int][]int
var rulesReversed map[int][]int
var pages [][]int
var incorrect [][]int

func main() {
	rules = make(map[int][]int)
	rulesReversed = make(map[int][]int)

	importData()
	part1()
	part2()
}

func part1() {
	var result int

	for i := range pages{
		var valid = true
		var forbidden []int
		for _, val := range pages[i] {
			if slices.Contains(forbidden, val){
				valid = false
				incorrect = append(incorrect, pages[i])
				break
			}
			if rules[val] != nil {
				forbidden = append(forbidden, rules[val]...)
			}
		}
		if valid { result += pages[i][len(pages[i])/2] }
	}

	fmt.Println("Result:", result)
}

func part2(){
	var result int

	for i := range incorrect {
		var forbidden []int
		for j, val := range incorrect[i] {
			if slices.Contains(forbidden, val){
				for z, lav := range incorrect[i]{
					if slices.Contains(rulesReversed[val], lav){
						incorrect[i] = adventutils.MoveIndex(incorrect[i], j, z)
						break
					}
				}
			}
			if rules[val] != nil {
				forbidden = append(forbidden, rules[val]...)
			}
		}
		result += incorrect[i][len(incorrect[i])/2]
	}

	fmt.Println("Result:", result)
}

func importData(){
	file, err := os.Open("2024/tasks/day5_sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	
	var rulemode = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			rulemode = false
			continue
		}
		if rulemode{ 
			var num1, num2 int
			fmt.Sscanf(scanner.Text(), "%d|%d", &num1, &num2)
			rules[num2] = append(rules[num2], num1)
			rulesReversed[num1] = append(rulesReversed[num1], num2)
		} else {
			parts := strings.Split(scanner.Text(), ",")
			var section []int
			for _, val := range parts {
				page, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				section = append(section, page)
			}
			pages = append(pages, section)
		}	
	}


}