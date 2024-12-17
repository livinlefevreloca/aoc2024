package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var day int
	var part int
	var input string

	flag.IntVar(&day, "d", 0, "The problem to run")
	flag.IntVar(&part, "p", 0, "The part to run")
	flag.StringVar(&input, "i", "", "The input file")

	flag.Parse()

	switch day {
	case 1:
		if part == 1 {
			day1Part1(input)
		} else if part == 2 {
			day1Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 2:
		if part == 1 {
			day2Part1(input)
		} else if part == 2 {
			day2Part2(input)

		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 3:
		if part == 1 {
			day3Part1(input)
		} else if part == 2 {
			day3Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 4:
		if part == 1 {
			day4Part1(input)
		} else if part == 2 {
			day4Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 5:
		if part == 1 {
			day5Part1(input)
		} else if part == 2 {
			day5Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 6:
		if part == 1 {
			day6Part1(input)
		} else if part == 2 {
			day6Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 7:
		if part == 1 {
			day7Part1(input)
		} else if part == 2 {
			day7Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 8:
		if part == 1 {
			day8Part1(input)
		} else if part == 2 {
			day8Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 9:
		if part == 1 {
			day9Part1(input)
		} else if part == 2 {
			day9Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 10:
		if part == 1 {
			day10Part1(input)
		} else if part == 2 {
			day10Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 11:
		if part == 1 {
			day11Part1(input)
		} else if part == 2 {
			day11Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 12:
		if part == 1 {
			day12Part1(input)
		} else if part == 2 {
			day12Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 13:
		if part == 1 {
			day13Part1(input)
		} else if part == 2 {
			day13Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 14:
		if part == 1 {
			day14Part1(input)
		} else if part == 2 {
			day14Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 15:
		if part == 1 {
			day15Part1(input)
		} else if part == 2 {
			day15Part2(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	case 16:
		if part == 1 {
			day16Part1(input)
		} else if part == 2 {
			day16Part2Good(input)
		} else {
			fmt.Printf("Unknown part: %d\n", part)
		}
	default:
		fmt.Printf("Unknown day: %d\n", day)
	}

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

// Day 1
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

// Day 2

func day2Part1(inputFile string) {

	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	levels := make([][]int, 0)
	for _, line := range lines {
		level := make([]int, 0)
		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			check(err)
			level = append(level, n)
		}
		if len(level) > 0 {
			levels = append(levels, level)
		}
	}

	total := 0
	for level := range levels {
		if isSafe(levels[level]) {
			total++
		}
	}
	fmt.Printf("Got total: %d\n", total)
}

func day2Part2(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	levels := make([][]int, 0)
	for _, line := range lines {
		level := make([]int, 0)
		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			check(err)
			level = append(level, n)
		}
		if len(level) > 0 {
			levels = append(levels, level)
		}
	}

	total := 0
	for level := range levels {
		if isSafeWithRemoval(levels[level]) {
			total++
		}
	}
	fmt.Printf("Got total: %d\n", total)
}

func isSafe(level []int) bool {
	prev := level[0]
	direction := 0
	for _, num := range level[1:] {
		if direction == 0 {
			if num != prev {
				direction = (num - prev) / absDiffInt(prev, num)
			}
		} else {
			if num != prev && direction != (num-prev)/absDiffInt(prev, num) {
				return false
			}
		}
		if absDiffInt(prev, num) < 1 || absDiffInt(prev, num) > 3 {
			return false
		}
		prev = num
	}

	return true
}

func isSafeWithRemoval(level []int) bool {
	if !isSafe(level) {
		for i := 0; i < len(level); i++ {
			levelWithRemoval := make([]int, 0)
			for j, num := range level {
				if j != i {
					levelWithRemoval = append(levelWithRemoval, num)
				}
			}
			if isSafe(levelWithRemoval) {
				return true
			}
		}
		return false
	}
	return true
}

// Day 3

func day3Part1(inputFile string) {
	memory := readInput(inputFile)
	total := find_muls(memory)
	fmt.Printf("Got total: %d\n", total)
}

func find_muls(memory string) int {
	milRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	muls := milRe.FindAllStringSubmatch(memory, -1)
	total := 0
	for _, mul := range muls {
		a, err := strconv.Atoi(mul[1])
		check(err)
		b, err := strconv.Atoi(mul[2])
		check(err)
		total += a * b
	}

	return total
}

func day3Part2(inputFile string) {
	memory := readInput(inputFile)
	total := find_muls_do_dont(memory)
	fmt.Printf("Got total: %d\n", total)
}

type Range struct {
	lower int
	upper int
	on    bool
}

func (r *Range) contains(i int) bool {
	return i <= r.upper && i >= r.lower
}

func find_muls_do_dont(memory string) int {
	milRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)
	muls := milRe.FindAllStringSubmatch(memory, -1)
	mulIndexes := milRe.FindAllStringIndex(memory, -1)
	doIndexes := doRe.FindAllStringIndex(memory, -1)
	dontIndexes := dontRe.FindAllStringIndex(memory, -1)

	do := true

	mulPtr := 0
	doPtr := 0
	dontPtr := 0
	total := 0
	for {
		if mulPtr >= len(muls) {
			break
		}

		mulIdx := mulIndexes[mulPtr][0]
		var doIdx int
		var dontIdx int

		if doPtr >= len(doIndexes) {
			doIdx = len(memory)
		} else {
			doIdx = doIndexes[doPtr][0]
		}

		if dontPtr >= len(dontIndexes) {
			dontIdx = len(memory)
		} else {
			dontIdx = dontIndexes[dontPtr][0]
		}

		nextCandidates := []int{mulIdx, doIdx, dontIdx}
		minIdx := min_of_3(nextCandidates)
		if minIdx == 0 {
			if do {
				mul := muls[mulPtr]
				a, err := strconv.Atoi(mul[1])
				check(err)
				b, err := strconv.Atoi(mul[2])
				check(err)
				total += a * b
			}
			mulPtr += 1
		} else if minIdx == 1 {
			do = true
			doPtr += 1
		} else {
			do = false
			dontPtr += 1
		}
	}

	return total
}

func min_of_3(three []int) int {
	if three[0] < three[1] {
		if three[0] < three[2] {
			return 0
		}
		return 2
	}
	if three[1] < three[2] {
		return 1
	}

	return 2

}

// Day 4

func day4Part1(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, letter := range line {
			row = append(row, letter)
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}
	total := find_xmas(grid)

	fmt.Printf("Got total: %d\n", total)
}

func find_xmas(grid [][]rune) int {
	total := 0
	for i, row := range grid {
		for k, letter := range row {
			if letter == 'X' {
				total += check_xmas(grid, i, k)
			}
		}
	}
	return total
}

type Direction struct {
	x int
	y int
}

func check_xmas(grid [][]rune, i int, j int) int {
	total := 0
	xmas := []rune{'X', 'M', 'A', 'S'}
	for _, direction := range []Direction{{0, 1}, {1, 0}, {1, 1}, {1, -1}, {0, -1}, {-1, 0}, {-1, 1}, {-1, -1}} {
		match := true
		for k := 0; k < 4; k++ {
			if i+direction.x*k < 0 || i+direction.x*k >= len(grid) {
				match = false
				break
			} else if j+direction.y*k < 0 || j+direction.y*k >= len(grid[i]) {
				match = false
				break
			} else if grid[i+direction.x*k][j+direction.y*k] != xmas[k] {
				match = false
				break
			}
		}
		if match {
			total += 1
		}
	}
	return total
}

func day4Part2(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, letter := range line {
			row = append(row, letter)
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}
	total := find_mas(grid)
	fmt.Printf("Got total: %d\n", total)
}

func find_mas(grid [][]rune) int {
	total := 0
	for i, row := range grid {
		for k, letter := range row {
			if letter == 'A' {
				total += check_mas(grid, i, k)
			}
		}
	}
	return total
}

func check_mas(grid [][]rune, i int, j int) int {
	// This should be burned with fire
	if i+1 < len(grid) && j+1 < len(grid[0]) && j > 0 && i > 0 {
		if (grid[i+1][j+1] == 'S' && grid[i-1][j-1] == 'M' || grid[i+1][j+1] == 'M' && grid[i-1][j-1] == 'S') && (grid[i+1][j-1] == 'S' && grid[i-1][j+1] == 'M' || grid[i+1][j-1] == 'M' && grid[i-1][j+1] == 'S') {
			return 1
		}
	}
	return 0
}

// Day 5

func day5Part1(inputFile string) {
	input := readInput(inputFile)

	lines := strings.Split(input, "\n")

	ordering := make(map[string][]string, 0)
	updates := make([][]string, 0)
	for _, line := range lines {
		if strings.Contains(line, "|") {
			before_and_after := strings.Split(line, "|")
			before := before_and_after[0]
			after := before_and_after[1]
			if befores, ok := ordering[after]; ok {
				befores = append(befores, before)
				ordering[after] = befores
			} else {
				ordering[after] = []string{before}
			}
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		} else {
			continue
		}
	}
	updates = findCorrectUpdates(ordering, updates)
	total := 0
	for _, update := range updates {
		middle := len(update) / 2
		num, err := strconv.Atoi(update[middle])
		check(err)
		total += num
	}
	fmt.Printf("Got total: %d\n", total)
}

func findCorrectUpdates(ordering map[string][]string, updates [][]string) [][]string {
	incorrect := make([][]string, 0)
	for _, update := range updates {
		correct := true
	Update:
		for i, step := range update {
			if befores, ok := ordering[step]; ok {
				for _, before := range befores {
					if slices.Contains(update[i:], before) {
						correct = false
						break Update
					}
				}
			} else {
				continue
			}
		}
		if correct {
			incorrect = append(incorrect, update)
		}
	}

	return incorrect

}

func day5Part2(inputFile string) {
	input := readInput(inputFile)
	lines := strings.Split(input, "\n")
	ordering := make(map[string][]string, 0)
	updates := make([][]string, 0)
	for _, line := range lines {
		if strings.Contains(line, "|") {
			before_and_after := strings.Split(line, "|")
			before := before_and_after[0]
			after := before_and_after[1]
			if befores, ok := ordering[after]; ok {
				befores = append(befores, before)
				ordering[after] = befores
			} else {
				ordering[after] = []string{before}
			}
		} else if strings.Contains(line, ",") {
			updates = append(updates, strings.Split(line, ","))
		} else {
			continue
		}
	}
	updates = findIncorrectUpdates(ordering, updates)

	for _, update := range updates {
		for i, step := range update {
			befores, ok := ordering[step]
			if ok {
				for _, before := range befores {
					if slices.Contains(update[:i], before) {
						stepIdx := slices.Index(update, step)
						beforeIdx := slices.Index(update, before)
						for stepIdx > beforeIdx {
							update[stepIdx], update[stepIdx-1] = update[stepIdx-1], update[stepIdx]
							stepIdx--
						}
					}
				}
			}
		}
	}

	total := 0
	for _, update := range updates {
		middle := len(update) / 2
		num, err := strconv.Atoi(update[middle])
		check(err)
		total += num
	}
	fmt.Printf("Got total: %d\n", total)
}

func findIncorrectUpdates(ordering map[string][]string, updates [][]string) [][]string {
	incorrect := make([][]string, 0)
	for _, update := range updates {
		correct := true
	Update:
		for i, step := range update {
			if befores, ok := ordering[step]; ok {
				for _, before := range befores {
					if slices.Contains(update[i:], before) {
						correct = false
						break Update
					}
				}
			} else {
				continue
			}
		}
		if !correct {
			incorrect = append(incorrect, update)
		}
	}

	return incorrect

}

// Day 6

type Position struct {
	x  int
	y  int
	dx int
	dy int
}

func day6Part1(inputFile string) {
	input := readInput(inputFile)

	grid, current, height, width := parseGrid(input)

	_, positions := findPath(current, grid, height, width)

	unique := make(map[Position]bool, 0)
	for position := range positions {
		key := Position{position.x, position.y, 0, 0}
		if _, ok := unique[key]; !ok {
			unique[key] = true
		}
	}

	total := len(unique)

	fmt.Printf("Got total: %d\n", total)
}

func findPath(current Position, grid [][]bool, height int, width int) (bool, map[Position]bool) {
	positions := make(map[Position]bool, 0)
	positions[Position{current.x, current.x, current.x, current.y}] = true
	cycle := false
	for {
		next := Position{
			current.x + current.dx,
			current.y + current.dy,
			current.dx,
			current.dy,
		}
		if (next.x < 0) || (next.x >= height) || (next.y < 0) || (next.y >= width) {
			break
		}

		if grid[next.x][next.y] {
			// 90 degree clockwise rotation
			current.dx, current.dy = current.dx*0+current.dy*1, current.dx*-1+current.dy*0
		} else {
			current = Position{current.x + current.dx, current.y + current.dy, current.dx, current.dy}
		}

		_, ok := positions[current]
		if !ok {
			positions[current] = true
		} else {
			cycle = true
			break
		}
	}

	return cycle, positions
}

func parseGrid(input string) ([][]bool, Position, int, int) {

	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	height := len(lines)
	width := len(lines[0])
	grid := make([][]bool, height)

	for i := 0; i < height; i++ {
		grid[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			grid[i][j] = false
		}
	}

	current := Position{}
	for i, line := range lines {
		for j, letter := range line {
			if letter == '.' {
			} else if letter == '#' {
				grid[i][j] = true
			} else if letter == '^' {
				current = Position{i, j, -1, 0}
			} else {
				continue
			}
		}
	}

	return grid, current, height, width

}

func day6Part2(inputFile string) {
	input := readInput(inputFile)
	grid, current, height, width := parseGrid(input)

	start := current
	_, initial_path := findPath(current, grid, height, width)
	initial_positions := make(map[Position]bool, 0)
	for position := range initial_path {
		key := Position{position.x, position.y, 0, 0}
		if _, ok := initial_positions[key]; !ok {
			initial_positions[key] = true
		}
	}

	cycles := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == start.x && j == start.y {
				continue
			}

			if _, ok := initial_positions[Position{i, j, 0, 0}]; !ok {
				continue
			}

			current = start
			old := grid[i][j]
			grid[i][j] = true
			cycle, _ := findPath(current, grid, height, width)
			if cycle {
				cycles++
			}
			grid[i][j] = old
		}
	}

	fmt.Printf("Got total: %d\n", cycles)
}

// Day 7
func day7Part1(inputFile string) {
	input := readInput(inputFile)
	calibrations := parseCalibration(input)
	operations := map[int]func(int, int) int{0: sum, 1: product}
	total := 0
	for cal, values := range calibrations {
		if isValidCalibration(cal, values, operations) {
			total += cal
		}
	}

	fmt.Printf("Got total: %d\n", total)
}

func day7Part2(inputFile string) {
	input := readInput(inputFile)
	calibrations := parseCalibration(input)
	total := 0
	operations := map[int]func(int, int) int{0: sum, 1: product, 2: concat}
	for cal, values := range calibrations {
		if isValidCalibration(cal, values, operations) {
			total += cal
		}
	}

	fmt.Printf("Got total: %d\n", total)
}

func parseCalibration(input string) map[int][]int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	calibrations := make(map[int][]int, 0)
	for _, line := range lines {
		calibration := strings.Split(line, ":")
		key_str := calibration[0]
		value_strs := strings.Split(strings.Trim(calibration[1], " "), " ")
		key, err := strconv.Atoi(key_str)
		check(err)
		values := make([]int, 0)
		for _, value_str := range value_strs {
			value, err := strconv.Atoi(value_str)
			check(err)
			values = append(values, value)
		}
		calibrations[key] = values
	}

	return calibrations
}

func isValidCalibration(calibration int, values []int, operations map[int]func(int, int) int) bool {
	currentStates := make([]int, 0)
	currentStates = append(currentStates, values[0])
	for _, value := range values[1:] {
		newStates := make([]int, 0)
		for _, state := range currentStates {
			for _, op := range operations {
				newState := op(state, value)
				if newState > calibration {
					continue
				}
				newStates = append(newStates, newState)
			}
		}
		currentStates = newStates
	}

	for _, state := range currentStates {
		if state == calibration {
			return true
		}
	}
	return false
}

func sum(a int, b int) int {
	return a + b
}

func product(a int, b int) int {
	return a * b
}

func concat(a int, b int) int {
	return a*10 ^ int(math.Log10(float64(b))+1) + b
}

func getOperation(i int, opIdx int) int {
	return (i >> opIdx) & 1
}

// Day 8
type Location struct {
	x int
	y int
}

type Slope struct {
	dx int
	dy int
}

func day8Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	locationMap := make(map[string][]Location, 0)
	antiNodes := make(map[Location]bool, 0)
	for i, line := range lines {
		for j, letter := range strings.Split(line, "") {
			if letter != "." {
				newLocation := Location{i, j}
				if _, ok := locationMap[letter]; ok {
					for _, location := range locationMap[letter] {
						slope := calculateSlope(newLocation, location)
						negativeSlope := Slope{-slope.dx, -slope.dy}
						pairs := []struct {
							loc   Location
							slope Slope
						}{
							{newLocation, slope},
							{newLocation, negativeSlope},
							{location, slope},
							{location, negativeSlope},
						}
						for _, pair := range pairs {
							checkLocation := Location{pair.loc.x + pair.slope.dx, pair.loc.y + pair.slope.dy}
							if isInBounds(checkLocation, len(lines), len(lines[0])) {
								if checkLocation != newLocation && checkLocation != location {
									antiNodes[checkLocation] = true
								}
							}
						}

					}
					locationMap[letter] = append(locationMap[letter], newLocation)
				} else {
					locationMap[letter] = []Location{newLocation}
				}
			}
		}
	}

	for i, line := range lines {

		for _, letter := range strings.Split(line, "") {
			fmt.Printf("%s", letter)
		}
		fmt.Printf(" ")
		for j, letter := range strings.Split(line, "") {
			if _, ok := antiNodes[Location{i, j}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf("%s", letter)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Got total: %d\n", len(antiNodes))
}

func isInBounds(location Location, height int, width int) bool {
	return location.x >= 0 && location.x < height && location.y >= 0 && location.y < width
}

func calculateSlope(location1 Location, location2 Location) Slope {
	dx := location2.x - location1.x
	dy := location2.y - location1.y
	return Slope{dx, dy}
}

func day8Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	locationMap := make(map[string][]Location, 0)
	antiNodes := make(map[Location]bool, 0)
	for i, line := range lines {
		for j, letter := range strings.Split(line, "") {
			if letter != "." {
				newLocation := Location{i, j}
				if _, ok := locationMap[letter]; ok {
					for _, location := range locationMap[letter] {
						slope := calculateSlope(newLocation, location)
						negativeSlope := Slope{-slope.dx, -slope.dy}
						offset := 1
						for {
							locationsToCheck := []Location{
								{newLocation.x + slope.dx*offset, newLocation.y + slope.dy*offset},
								{newLocation.x + negativeSlope.dx*offset, newLocation.y + negativeSlope.dy*offset},
								{location.x + slope.dx*offset, location.y + slope.dy*offset},
								{location.x + negativeSlope.dx*offset, location.y + negativeSlope.dy*offset},
							}
							outOfBoundsCount := 0
							for _, checkLocation := range locationsToCheck {
								if isInBounds(checkLocation, len(lines), len(lines[0])) {
									antiNodes[checkLocation] = true
								} else {
									outOfBoundsCount++
								}
							}
							if outOfBoundsCount == 4 {
								break
							}
							offset++
						}
					}
					locationMap[letter] = append(locationMap[letter], newLocation)
				} else {
					locationMap[letter] = []Location{newLocation}
				}
			}
		}
	}

	for i, line := range lines {

		for _, letter := range strings.Split(line, "") {
			fmt.Printf("%s", letter)
		}
		fmt.Printf(" ")
		for j, letter := range strings.Split(line, "") {
			if _, ok := antiNodes[Location{i, j}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf("%s", letter)
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("Got total: %d\n", len(antiNodes))
}

// Day 9
func day9Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	fileArray := make([]int, 0)
	isBlock := true
	fileIdx := 0
	for _, letter := range strings.Split(trimmed, "") {
		num, err := strconv.Atoi(letter)
		check(err)
		insert := -1
		if isBlock {
			insert = fileIdx
			fileIdx++
		}
		for i := 0; i < num; i++ {
			fileArray = append(fileArray, insert)
		}
		isBlock = !isBlock
	}

	free := 0
	for i := len(fileArray) - 1; i > 0; i-- {
		if fileArray[i] == -1 {
			continue
		}
		foundEnd := false
		for fileArray[free] != -1 {
			free++
			if free >= i {
				foundEnd = true
				break
			}
		}
		if foundEnd {
			break
		}
		fileArray[free], fileArray[i] = fileArray[i], fileArray[free]
	}

	checksum := 0
	for i := 0; i < len(fileArray); i++ {
		if fileArray[i] == -1 {
			continue
		}
		checksum += fileArray[i] * i
	}

	fmt.Printf("Got checksum: %d\n", checksum)

}

type Space struct {
	start int
	end   int
}

func day9Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	fileArray := make([]int, 0)
	isBlock := true
	fileIdx := 0
	for _, letter := range strings.Split(trimmed, "") {
		num, err := strconv.Atoi(letter)
		check(err)
		insert := -1
		if isBlock {
			insert = fileIdx
			fileIdx++
		}
		for i := 0; i < num; i++ {
			fileArray = append(fileArray, insert)
		}
		isBlock = !isBlock
	}

	start := 0
	end := len(fileArray) - 1
	for end > 0 {
		// Find the file directtly to the left of end
		currentFile := Space{end, end}
		for currentFile.start-1 >= 0 && fileArray[currentFile.start-1] == fileArray[currentFile.end] {
			currentFile.start--
		}

		// Find the nearest freespace to right of start
		freeSpace := Space{start, start}
		for freeSpace.start < len(fileArray) && fileArray[freeSpace.start] != -1 {
			freeSpace.start++
		}
		freeSpace.end = freeSpace.start
		for freeSpace.end+1 < len(fileArray) && fileArray[freeSpace.end+1] == -1 {
			freeSpace.end++
		}

		// If the freespace overlaps or is to the right of the current file
		// The file cannot be moved to the left becuase there are no sufficienly
		// large freespace to the left. In this case we reset start to the
		// beginning of the array and move to so that it points to the end of the
		// next file after the one that could not be moved
		if freeSpace.end >= currentFile.start {
			start = 0
			end = currentFile.start - 1
			for end >= 0 && fileArray[end] == -1 {
				end--
			}
			continue
		}

		// Calculate the size of the file and the freespace
		fileSize := currentFile.end - currentFile.start + 1
		freeSize := freeSpace.end - freeSpace.start + 1

		// If the freespace is larger than the file, move the file to the freespace
		if freeSize >= fileSize {
			for i := 0; i < fileSize; i++ {
				fileArray[freeSpace.start+i] = fileArray[currentFile.start+i]
				fileArray[currentFile.start+i] = -1
			}
			start = 0
			end = currentFile.start - 1
			for end >= 0 && fileArray[end] == -1 {
				end--
			}
			// If the freespace is smaller than the file, increment start past the
			// freespace so we can try to find a larger one further to the right in the array
		} else {
			start = freeSpace.end + 1
			if start >= end {
				end = currentFile.start - 1
				for end >= 0 && fileArray[end] == -1 {
					end--
				}
				start = 0
				continue
			}
		}

	}

	// Once we have compacted the file compute the checksum
	checksum := 0
	for i := 0; i < len(fileArray); i++ {
		if fileArray[i] == -1 {
			continue
		}
		checksum += fileArray[i] * i
	}

	fmt.Printf("Got checksum: %d\n", checksum)
}

// Use Memoization of freespace sizes to speed up the process.
// Original solution took ~260ms, this one takes 25ms
func day9Part2Improved(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	fileArray := make([]int, 0)
	isBlock := true
	fileIdx := 0
	for _, letter := range strings.Split(trimmed, "") {
		num, err := strconv.Atoi(letter)
		check(err)
		insert := -1
		if isBlock {
			insert = fileIdx
			fileIdx++
		}
		for i := 0; i < num; i++ {
			fileArray = append(fileArray, insert)
		}
		isBlock = !isBlock
	}

	// Build a map of free spaces
	freeSpacesMap := make(map[int][]Space, 0)
	for i := 0; i < len(fileArray); i++ {
		if fileArray[i] != -1 {
			continue
		}
		space := Space{i, i}
		for i < len(fileArray) && fileArray[i] == -1 {
			space.end = i
			i++
		}
		size := space.end - space.start + 1
		if _, ok := freeSpacesMap[size]; ok {
			freeSpacesMap[size] = append(freeSpacesMap[size], space)
			slices.SortFunc(freeSpacesMap[size], func(i, j Space) int {
				return i.start - j.start
			})

		} else {
			freeSpacesMap[size] = []Space{space}
		}
	}

	// While there are files to move
	end := len(fileArray) - 1
	for end > 0 {
		// Find the rightmost most file from end
		currentFile := Space{end, end}
		for currentFile.start-1 >= 0 && fileArray[currentFile.start-1] == fileArray[currentFile.end] {
			currentFile.start--
		}
		fileSize := currentFile.end - currentFile.start + 1

		// Check if there is a free space of the same size or larger
		size := fileSize
		freeSpaces, ok := freeSpacesMap[size]
		for !ok && size <= len(freeSpacesMap) {
			size++
			freeSpaces, ok = freeSpacesMap[size]
		}

		// If there is no free space of the same size or larger, or
		// the freespace found is too the right of the current file
		// We skip the file. we know there will not be another free space
		// since we keep them in sorted order by start index
		if !ok || len(freeSpaces) == 0 || freeSpaces[0].start > currentFile.start {
			end = currentFile.start - 1
			for end >= 0 && fileArray[end] == -1 {
				end--
			}
			continue
		}

		// Remove the matched freespace from  its size group
		freeSpace := freeSpaces[0]
		freeSpacesMap[size] = freeSpaces[1:]

		// Swap the file into the free space counting down the
		// number of free blocks filled in with the file
		freeSize := freeSpace.end - freeSpace.start + 1
		for i := 0; i < fileSize; i++ {
			fileArray[freeSpace.start] = fileArray[currentFile.start+i]
			fileArray[currentFile.start+i] = -1
			freeSize--
			freeSpace.start++
		}

		// Update the free space map with the new free space
		// putting it into the correct size group and keeping
		// the groups sorted by start index
		spaces := freeSpacesMap[freeSize]
		for i, space := range spaces {
			if space.start > freeSpace.start {
				freeSpacesMap[freeSize] = append(spaces[:i], append([]Space{freeSpace}, spaces[i:]...)...)
				break
			}
		}

		// move end to the end of the next non freespace block
		end = currentFile.start - 1
		for end >= 0 && fileArray[end] == -1 {
			end--
		}
	}

	// Calculate the checksum
	checksum := 0
	for i := 0; i < len(fileArray); i++ {
		if fileArray[i] == -1 {
			continue
		}
		checksum += fileArray[i] * i
	}

	fmt.Printf("Got checksum: %d\n", checksum)
}

// Day 10

type Pair struct {
	x int
	y int
}

func day10Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	grid := make([][]int, 0)
	possibleTrailHeads := make([]Pair, 0)
	for i, line := range lines {
		row := make([]int, 0)
		for j, num := range strings.Split(line, "") {
			n, err := strconv.Atoi(num)
			if num == "0" {
				possibleTrailHeads = append(possibleTrailHeads, Pair{i, j})
			}
			check(err)
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	total := 0
	for _, trailHead := range possibleTrailHeads {
		score := getScore(grid, trailHead)
		total += score
	}

	fmt.Printf("Got total: %d\n", total)

}

func getScore(grid [][]int, location Pair) int {
	currentStates := make(map[Pair]bool, 0)
	currentStates[location] = true
	trailHeads := make(map[Pair]bool, 0)

	for len(currentStates) > 0 {
		newStates := make(map[Pair]bool, 0)
		for state := range currentStates {
			for _, delta := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Pair{state.x + delta.x, state.y + delta.y}
				if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
					continue
				}

				if grid[next.x][next.y]-grid[state.x][state.y] == 1 {
					if grid[next.x][next.y] == 9 {
						trailHeads[next] = true
					} else {
						newStates[next] = true
					}
				}
			}
		}
		currentStates = newStates
	}

	return len(trailHeads)
}

func day10Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	grid := make([][]int, 0)
	possibleTrailHeads := make([]Pair, 0)
	for i, line := range lines {
		row := make([]int, 0)
		for j, num := range strings.Split(line, "") {
			n, err := strconv.Atoi(num)
			if num == "0" {
				possibleTrailHeads = append(possibleTrailHeads, Pair{i, j})
			}
			check(err)
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	total := 0
	for _, trailHead := range possibleTrailHeads {
		score := getTrailScore(grid, trailHead)
		total += score
	}

	fmt.Printf("Got total: %d\n", total)

}

func getTrailScore(grid [][]int, location Pair) int {
	currentTrails := make([][]Pair, 0)
	currentTrails = append(currentTrails, []Pair{location})
	allTrails := make([][]Pair, 0)

	for len(currentTrails) > 0 {
		newTrails := make([][]Pair, 0)
		for _, trail := range currentTrails {
			state := trail[len(trail)-1]
			for _, delta := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Pair{state.x + delta.x, state.y + delta.y}
				if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
					continue
				}

				if grid[next.x][next.y]-grid[state.x][state.y] == 1 {
					if grid[next.x][next.y] == 9 {
						newTrail := make([]Pair, len(trail))
						copy(newTrail, trail)
						newTrail = append(newTrail, next)
						allTrails = append(allTrails, append(trail, next))
					} else {
						newTrail := make([]Pair, len(trail))
						copy(newTrail, trail)
						newTrail = append(newTrail, next)
						newTrails = append(newTrails, newTrail)
					}
				}
			}
		}
		currentTrails = newTrails
	}
	return len(allTrails)

}

func day11Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	nums := strings.Split(trimmed, " ")
	total := getStoneCounts(nums, 25)
	fmt.Printf("Got total: %d\n", total)
}

func day11Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	nums := strings.Split(trimmed, " ")
	total := getStoneCounts(nums, 75)
	fmt.Printf("Got total: %d\n", total)
}

func getStoneCounts(nums []string, iterations int) int {
	counts := make(map[string]int, 0)
	for _, num := range nums {
		if _, ok := counts[num]; ok {
			counts[num]++
		} else {
			counts[num] = 1
		}
	}
	for i := 0; i < iterations; i++ {
		newCounts := make(map[string]int, 0)
		for num, count := range counts {
			if num == "0" {
				if _, ok := newCounts["1"]; ok {
					newCounts["1"] += count
				} else {
					newCounts["1"] = count
				}
			} else if len(num)%2 == 0 {
				num1, num2 := num[:len(num)/2], num[len(num)/2:]
				number1, err := strconv.Atoi(num1)
				check(err)
				number2, err := strconv.Atoi(num2)
				check(err)
				num1, num2 = strconv.Itoa(number1), strconv.Itoa(number2)
				if _, ok := newCounts[num1]; ok {
					newCounts[num1] += count
				} else {
					newCounts[num1] = count
				}
				if _, ok := newCounts[num2]; ok {
					newCounts[num2] += count
				} else {
					newCounts[num2] = count
				}
			} else {
				number, err := strconv.Atoi(num)
				check(err)
				newNumber := number * 2024
				newNum := strconv.Itoa(newNumber)
				if _, ok := newCounts[newNum]; ok {
					newCounts[newNum] += count
				} else {
					newCounts[newNum] = count
				}
			}
		}
		counts = newCounts
	}
	total := 0
	for _, count := range counts {
		total += count
	}

	return total
}

// Day 12

func day12Part1(inputFile string) {

	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, letter := range line {
			row = append(row, letter)
		}
		grid = append(grid, row)
	}

	plots := make([]map[Pair]rune, 0)
	usedPlots := make(map[Pair]bool, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			plot := make(map[Pair]rune, 0)
			current := Pair{i, j}
			if _, ok := usedPlots[current]; ok {
				continue
			}
			walkPlot(grid, plot, current, usedPlots)
			plots = append(plots, plot)
		}
	}

	total := 0
	for _, plot := range plots {
		fmt.Printf("Got plot: %v\n", plot)
		area := len(plot)
		perimeter := calcPerimeter(plot)
		total += area * perimeter
	}
	fmt.Printf("Got total: %d\n", total)

}

func walkPlot(grid [][]rune, plot map[Pair]rune, current Pair, seen map[Pair]bool) {
	if _, ok := plot[current]; !ok {
		plot[current] = grid[current.x][current.y]
	} else {
		return
	}
	for _, delta := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		next := Pair{current.x + delta.x, current.y + delta.y}
		if _, ok := seen[next]; ok || isOutOfBounds(next, len(grid), len(grid[0])) {
			continue
		}
		if grid[next.x][next.y] == plot[current] {
			seen[next] = true
			walkPlot(grid, plot, next, seen)
		}
	}
}

func isOutOfBounds(point Pair, max_x int, max_y int) bool {
	return point.x < 0 || point.x >= max_x || point.y < 0 || point.y >= max_y
}

func calcPerimeter(plot map[Pair]rune) int {
	perimeter := 0
	for point := range plot {
		for _, delta := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := Pair{point.x + delta.x, point.y + delta.y}
			if _, ok := plot[next]; !ok {
				perimeter++
			}
		}
	}
	return perimeter
}

func day12Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	grid := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, letter := range line {
			row = append(row, letter)
		}
		grid = append(grid, row)
	}

	plots := make([]map[Pair]rune, 0)
	usedPlots := make(map[Pair]bool, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			plot := make(map[Pair]rune, 0)
			current := Pair{i, j}
			if _, ok := usedPlots[current]; ok {
				continue
			} else {
				usedPlots[current] = true
			}

			if _, ok := plot[current]; !ok {
				plot[current] = grid[current.x][current.y]
			}

			searcher := Pair{current.x, current.y}
			walkPlot(grid, plot, searcher, usedPlots)
			plots = append(plots, plot)
		}
	}

	total := 0
	for _, plot := range plots {
		area := len(plot)
		sides := calcSides(plot, len(grid), len(grid[0]))
		total += area * sides
	}
	fmt.Printf("Got total: %d\n", total)
}

func calcSides(plot map[Pair]rune, max_x int, max_y int) int {
	pxs := make(map[int][]Pair, 0)
	pys := make(map[int][]Pair, 0)
	nxs := make(map[int][]Pair, 0)
	nys := make(map[int][]Pair, 0)
	for point := range plot {
		for _, delta := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			next := Pair{point.x + delta.x, point.y + delta.y}
			if _, ok := plot[next]; !ok || isOutOfBounds(next, max_x, max_y) {
				if delta.y == 1 {
					if side, ok := pys[next.y]; !ok {
						side := []Pair{next}
						pys[next.y] = side
					} else {
						pys[next.y] = append(side, next)
					}
				} else if delta.y == -1 {
					if side, ok := nys[next.y]; !ok {
						side := []Pair{next}
						nys[next.y] = side
					} else {
						nys[next.y] = append(side, next)
					}
				} else if delta.x == 1 {
					if side, ok := pxs[next.x]; !ok {
						side := []Pair{next}
						pxs[next.x] = side
					} else {
						pxs[next.x] = append(side, next)
					}
				} else if delta.x == -1 {
					if side, ok := nxs[next.x]; !ok {
						side := []Pair{next}
						nxs[next.x] = side
					} else {
						nxs[next.x] = append(side, next)
					}
				}

			}
		}
	}

	sides := make([][]Pair, 0)
	for _, side := range pxs {
		splitSides := splitSide(side, func(a, b Pair) int {
			return a.y - b.y
		})
		sides = append(sides, splitSides...)
	}
	for _, side := range pys {
		splitSides := splitSide(side, func(a, b Pair) int {
			return a.x - b.x
		})
		sides = append(sides, splitSides...)
	}
	for _, side := range nxs {
		splitSides := splitSide(side, func(a, b Pair) int {
			return a.y - b.y
		})
		sides = append(sides, splitSides...)
	}
	for _, side := range nys {
		splitSides := splitSide(side, func(a, b Pair) int {
			return a.x - b.x
		})
		sides = append(sides, splitSides...)
	}

	return len(sides)
}

