package main

import (
	"aoc2015/util"
	"fmt"
)

func part1(data string) int {
	x, y := 0, 0

	deliveries := map[string]int{
		"0x0": 1,
	}

	for _, instruction := range data {
		switch instruction {
		case '^':
			y += 1
		case 'v':
			y -= 1
		case '<':
			x -= 1
		case '>':
			x += 1
		default:
			panic(instruction)
		}

		deliveries[fmt.Sprintf("%dx%d", x, y)] += 1
	}

	return len(deliveries)
}

func part2(data string) int {
	x1, y1 := 0, 0
	x2, y2 := 0, 0

	deliveries := map[string]int{
		"0x0": 1,
	}

	for index, instruction := range data {
		if index%2 == 0 {
			switch instruction {
			case '^':
				y1 += 1
			case 'v':
				y1 -= 1
			case '<':
				x1 -= 1
			case '>':
				x1 += 1
			default:
				panic(instruction)
			}

			deliveries[fmt.Sprintf("%dx%d", x1, y1)] += 1
		} else {
			switch instruction {
			case '^':
				y2 += 1
			case 'v':
				y2 -= 1
			case '<':
				x2 -= 1
			case '>':
				x2 += 1
			default:
				panic(instruction)
			}

			deliveries[fmt.Sprintf("%dx%d", x2, y2)] += 1
		}
	}

	return len(deliveries)
}

func main() {
	testData := util.ReadFile("./d03/test1.txt")
	data := util.ReadFile("./d03/input.txt")

	util.TestRun1(part1(testData))
	util.TestRun2(part2(testData))

	util.SolutionPart1(part1(data))
	util.SolutionPart2(part2(data))
}
