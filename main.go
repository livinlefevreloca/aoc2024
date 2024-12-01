package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var problem int
	var part int
	var input string

	flag.IntVar(&problem, "P", 0, "The problem to run")
	flag.IntVar(&part, "p", 0, "The part to run")
	flag.StringVar(&input, "i", "", "The input file")

	flag.Parse()

	if problem == 0 || part == 0 || input == "" {
		panic("Missing required args")
	}

	if problem == 1 {
		if part == 1 {
			day1Part1(input)
		} else if part == 2 {
			day1Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	} else {
		fmt.Printf("Unknown problem: %d\n", problem)
	}

}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func day1Part1(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	col1 := make([]int, len(lines))
	col2 := make([]int, len(lines))
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			continue
		}
		num1, err := strconv.Atoi(numbers[0])
		check(err)

		num2, err := strconv.Atoi(numbers[1])
		check(err)

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	sort.Slice(col1, func(i, j int) bool { return col1[i] < col1[j] })
	sort.Slice(col2, func(i, j int) bool { return col2[i] < col2[j] })

	total := 0
	for i := 0; i < len(col1); i++ {
		total += absDiffInt(col1[i], col2[i])
	}

	fmt.Printf("Got total: %d\n", total)
}

func day1Part2(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	col1 := make([]int, len(lines))
	col2 := make([]int, len(lines))
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			continue
		}
		num1, err := strconv.Atoi(numbers[0])
		check(err)

		num2, err := strconv.Atoi(numbers[1])
		check(err)

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	counts := make(map[int]int)

	for _, num := range col2 {
		_, ok := counts[num]
		if ok {
			counts[num]++
		} else {
			counts[num] = 1
		}
	}

	total := 0
	for _, num := range col1 {
		appearances, ok := counts[num]
		if !ok {
			continue
		}
		total += num * appearances
	}

	fmt.Printf("Got total: %d\n", total)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(file string) string {
	data, err := os.ReadFile(file)
	check(err)
	return string(data)
}
