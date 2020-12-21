package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"sort"
	"strings"
)

func main() {
	Day21()
}

func Day21() {
	fmt.Println("==============  Day21 =============")
	fmt.Println("==============  TEST  =============")
	output, canonicalTestNames := countNonAllergentIngredients(filereader.ReadFile("./input/day-21/test.txt"))
	if output == 5 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	if canonicalTestNames == "mxmxvkd,sqjhc,fvjkl" {
		fmt.Println("Test succeeded:", canonicalTestNames)
	} else {
		fmt.Println("Test failed:", canonicalTestNames)
	}

	fmt.Println("============== OUTPUT =============")
	count, canonicalNames := countNonAllergentIngredients(filereader.ReadFile("./input/day-21/input.txt"))
	fmt.Println("Ingredients without allergens:", count)
	fmt.Println("Ingredients without allergens canonical:", canonicalNames)
}

type Food struct {
	ingredients []string
	allergens   []string
}

func countNonAllergentIngredients(foodlist []string) (int, string) {
	allergens := make(map[string][]string)
	foods := []Food{}

	ingredientsCount := map[string]int{}

	for _, food := range foodlist {
		ingredients := strings.Split(food[0:strings.Index(food, "(contains")-1], " ")
		allergenList := strings.Split(food[strings.Index(food, "(contains ")+len("(contains "):len(food)-1], ", ")
		foods = append(foods, Food{ingredients: ingredients, allergens: allergenList})

		for _, ingredient := range ingredients {
			ingredientsCount[ingredient]++
		}

		for _, allergen := range allergenList {
			if _, exists := allergens[allergen]; !exists {
				allergens[allergen] = ingredients
			} else {
				allergens[allergen] = innerJoin(allergens[allergen], ingredients)
			}
		}
	}

	oneStillMultiple := true

	for oneStillMultiple {
		oneStillMultiple = false

		for allergen, possible := range allergens {
			if len(possible) != 1 {
				oneStillMultiple = true
			} else {
				for otherAllergen, otherIngredients := range allergens {
					if otherAllergen != allergen {
						allergens[otherAllergen] = removeFromSlice(otherIngredients, possible[0])
					}
				}
			}
		}
	}

	for _, ingredients := range allergens {
		delete(ingredientsCount, ingredients[0])
	}

	countAllergieFreeIngredient := 0
	for _, allergieFreeIngredientCount := range ingredientsCount {
		countAllergieFreeIngredient += allergieFreeIngredientCount
	}

	var allergenNames []string
	for allergenName := range allergens {
		allergenNames = append(allergenNames, allergenName)
	}
	sort.Strings(allergenNames)

	var ingredientNames []string
	for _, allergenName := range allergenNames {
		ingredientNames = append(ingredientNames, allergens[allergenName][0])
	}

	return countAllergieFreeIngredient, strings.Join(ingredientNames, ",")
}

func innerJoin(set1, set2 []string) []string {
	map1 := map[string]bool{}
	for _, v := range set1 {
		map1[v] = true
	}

	var unionSet []string
	for _, v := range set2 {
		if map1[v] {
			unionSet = append(unionSet, v)
		}
	}
	return unionSet
}

func removeFromSlice(sli []string, strToFind string) []string {
	for i, v := range sli {
		if v == strToFind {
			sli[i] = sli[len(sli)-1]
			sli = sli[:len(sli)-1]
		}
	}
	return sli
}
