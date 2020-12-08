package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strings"
)

func main() {
	Day8()
}

func Day8() {
	fmt.Println("============== Day8 ==============")
	fmt.Println("============== TEST ==============")

	testBootCode := filereader.ReadFile("./input/day-8/test-bootcode.txt")
	bootCode := filereader.ReadFile("./input/day-8/bootcode.txt")
	handHeld := new(HandHeld)
	handHeld.bootcode = testBootCode
	output, _ := handHeld.DiagBootCode(false)

	if output == 5 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = handHeld.fixBootCode()

	if output == 8 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT ============")
	handHeld.bootcode = bootCode
	output, _ = handHeld.DiagBootCode(false)
	fmt.Println("Diag output:", output)
	fmt.Println("Run output:", handHeld.fixBootCode())
}

func (handHeld *HandHeld) fixBootCode() int {
	outputValue := 0
	output := -1
	outputCode := -1
	initialBootcode := handHeld.bootcode

	//fmt.Println("initial bootcode:", initialBootcode)
	fixedBootCode := make([]string, len(initialBootcode))

	for index, instruction := range initialBootcode {
		//fmt.Println("bootcode:", initialBootcode)
		copy(fixedBootCode, initialBootcode)

		if strings.HasPrefix(instruction, "nop") {
			//fmt.Println("fixed bootcode:", fixedBootCode)
			fixedBootCode[index] = strings.Replace(instruction, "nop", "jmp", 1)
			handHeld.bootcode = fixedBootCode
			output, outputCode = handHeld.DiagBootCode(false)
			if outputCode == 0 {
				output, outputCode = handHeld.DiagBootCode(true)
				outputValue = output
			}
		} else if strings.HasPrefix(instruction, "jmp") {
			//fmt.Println("fixed bootcode:", fixedBootCode)
			fixedBootCode[index] = strings.Replace(instruction, "jmp", "nop", 1)
			handHeld.bootcode = fixedBootCode
			output, outputCode = handHeld.DiagBootCode(false)
			if outputCode == 0 {
				output, outputCode = handHeld.DiagBootCode(true)
				outputValue = output
			}
		}
	}
	return outputValue
}
