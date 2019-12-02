package main

import "log"

var iterations int

var input = []int{1, 12, 2, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 19, 5, 23, 2, 10, 23, 27, 2, 27, 13, 31, 1, 10, 31, 35, 1, 35, 9, 39, 2, 39, 13, 43, 1, 43, 5, 47, 1, 47, 6, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 9, 59, 63, 2, 6, 63, 67, 1, 13, 67, 71, 1, 9, 71, 75, 2, 13, 75, 79, 1, 79, 10, 83, 2, 83, 9, 87, 1, 5, 87, 91, 2, 91, 6, 95, 2, 13, 95, 99, 1, 99, 5, 103, 1, 103, 2, 107, 1, 107, 10, 0, 99, 2, 0, 14, 0}

func main() {
	res := performOperation(&input, 0)
	log.Println(res)
	log.Println(iterations)
}

func performOperation(program *[]int, offset int) *[]int {
	iterations++
	loc1 := (*program)[offset+1]
	loc2 := (*program)[offset+2]
	resTarget := (*program)[offset+3]
	log.Printf("started with resTarget == %v", (*program)[resTarget])
	if (*program)[offset] == 1 {
		log.Println("add")
		(*program)[resTarget] = (*program)[loc1] + (*program)[loc2]
		log.Printf("result of %v when adding %v by %v", (*program)[resTarget], (*program)[loc1], (*program)[loc2])
	} else if (*program)[offset] == 2 {
		log.Println("multiply")
		(*program)[resTarget] = (*program)[loc1] * (*program)[loc2]
		log.Printf("result of %v when multiplying %v by %v", (*program)[resTarget], (*program)[loc1], (*program)[loc2])
	}
	log.Printf("ended with resTarget == %v", (*program)[resTarget])
	// this is the last thing this function does, always
	nextOffset := searchForOpcode(input, offset)
	if nextOffset == -1 {
		log.Println("HALT")
		return program
	}
	return performOperation(program, nextOffset)
}

func searchForOpcode(program []int, currOpcodeOffset int) int {
	for i := currOpcodeOffset; i <= currOpcodeOffset+4; i++ {
		if program[i] == 99 {
			return -1
		}
	}
	return currOpcodeOffset + 4
}
