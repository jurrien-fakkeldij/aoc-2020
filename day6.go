package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strings"
)

func main() {
	Day6()
}

func Day6() {
	fmt.Println("============== Day5 ==============")
	fmt.Println("============== TEST ==============")
	testDeclarationForm := filereader.ReadFile("./input/day-6/test.txt")
	count := countYesAnswers(testDeclarationForm)
	if count == 11 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}

	count = countYesAnwersByEverone(testDeclarationForm)
	if count == 6 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}

	fmt.Println("============== OUTPUT ============")
	actualDeclarationForm := filereader.ReadFile("./input/day-6/input.txt")
	count = countYesAnswers(actualDeclarationForm)
	fmt.Println("Counted unqiue answers within all groups:", count)

	count = countYesAnwersByEverone(actualDeclarationForm)
	fmt.Println("Counted only the yes for the entire group:", count)
}

func countYesAnwersByEverone(declarationForms []string) int {
	count := 0
	startIndex := 0
	for index, line := range declarationForms {
		if len(line) == 0 {
			count += countGroupAllYesAnswers(declarationForms[startIndex:index])
			startIndex = index + 1
		} else if len(declarationForms)-1 == index {
			count += countGroupAllYesAnswers(declarationForms[startIndex : index+1])
		}
	}
	return count
}

func countGroupAllYesAnswers(declarationForms []string) int {
	count := 0
	alphabet := []string{"a", "b", "c", "d", "e", "f", "h", "g", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	allAnswers := groupAnswers(declarationForms)
	for _, letter := range alphabet {
		if strings.Count(allAnswers, letter) == len(declarationForms) {
			count++
		}
	}
	return count
}

func countYesAnswers(declarationForms []string) int {
	count := 0
	startIndex := 0
	for index, line := range declarationForms {
		if len(line) == 0 {
			count += countGroupYesAnswers(declarationForms[startIndex:index])
			startIndex = index + 1
		} else if len(declarationForms)-1 == index {
			count += countGroupYesAnswers(declarationForms[startIndex : index+1])
		}
	}
	return count
}

func groupAnswers(declarationForm []string) string {
	var answers strings.Builder
	for _, line := range declarationForm {
		answers.WriteString(line)
	}
	return answers.String()
}

func countGroupYesAnswers(declartionForm []string) int {
	var answers strings.Builder
	for _, line := range declartionForm {
		answers.WriteString(line)
	}

	answerRunes := []rune(answers.String())
	uniqueAnswers := unique(answerRunes)
	return len(uniqueAnswers)
}

func unique(runeSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
