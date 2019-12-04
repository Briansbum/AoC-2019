package main

import (
	"log"
)

var iterations int

var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 10, 23, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 35, 9, 39, 2, 39, 13, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 9, 59, 63, 2, 6, 63, 67, 1, 13, 67, 71, 1, 9, 71, 75, 2, 13, 75, 79, 1, 79, 10, 83, 2, 83, 9, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 99, 5, 103, 1, 103, 2, 107, 1, 107, 10, 0, 99, 2, 0, 14, 0}

func main() {
	// part 1
	inputCopy := input
	inputCopy[1] = 12
	inputCopy[2] = 2
	log.Println(performOperation(inputCopy))

	// part 2
	inputCopy = input
	log.Println(findParams(inputCopy, 19690720))
}

func performOperation(program []int) []int {
	for cur := 0; cur <= len(program); cur++ {
		log.Println(cur)
		log.Println(cur + 4)
		log.Println(len(program))
		if cur+4 > len(program) {
			break
		}
		switch program[cur] {
		case 1:
			program[program[cur+3]] = program[program[cur+1]] + program[program[cur+2]]
			cur = cur + 3
		case 2:
			program[program[cur+3]] = program[program[cur+1]] * program[program[cur+2]]
			cur = cur + 3
		case 99:
			return program
		}
	}
	return program
}

func findParams(program []int, desiredResult int) (int, int) {
	p1, p2 := 0, 0
	for p1 < 100 {
		for p2 < 100 {
			programCopy := make([]int, len(program))
			copy(programCopy, program)

			programCopy[1] = p1
			programCopy[2] = p2

			if performOperation(programCopy)[0] == desiredResult {
				return p1, p2
			}

			p2++
		}
		p1++
	}
	return 0, 0
}
