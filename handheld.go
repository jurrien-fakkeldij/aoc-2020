package main

import (
	"fmt"
	"strconv"
	"strings"
)

type HandHeld struct {
	bootcode []string
}

func (handheld *HandHeld) DiagBootCode(withDebug bool) (int, int) {
	runMap := make(map[int]int)
	accumulator := 0
	index := 0

	instructions := make(chan string, 1)
	instructions <- handheld.bootcode[index]

	output := -1

	for instruction := range instructions {
		if runMap[index] != 0 {
			//diag done
			//fmt.Println("Diagnosis is done.")
			close(instructions)
		} else {
			runMap[index]++
			if withDebug {
				fmt.Println("handleInstruction:", instruction, "index:", index)
			}
			accumulatorAddition, indexTraversal := handleInstruction(instruction)
			index += indexTraversal
			accumulator += accumulatorAddition
			if index == len(handheld.bootcode) {
				close(instructions)
				output = 0
			} else {
				instructions <- handheld.bootcode[index]
			}
		}
	}
	return accumulator, output
}

func handleInstruction(instruction string) (int, int) {
	operation := strings.Split(instruction, " ")[0]
	argument := strings.Split(instruction, " ")[1]
	argumentInteger, _ := strconv.Atoi(argument)

	accumulatorAddition := 0
	indexTraversal := 0

	//fmt.Println("operation:", operation, "argument:", argumentInteger)

	switch operation {
	case "acc":
		accumulatorAddition = argumentInteger
		indexTraversal = 1
	case "jmp":
		accumulatorAddition = 0
		indexTraversal = argumentInteger
	case "nop":
		accumulatorAddition = 0
		indexTraversal = 1
	}

	return accumulatorAddition, indexTraversal
}