func splitSide(side []Pair, sortFunc func(a, b Pair) int) [][]Pair {
	slices.SortFunc(side, sortFunc)
	splitSides := make([][]Pair, 0)
	lastSplit := 0
	for i := 0; i < len(side); i++ {
		if i == 0 {
			continue
		}
		if absDiffInt(side[i].y, side[i-1].y) > 1 || absDiffInt(side[i].x, side[i-1].x) > 1 {
			splitSides = append(splitSides, side[:i])
			lastSplit = i
		}
	}
	splitSides = append(splitSides, side[lastSplit:])

	return splitSides
}

// Day 13

type Button struct {
	x int
	y int
}

type Machine struct {
	A     Button
	B     Button
	Prize Pair
}

func day13Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")

	total := 0
	for _, chunk := range strings.Split(trimmed, "\n\n") {
		machine := parseMachine(chunk, 0)
		fmt.Printf("Got machine: %v\n", machine)
		pushes := solveMachine(machine, []string{}, Pair{0, 0}, make(map[Pair]bool))
		counts := make(map[string]int, 0)
		for _, push := range pushes {
			if _, ok := counts[push]; ok {
				counts[push]++
			} else {
				counts[push] = 1
			}
		}
		total += counts["A"]*3 + counts["B"]
	}

	fmt.Printf("Got total: %d\n", total)

}

