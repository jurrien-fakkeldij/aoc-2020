package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
)

func main() {
	Day1()
}

func Day1() {
	fmt.Println("============== Day1 ==============")
	Part1()
	Part2()
}

func Part1() {
	fmt.Println("============== Part1 =============")
	testExpenseReport, _ := transformer.SliceAtoi(filereader.ReadFile("./input/day-1/test.txt"))
	solution := fixExpenseReport(testExpenseReport)

	if solution != 514579 {
		fmt.Println("Test went wrong:", solution)
	} else {
		fmt.Println("Test passed:", solution)
	}

	expenseReport, _ := transformer.SliceAtoi(filereader.ReadFile("./input/day-1/input.txt"))
	solution = fixExpenseReport(expenseReport)
	fmt.Println("Solution:", solution)
}

func fixExpenseReport(expenses []int) int {
	correctExpense := 0
	for expensePointer, expense := range expenses {
		for _, anotherExpense := range expenses[expensePointer+1:] {
			if expense+anotherExpense == 2020 {
				correctExpense = expense * anotherExpense
				break
			}
		}
	}
	return correctExpense
}

func Part2() {
	fmt.Println("============== Part2 =============")
	testExpenseReport, _ := transformer.SliceAtoi(filereader.ReadFile("./input/day-1/test.txt"))
	solution := fixExpenseReportPart2(testExpenseReport)

	if solution != 241861950 {
		fmt.Println("Test went wrong:", solution)
	} else {
		fmt.Println("Test passed:", solution)
	}

	expenseReport, _ := transformer.SliceAtoi(filereader.ReadFile("./input/day-1/input.txt"))
	solution = fixExpenseReportPart2(expenseReport)
	fmt.Println("Solution:", solution)
}

func fixExpenseReportPart2(expenses []int) int {
	correctExpense := 0
	for expensePointer, expense := range expenses {
		for secondExpensePointer, anotherExpense := range expenses[expensePointer+1:] {
			for _, thirdExpense := range expenses[secondExpensePointer+1:] {
				if expense+anotherExpense+thirdExpense == 2020 {
					correctExpense = expense * anotherExpense * thirdExpense
					break
				}
			}
		}
	}
	return correctExpense
}
