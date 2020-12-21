package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Day18()
}

func Day18() {
	fmt.Println("==============  Day18 =============")
	fmt.Println("==============  TEST  =============")
	testProblem1 := "2 * 3 + (4 * 5)"
	output := EvaluateAdvanced(testProblem1)
	if output == 46 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testProblem2 := "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	output = EvaluateAdvanced(testProblem2)
	if output == 1445 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testProblem3 := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	output = EvaluateAdvanced(testProblem3)
	if output == 669060 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testProblem4 := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	output = EvaluateAdvanced(testProblem4)
	if output == 23340 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}
	fmt.Println("============== OUTPUT =============")
	homework := filereader.ReadFile("./input/day-18/input.txt")
	output = 0
	for _, problem := range homework {
		output = output + EvaluateAdvanced(problem)
	}
	fmt.Println("Sum: ", output)
}

func Evaluate(expression string) int {
	idxs := findBracketIdxs(expression)

	if len(idxs) == 0 {
		o := 0

		op := ""
		for _, s := range strings.Split(expression, " ") {
			if s == "+" || s == "*" {
				op = s
			} else {
				i, _ := strconv.Atoi(s)
				if op == "+" {
					o += i
				} else if op == "*" {
					o *= i
				} else {
					o = i
				}
				op = ""
			}
		}

		return o
	}

	subExpression := expression[idxs[0]+1 : idxs[1]]
	nextExpression := expression[:idxs[0]] + strconv.Itoa(Evaluate(subExpression)) + expression[idxs[1]+1:]
	return Evaluate(nextExpression)
}

func EvaluateAdvanced(expression string) int {
	idxs := findBracketIdxs(expression)

	if len(idxs) == 0 {
		idxs := findPlusIdxs(expression)

		if len(idxs) == 0 {
			o := 1
			for _, s := range strings.Split(expression, " ") {
				if s == "*" {
					continue
				}
				i, _ := strconv.Atoi(s)
				o *= i
			}
			return o
		}

		subExpression := expression[idxs[0] : idxs[1]+1]
		nextExpression := expression[:idxs[0]] + strconv.Itoa(Evaluate(subExpression)) + expression[idxs[1]+1:]
		return EvaluateAdvanced(nextExpression)
	}

	subExpression := expression[idxs[0]+1 : idxs[1]]
	nextExpression := expression[:idxs[0]] + strconv.Itoa(EvaluateAdvanced(subExpression)) + expression[idxs[1]+1:]
	return EvaluateAdvanced(nextExpression)
}

func findBracketIdxs(expression string) []int {
	o := []int{}
	c := 0
	for idx, char := range expression {
		if char == '(' {
			c++
			if len(o) == 0 {
				o = append(o, idx)
			}
		} else if char == ')' {
			c--
			if c == 0 {
				o = append(o, idx)
				break
			}
		}
	}
	return o
}

func findPlusIdxs(expression string) []int {
	plusIdx := -1
	for idx, char := range expression {
		if char == '+' {
			plusIdx = idx
			break
		}
	}
	if plusIdx < 0 {
		return []int{}
	}
	lower, upper := plusIdx-2, plusIdx+2
	for {
		if lower == 0 {
			break
		} else if expression[lower] == ' ' {
			lower++
			break
		} else {
			lower--
		}
	}
	for {
		if upper == len(expression)-1 {
			break
		} else if expression[upper] == ' ' {
			upper--
			break
		} else {
			upper++
		}
	}
	return []int{lower, upper}
}