func solveMachine(machine Machine, pushes []string, state Pair, seen map[Pair]bool) []string {
	if _, ok := seen[state]; ok {
		return []string{}
	} else {
		seen[state] = true
	}
	if state == machine.Prize {
		return pushes
	} else if state.x > machine.Prize.x || state.y > machine.Prize.y {
		return []string{}
	} else {
		pushes = append(pushes, "A")
		pushed := solveMachine(machine, pushes, Pair{state.x + machine.A.x, state.y + machine.A.y}, seen)
		if len(pushed) > 0 {
			return pushed
		}
		pushes := pushes[:len(pushes)-1]
		pushes = append(pushes, "B")
		pushed = solveMachine(machine, pushes, Pair{state.x + machine.B.x, state.y + machine.B.y}, seen)
		if len(pushed) > 0 {
			return pushed
		}
	}

	return []string{}
}

func parseMachine(input string, offset int) Machine {
	machine := Machine{}
	xre := regexp.MustCompile(`X.(?P<x>\d+)`)
	yre := regexp.MustCompile(`Y.(?P<y>\d+)`)
	for i, line := range strings.Split(input, "\n") {
		xstr := xre.FindStringSubmatch(line)
		ystr := yre.FindStringSubmatch(line)
		if i == 0 {
			x, err := strconv.Atoi(xstr[1])
			check(err)
			y, err := strconv.Atoi(ystr[1])
			machine.A = Button{x, y}
		} else if i == 1 {
			x, err := strconv.Atoi(xstr[1])
			check(err)
			y, err := strconv.Atoi(ystr[1])
			machine.B = Button{x, y}
		} else {
			x, err := strconv.Atoi(xstr[1])
			check(err)
			y, err := strconv.Atoi(ystr[1])
			machine.Prize = Pair{x + offset, y + offset}
		}
	}
	return machine
}

