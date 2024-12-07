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

	total := 0
	for cal, values := range calibrations {
		operations := len(values) - 1
		combinations := intPow(2, operations)
		for i := 0; i < combinations; i++ {
			currentTotal := values[0]
			for j := 0; j < operations; j++ {
				op := getOperation(i, j)
				if op == 1 {
					currentTotal += values[j+1]
				} else {
					currentTotal *= values[j+1]
				}
			}
			if currentTotal == cal {
				total += cal
				break
			}
		}
	}

	fmt.Printf("Got total: %d\n", total)
}

func intPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func product(values []int) int {
	total := 1
	for _, value := range values {
		total *= value
	}
	return total
}

func getOperation(i int, opIdx int) int {
	return (i >> opIdx) & 1
}

func day7Part2(inputFile string) {
	input := readInput(inputFile)
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

	total := 0
	for cal, values := range calibrations {
		var err error
		combinationsCount := intPow(3, len(values)-1)
		operations := len(values) - 1
		nextOp := nextOperation([]int{0, 1, 2}, operations)
		for i := 0; i < combinationsCount; i++ {
			ops := nextOp()
			currentTotal := values[0]
			for j, op := range ops {
				if op == 0 {
					currentTotal += values[j+1]
				} else if op == 1 {
					currentTotal *= values[j+1]
				} else {
					left := strconv.Itoa(currentTotal)
					right := strconv.Itoa(values[j+1])
					currentTotal, err = strconv.Atoi(strings.Join([]string{left, right}, ""))
					check(err)
				}
			}
			if currentTotal == cal {
				total += cal
				break
			}
		}
	}

	fmt.Printf("Got total: %d\n", total)
}

func nextOperation(operations []int, repeat int) func() []int {
	p := make([]int, repeat)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = operations[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(operations) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}