func day13Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")

	total := 0
	for _, chunk := range strings.Split(trimmed, "\n\n") {
		machine := parseMachine(chunk, 10000000000000)
		fmt.Printf("Got machine: %v\n", machine)
		as, bs := solveMachine2(machine)
		fmt.Printf("Got as: %d, bs: %d\n", as, bs)
		if as > 0 && bs > 0 {
			total += as*3 + bs
		}
	}

	fmt.Printf("Got total: %d\n", total)

}

func solveMachine2(machine Machine) (int, int) {
	Ay := machine.A.y
	Ax := machine.A.x
	By := machine.B.y
	Bx := machine.B.x
	PrizeY := machine.Prize.y
	PrizeX := machine.Prize.x

	j := (PrizeY*Ax - Ay*PrizeX) / (Ax*By - Ay*Bx)
	i := (PrizeX - Bx*j) / Ax

	if Ax*i+Bx*j == PrizeX && Ay*i+By*j == PrizeY {
		return i, j
	}

	return -1, -1
}

// Day 14

type Robot struct {
	x  int
	y  int
	dx int
	dy int
}

func day14Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	robots := make([]Robot, 0)
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	// fmt.Printf("Got robots: %v\n", robots)
	grid := Pair{101, 103}
	for i := 0; i < 100; i++ {
		robots = simulate(robots, grid)
	}

	printGrid(robots, grid)

	middle_x := grid.x / 2
	middle_y := grid.y / 2
	qauds := make(map[int]int, 4)
	for i := 0; i < 4; i++ {
		qauds[i] = 0
	}
	for _, robot := range robots {
		if robot.x == middle_x || robot.y == middle_y {
			continue
		}
		if robot.x < middle_x && robot.y < middle_y {
			qauds[0]++
		} else if robot.x > middle_x && robot.y < middle_y {
			qauds[1]++
		} else if robot.x > middle_x && robot.y > middle_y {
			qauds[2]++
		} else {
			qauds[3]++
		}
	}
	product := 1
	for _, count := range qauds {
		product *= count
	}
	fmt.Printf("Got product: %d\n", product)

}
func parseRobot(line string) Robot {
	first_second := strings.Split(line, " ")
	x_y := strings.Split(first_second[0][2:], ",")
	x, err := strconv.Atoi(x_y[0])
	check(err)
	y, err := strconv.Atoi(x_y[1])
	check(err)
	dx_dy := strings.Split(first_second[1][2:], ",")
	dx, err := strconv.Atoi(dx_dy[0])
	check(err)
	dy, err := strconv.Atoi(dx_dy[1])
	check(err)
	return Robot{x, y, dx, dy}
}

func simulate(robots []Robot, grid Pair) []Robot {
	newRobots := make([]Robot, 0)
	for _, robot := range robots {
		if robot.x+robot.dx >= grid.x {
			robot.x = robot.dx - (grid.x - robot.x)
		} else if robot.x+robot.dx < 0 {
			robot.x = grid.x + (robot.dx + robot.x)
		} else {
			robot.x += robot.dx
		}

		if robot.y+robot.dy >= grid.y {
			robot.y = robot.dy - (grid.y - robot.y)
		} else if robot.y+robot.dy < 0 {
			robot.y = grid.y + (robot.dy + robot.y)
		} else {
			robot.y += robot.dy
		}
		newRobots = append(newRobots, robot)
	}

	return newRobots
}

func day14Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	robots := make([]Robot, 0)
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	// fmt.Printf("Got robots: %v\n", robots)
	postitions := make(map[int][]Robot, 0)
	grid := Pair{101, 103}
	for i := 0; i < 10403; i++ {
		robots = simulate(robots, grid)
		postitions[i] = robots
	}

	distances := make(map[float64]int, 0)
	count := 0
	for i, robots := range postitions {
		fmt.Printf("Calcualting iteration: %d, count: %d\n", i, count)
		count++
		robotMap := make(map[Pair]bool, 0)
		for _, robot := range robots {
			robotMap[Pair{robot.x, robot.y}] = true
		}
		middle_x := grid.x / 2
		middle_y := grid.y / 2
		ssd := 0.0
		for _, robot := range robots {
			if robot.x == middle_x || robot.y == middle_y {
				continue
			}
			ssd += math.Pow(float64(robot.x-middle_x), 2) + math.Pow(float64(robot.y-middle_y), 2)
		}
		distances[ssd] = i
	}

	minSSD := math.Pow(2, 31) - 1
	for distance := range distances {
		if distance < minSSD {
			minSSD = distance
		}

	}

	iteration := distances[minSSD]
	fmt.Printf("Got iteration: %d\n", iteration)
	printGrid(postitions[iteration], grid)
}

func pointDistance(p1, p2 Robot) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

func printGrid(robots []Robot, grid Pair) {
	for j := 20; j < grid.y-50; j++ {
		for k := 41; k < grid.x-29; k++ {
			count := 0
			for _, robot := range robots {
				if robot.x == k && robot.y == j {
					count++
				}
			}
			if count == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("X")
			}
		}
		fmt.Printf("\n")
	}
}

// Day 15
func day15Part1(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")

	parts := strings.Split(trimmed, "\n\n")
	grid := make([][]string, 0)
	var robot Pair
	for i, line := range strings.Split(parts[0], "\n") {
		row := make([]string, 0)
		for j, letter := range strings.Split(line, "") {
			if letter == "@" {
				robot = Pair{i, j}
				row = append(row, "@")
			} else {
				row = append(row, letter)
			}
		}
		grid = append(grid, row)
	}

	moves := parseMoves(parts[1])

	for _, move := range moves {
		next := Pair{robot.x + move.x, robot.y + move.y}
		// If the next position is a wall, we skip
		if grid[next.x][next.y] == "#" {
			continue
		} else if grid[next.x][next.y] == "O" {
			// if the next position is a box, we check for aligned boxes
			// until we hit a wall or open space. We then move all the boxes
			// found on the line by one step in the direction of the move
			toMove := make([]Pair, 0)
			toMove = append(toMove, next)
			nextNext := Pair{next.x + move.x, next.y + move.y}
			for grid[nextNext.x][nextNext.y] == "O" {
				toMove = append(toMove, nextNext)
				nextNext = Pair{nextNext.x + move.x, nextNext.y + move.y}
			}
			// if nextNext is a wall after exiting the loop, we cant move
			// anything on the line so we skip
			if grid[nextNext.x][nextNext.y] == "#" {
				continue
			} else {
				// If nextNext is open space, we move all the boxes on the line
				// by one step in the direction of the move
				for i := len(toMove) - 1; i >= 0; i-- {
					box := toMove[i]
					grid[box.x][box.y] = "."
					grid[box.x+move.x][box.y+move.y] = "O"
				}
				// We then move the robot to the next position
				grid[next.x][next.y] = "@"
				grid[robot.x][robot.y] = "."
				robot = next
			}
		} else {
			// if the next position is open space, we move the
			// robot to the next position
			grid[next.x][next.y] = "@"
			grid[robot.x][robot.y] = "."
			robot = next
		}
	}
	printGrid15(grid, robot)
	sum := 0
	for i, row := range grid {
		for j, letter := range row {
			if letter == "O" {
				sum += i*100 + j
			}
		}
	}

	fmt.Printf("Got sum: %d\n", sum)
}

func parseMoves(input string) []Pair {
	moves := make([]Pair, 0)
	for _, r := range strings.Split(input, "") {
		if r == "^" {
			moves = append(moves, Pair{-1, 0})
		} else if r == "v" {
			moves = append(moves, Pair{1, 0})
		} else if r == "<" {
			moves = append(moves, Pair{0, -1})
		} else if r == ">" {
			moves = append(moves, Pair{0, 1})
		}
	}

	return moves
}

func printGrid15(grid [][]string, robot Pair) {
	for _, row := range grid {
		for _, letter := range row {
			fmt.Printf("%s", letter)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func day15Part2(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")

	parts := strings.Split(trimmed, "\n\n")
	grid := make([][]string, 0)
	var robot Pair
	for i, line := range strings.Split(parts[0], "\n") {
		row := make([]string, 0)
		for j, letter := range strings.Split(line, "") {
			if letter == "@" {
				robot = Pair{i, j * 2}
				row = append(row, "@")
				row = append(row, ".")
			} else if letter == "#" {
				row = append(row, "#")
				row = append(row, "#")
			} else if letter == "." {
				row = append(row, ".")
				row = append(row, ".")
			} else if letter == "O" {
				row = append(row, "[")
				row = append(row, "]")
			}
		}
		grid = append(grid, row)
	}

	moves := parseMoves(parts[1])

	for _, move := range moves {
		next := Pair{robot.x + move.x, robot.y + move.y}
		// If the next position is a wall, we skip
		if grid[next.x][next.y] == "#" {
			continue
		} else if grid[next.x][next.y] == "[" || grid[next.x][next.y] == "]" {
			// if the next position is a part of a box we check the direction
			// we are pushing in and take the appropriate action
			if move.x == 0 {
				// If we are pushing in the x direction, we can follow
				// a procedure similar to the one in part 1 where we keep
				// track of boxes on the line until we hit a wall or open space
				// and then move all the boxes on the line by one step in the
				toMove := make([]Pair, 0)
				toMove = append(toMove, next)
				i := 0
				for grid[next.x][next.y+move.y*i] != "#" && grid[next.x][next.y+move.y*i] != "." {
					i++
				}
				if grid[next.x][next.y+move.y*i] == "#" {
					continue
				} else {
					for j := next.y + move.y*i; j != next.y; j -= move.y {
						grid[next.x][j], grid[next.x][j-move.y] = grid[next.x][j-move.y], grid[next.x][j]
					}
					grid[next.x][next.y] = "@"
					grid[robot.x][robot.y] = "."
					robot = next
				}
			} else {
				// If we are moving in the y direction we need to determine all of the
				// affected coordinates from push on the box in the move direction.
				// The coordinates branch out from the source box depending on how much overlap
				// the source box has with the box above it. For example. If we we are moving ^
				//  ......    ..[]..
				//  ..[].. -> ..[]..
				//  ..[]..    ...@..
				//  ...@.     ......
				// Boxes in line like the one above dont branch outt the only move up.
				//
				// Boxes not in line with one and other branch out in the direction of the other
				// half of the box being pushed. For example:
				//  ......    ...[].
				//  ...[].    ..[]..
				//  ..[].. -> ..[]..
				//  ..[]..	  ...@..
				//  ...@..	  ......
				// This can get complicated when their are many boxes branching out
				affectedCords := make([]Pair, 0)
				// starting affected cords is the current position of the robot
				affectedCords = append(affectedCords, robot)
				// next affected cords starts as the current position of the robot
				nextAffectedCords := affectedCords
				for {
					// calculate the next cordinate affect by the push passing the last
					// iteration of affected cords and the move direction
					nextAffectedCords = getNextAffectCords(nextAffectedCords, move, grid)
					// Sort the cords by y value so we can swap them one row at a time
					// They will be unordered because a map is used in getNextAffectCords
					// to prevent duplicates
					sort.Slice(nextAffectedCords, func(i, j int) bool {
						return nextAffectedCords[i].x < nextAffectedCords[j].x
					})

					// Set flags to intial values
					allClear := true
					wallHit := false
					for _, cord := range nextAffectedCords {
						// If one of the next affected cords is a wall, we set wallHit to true
						if grid[cord.x][cord.y] == "#" && grid[cord.x-move.x][cord.y] != "." {
							wallHit = true
							// if one of the next affected cords is a box, we set allClear to false
							// because we will still need to calculate the effect this box has on the
							// next row of boxes
						} else if grid[cord.x][cord.y] == "]" || grid[cord.x][cord.y] == "[" {
							allClear = false
						}
					}

					// If we hit a wall, we break out of the loop and not move
					if wallHit {
						break
					}
					// If the set of next affected cords is all clear,
					// We DONT add the current set to affectedCords and
					// swap all of the cords in affectedCords with their
					// corresponding cords move 1 step in the move direction
					// Finally we swap the robot with the next position
					if allClear {
						for k := len(affectedCords) - 1; k >= 0; k-- {
							cord := affectedCords[k]
							if grid[cord.x][cord.y] == "]" || grid[cord.x][cord.y] == "[" {
								grid[cord.x][cord.y], grid[cord.x+move.x][cord.y] = grid[cord.x+move.x][cord.y], grid[cord.x][cord.y]
							}
						}
						grid[next.x][next.y] = "@"
						grid[robot.x][robot.y] = "."
						robot = next
						break
					} else {
						// If we didnt hit a wall and the next affected cords are not all clear
						// We need to continue to the next row of boxes. Add the current set of
						// nextAffectedCords to affectedCords and continue the loop
						affectedCords = append(affectedCords, nextAffectedCords...)
						continue
					}
				}
			}
		} else {
			// If the next position is open space, we move the robot to it
			grid[next.x][next.y] = "@"
			grid[robot.x][robot.y] = "."
			robot = next
		}
	}
	printGrid15(grid, robot)
	sum := 0
	for i, row := range grid {
		for j, letter := range row {
			if letter == "[" {
				sum += i*100 + j
			}
		}
	}

	fmt.Printf("Got sum: %d\n", sum)
}

func getNextAffectCords(cords []Pair, Xdirection Pair, grid [][]string) []Pair {
	// Keep track of affected coords in a map to prevent duplicates
	nextAffectedCords := make(map[Pair]bool, 0)
	for _, cord := range cords {
		// get the letter at the current cord
		kind := grid[cord.x][cord.y]
		if kind == "@" {
			// if the current cord is the robot, check which side if the box
			// the robot is adjacent to and add the adjacent cord and its
			// other side to the next affected cords
			boxSide := grid[cord.x+Xdirection.x][cord.y]
			checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y})
			if boxSide == "[" {
				checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y + 1})
			} else if boxSide == "]" {
				checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y - 1})
			}
		}
		// If the current cord is the left side of a box, add  the coord one step
		// in the move direction to the next affected cords. If the added cord is
		// the right side of a box, add the left side as well.
		if kind == "[" {
			checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y})
			if grid[cord.x+Xdirection.x][cord.y] == "]" {
				checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y - 1})
			}
		} else if kind == "]" {
			// If the current cord is the right side of a box, add  the coord one step
			// in the move direction to the next affected cords. If the added cord is
			// the left side of a box, add the right side as well.
			checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y})
			if grid[cord.x+Xdirection.x][cord.y] == "[" {
				checkedAddToMap(nextAffectedCords, Pair{cord.x + Xdirection.x, cord.y + 1})
			}
		}
	}
	return getKeys(nextAffectedCords)
}

func checkedAddToMap(m map[Pair]bool, p Pair) {
	if _, ok := m[p]; !ok {
		m[p] = true
	}
}

func getKeys(m map[Pair]bool) []Pair {
	keys := make([]Pair, 0)
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Day 16

type State struct {
	score     int
	current   Pair
	direction Pair
	seen      map[Pair]bool
}

func day16Part1(inputFile string) int {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	start := Pair{0, 0}
	end := Pair{0, 0}
	grid := make([][]string, 0)
	for i, line := range lines {
		row := make([]string, 0)
		for j, letter := range strings.Split(line, "") {
			if letter == "S" {
				start = Pair{i, j}
				row = append(row, ".")
			} else if letter == "E" {
				end = Pair{i, j}
				row = append(row, ".")
			} else {
				row = append(row, letter)
			}
		}
		grid = append(grid, row)
	}

	minCost := math.MaxInt32
	states := make([]State, 0)
	seen1 := make(map[Pair]bool, 0)
	states = append(states, State{0, start, Pair{0, 1}, seen1})
	seen2 := make(map[Pair]bool, 0)
	states = append(states, State{1000, start, Pair{-1, 0}, seen2})
	costs := make(map[Pair]int, 0)
	for len(states) > 0 {
		newStates := make([]State, 0)
		toDisplay := make([]Pair, 0)
		for _, state := range states {
			toDisplay = append(toDisplay, state.current)
		}
		for _, state := range states {
			printGrid16(grid, state.current, end)
			for _, dir := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Pair{state.current.x + dir.x, state.current.y + dir.y}
				var newCost int
				if dir == state.direction {
					newCost = state.score + 1
				} else {
					newCost = state.score + 1000 + 1
				}
				if _, ok := state.seen[next]; ok || isOutOfBounds(next, len(grid), len(grid[0])) {
					continue
				}
				if grid[next.x][next.y] == "#" {
					continue
				}
				if next == end {
					if newCost < minCost {
						minCost = newCost
					}
				}

				if _, ok := costs[next]; ok {
					if newCost >= costs[next] {
						continue
					} else {
						costs[next] = newCost
					}
				} else {
					costs[next] = newCost
				}
				newSeen := make(map[Pair]bool, 0)
				for k, v := range state.seen {
					newSeen[k] = v
				}
				newSeen[next] = true
				newStates = append(newStates, State{newCost, next, dir, newSeen})
			}
		}
		states = newStates
	}

	fmt.Printf("Got min cost: %d\n", minCost)

	return minCost
}

func printGrid16(grid [][]string, start Pair, end Pair) {
	for i, row := range grid {
		for j, letter := range row {
			if i == start.x && j == start.y {
				fmt.Printf("S")
			} else if i == end.x && j == end.y {
				fmt.Printf("E")
			} else {
				fmt.Printf("%s", letter)
			}
		}
		fmt.Printf("\n")
	}
}
func day16Part2Bad(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	start := Pair{0, 0}
	end := Pair{0, 0}
	grid := make([][]string, 0)
	for i, line := range lines {
		row := make([]string, 0)
		for j, letter := range strings.Split(line, "") {
			if letter == "S" {
				start = Pair{i, j}
				row = append(row, ".")
			} else if letter == "E" {
				end = Pair{i, j}
				row = append(row, ".")
			} else {
				row = append(row, letter)
			}
		}
		grid = append(grid, row)
	}

	minCost := day16Part1(inputFile)
	costs := make(map[Pair]int, 0)
	costs[end] = 0
	targets := make([]Pair, 0)
	for i := range grid {
		for j := range grid[i] {
			targets = append(targets, Pair{i, j})
		}
	}

	// Do stupid parallelization of procedure in part1
	// for every point in the grid so that we have its
	// shortest distance to the end point
	done := make(chan bool, 0)
	runningCount := 0
	costMutex := sync.Mutex{}
	runWaitGroup := sync.WaitGroup{}
	for len(targets) > 0 {
		if runningCount == 64 {
			<-done
			runningCount--
		}
		target := targets[0]
		targets = targets[1:]
		go getShortPathToPoint(grid, start, end, minCost, costs, target, &costMutex, &runWaitGroup, done)
		runningCount++
	}
	runWaitGroup.Wait()

	// Now that we have the shortest distance to the end point
	// for each point we can find all the shortest path from start to end
	minPaths := make([]State, 0)
	states := make([]State, 0)
	seen1 := make(map[Pair]bool, 0)
	seen1[start] = true
	states = append(states, State{0, start, Pair{0, 1}, seen1})
	seen2 := make(map[Pair]bool, 0)
	seen2[start] = true
	states = append(states, State{0, start, Pair{-1, 0}, seen2})
	for len(states) > 0 {
		newStates := make([]State, 0)
		for _, state := range states {
			for _, dir := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Pair{state.current.x + dir.x, state.current.y + dir.y}
				var newCost int
				if dir == state.direction {
					newCost = state.score + 1
				} else {
					newCost = state.score + 1000 + 1
				}
				if _, ok := state.seen[next]; ok || isOutOfBounds(next, len(grid), len(grid[0])) {
					continue
				}
				if grid[next.x][next.y] == "#" {
					continue
				}

				var cost int
				if _, ok := costs[next]; ok {
					cost = costs[next]
				} else {
					continue
				}
				if state.score+cost > minCost {
					continue
				}

				if next != end {
					newSeen := make(map[Pair]bool, 0)
					for k, v := range state.seen {
						newSeen[k] = v
					}
					newSeen[next] = true
					newStates = append(newStates, State{newCost, next, dir, newSeen})
				} else {
					if newCost == minCost {
						newSeen := make(map[Pair]bool, 0)
						for k, v := range state.seen {
							newSeen[k] = v
						}
						newSeen[next] = true
						minPaths = append(minPaths, State{newCost, next, dir, newSeen})
					}
				}
			}
		}
		states = newStates
	}

	fmt.Printf("Got min paths: %d\n", len(minPaths))

	allNodes := make(map[Pair]bool, 0)
	for _, path := range minPaths {
		for node := range path.seen {
			allNodes[node] = true
		}
	}

	printGrid16Tiles(grid, start, end, allNodes)

	fmt.Printf("Got all nodes: %d\n", len(allNodes))

}

func printGrid16Tiles(grid [][]string, start Pair, end Pair, tiles map[Pair]bool) {
	for i, row := range grid {
		for j, letter := range row {
			if _, ok := tiles[Pair{i, j}]; ok {
				fmt.Printf("O")
			} else if i == start.x && j == start.y {
				fmt.Printf("S")
			} else if i == end.x && j == end.y {
				fmt.Printf("E")
			} else {
				fmt.Printf("%s", letter)
			}
		}
		fmt.Printf("\n")
	}
}

func getShortPathToPoint(
	grid [][]string,
	start Pair,
	end Pair,
	minCost int,
	costs map[Pair]int,
	target Pair,
	costMutex *sync.Mutex,
	runWaitGroup *sync.WaitGroup,
	done chan bool,
) {
	runWaitGroup.Add(1)
	fmt.Printf("target: %v\n", target)
	states := make([]State, 0)
	seen1 := make(map[Pair]bool, 0)
	states = append(states, State{0, end, Pair{0, -1}, seen1})
	seen2 := make(map[Pair]bool, 0)
	states = append(states, State{0, end, Pair{1, 0}, seen2})
	for len(states) > 0 {
		newStates := make([]State, 0)
		for _, state := range states {
			for _, dir := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				next := Pair{state.current.x + dir.x, state.current.y + dir.y}
				var newCost int
				if dir == state.direction {
					newCost = state.score + 1
				} else {
					newCost = state.score + 1000 + 1
				}
				if newCost > minCost {
				}

				costMutex.Lock()
				if _, ok := state.seen[next]; ok || isOutOfBounds(next, len(grid), len(grid[0])) {
					costMutex.Unlock()
					continue
				}
				if grid[next.x][next.y] == "#" {
					costMutex.Unlock()
					continue
				}

				if _, ok := costs[next]; ok {
					if newCost > costs[next] {
						costMutex.Unlock()
						continue
					} else {
						costs[next] = newCost
					}
				} else {
					costs[next] = newCost
				}
				costMutex.Unlock()

				if next != target {
					newSeen := make(map[Pair]bool, 0)
					for k, v := range state.seen {
						newSeen[k] = v
					}
					newSeen[next] = true
					newState := State{newCost, next, dir, newSeen}
					newStates = append(newStates, newState)
				}
			}
		}
		states = newStates
	}
	runWaitGroup.Done()
	done <- true
}

type Iteration struct {
	current   Pair
	direction Pair
}

func day16Part2Good(inputFile string) {
	input := readInput(inputFile)
	trimmed := strings.Trim(input, "\n")
	lines := strings.Split(trimmed, "\n")
	start := Pair{0, 0}
	end := Pair{0, 0}
	grid := make([][]string, 0)
	for i, line := range lines {
		row := make([]string, 0)
		for j, letter := range strings.Split(line, "") {
			if letter == "S" {
				start = Pair{i, j}
				row = append(row, ".")
			} else if letter == "E" {
				end = Pair{i, j}
				row = append(row, ".")
			} else {
				row = append(row, letter)
			}
		}
		grid = append(grid, row)
	}

	queue := make([]Iteration, 0)
	distances := make(map[Pair]int, 0)
	for i := range grid {
		for j := range grid[i] {
			distances[Pair{i, j}] = math.MaxInt32
		}
	}
	distances[end] = 0
	startIter := Iteration{end, Pair{1, 0}}
	queue = append(queue, startIter)
	for len(queue) > 0 {
		iter := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for _, dir := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if dir.x == -iter.direction.x && dir.y == -iter.direction.y {
				continue
			}

			next := Pair{iter.current.x + dir.x, iter.current.y + dir.y}
			if isOutOfBounds(next, len(grid), len(grid[0])) {
				continue
			}
			if grid[next.x][next.y] == "#" {
				continue
			}

			cost := 1
			if dir != iter.direction {
				cost = 1000 + 1
			}

			if distances[next] > distances[iter.current]+cost {
				distances[next] = distances[iter.current] + cost
				queue = append(queue, Iteration{next, dir})
			}
		}

		sort.Slice(queue, func(i, j int) bool {
			return distances[queue[i].current] > distances[queue[j].current]
		})
	}

	minCost := distances[start]
	fmt.Printf("Got min cost: %d\n", minCost)

	// Now that we have the shortest distance to the end point
	// for each point we can find all the shortest path from start to end

	minPaths := make([]State, 0)
	states := make([]State, 0)

	seen1 := make(map[Pair]bool, 0)
	seen1[start] = true
	states = append(states, State{0, start, Pair{0, 1}, seen1})

	seen2 := make(map[Pair]bool, 0)
	seen2[start] = true
	states = append(states, State{0, start, Pair{-1, 0}, seen2})
	for len(states) > 0 {
		newStates := make([]State, 0)
		for _, state := range states {
			for _, dir := range []Pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				if dir.x == -state.direction.x && dir.y == -state.direction.y {
					continue
				}
				next := Pair{state.current.x + dir.x, state.current.y + dir.y}

				var newCost int
				if dir == state.direction {
					newCost = state.score + 1
				} else {
					newCost = state.score + 1000 + 1
				}

				if _, ok := state.seen[next]; ok || isOutOfBounds(next, len(grid), len(grid[0])) {
					continue
				}

				if grid[next.x][next.y] == "#" {
					continue
				}

				if state.score+distances[next] > minCost {
					continue
				}

				if next != end {
					newSeen := make(map[Pair]bool, 0)
					for k, v := range state.seen {
						newSeen[k] = v
					}
					newSeen[next] = true
					newStates = append(newStates, State{newCost, next, dir, newSeen})
				} else {
					if newCost <= minCost {
						newSeen := make(map[Pair]bool, 0)
						for k, v := range state.seen {
							newSeen[k] = v
						}
						newSeen[next] = true
						minPaths = append(minPaths, State{newCost, next, dir, newSeen})
					}
				}
			}
		}
		states = newStates
	}

	allNodes := make(map[Pair]bool, 0)
	for _, path := range minPaths {
		for node := range path.seen {
			allNodes[node] = true
		}
	}

	printGrid16Tiles(grid, start, end, allNodes)
	fmt.Printf("Got all nodes: %d\n", len(allNodes))
}
