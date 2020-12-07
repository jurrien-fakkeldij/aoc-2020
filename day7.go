package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strconv"
	"strings"
)

type recipe struct {
	name        string
	ingredients map[string]int
}

func main() {
	Day7()
}

func Day7() {
	fmt.Println("============== Day7 ==============")
	fmt.Println("============== TEST ==============")
	testRecipes := filereader.ReadFile("./input/day-7/test.txt")
	testRecipes2 := filereader.ReadFile("./input/day-7/test-2.txt")
	actualRecipes := filereader.ReadFile("./input/day-7/input.txt")

	recipes := createRecipes(testRecipes)
	recipes2 := createRecipes(testRecipes2)
	for _, recipe := range recipes {
		fmt.Printf("%+v \n", recipe)
	}

	uniqueBags := findUniqueBagsForBag("shiny gold", recipes)
	amount := len(uniqueBags)
	if amount == 4 {
		fmt.Println("Test succeeded:", amount)
	} else {
		fmt.Println("Test failed:", amount)
	}

	amountOfBagsNeeded := countAmountOfBagsNeeded("shiny gold", recipes)
	if amountOfBagsNeeded == 32 {
		fmt.Println("Test succeeded:", amountOfBagsNeeded)
	} else {
		fmt.Println("Test failed:", amountOfBagsNeeded)
	}

	amountOfBagsNeeded = countAmountOfBagsNeeded("shiny gold", recipes2)
	if amountOfBagsNeeded == 126 {
		fmt.Println("Test succeeded:", amountOfBagsNeeded)
	} else {
		fmt.Println("Test failed:", amountOfBagsNeeded)
	}

	fmt.Println("============== OUTPUT ============")
	recipes = createRecipes(actualRecipes)
	uniqueBags = findUniqueBagsForBag("shiny gold", recipes)
	amount = len(uniqueBags)
	fmt.Println("Unique bags to be used:", amount)

	amountOfBagsNeeded = countAmountOfBagsNeeded("shiny gold", recipes)
	fmt.Println("Bags to be used:", amountOfBagsNeeded)
}

func countAmountOfBagsNeeded(recipeName string, recipes map[string]*recipe) int {
	recipe := recipes[recipeName]
	count := 0
	for ingredientName, ingredientAmount := range recipe.ingredients {
		count += countAmountOfBagsNeeded(ingredientName, recipes) * ingredientAmount
		count += ingredientAmount
	}
	return count
}

func findUniqueBagsForBag(recipeName string, recipes map[string]*recipe) []string {
	recipesList := []string{}

	for _, recipe := range recipes {
		_, present := recipe.ingredients[recipeName]
		if present {
			recipesList = append(recipesList, findUniqueBagsForBag(recipe.name, recipes)...)
			recipesList = append(recipesList, []string{recipe.name}...)
		}
	}

	//fmt.Println("recipeName:", recipeName, " list: ", recipesList)

	return uniqueStrings(recipesList)
}

func createRecipes(recipesInput []string) map[string]*recipe {
	recipes := make(map[string]*recipe)
	for _, recipeInput := range recipesInput {
		recipe := createRecipe(recipeInput)
		recipes[recipe.name] = recipe

	}

	return recipes
}

func uniqueStrings(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func createRecipe(recipeInput string) *recipe {
	recipe := new(recipe)
	recipe.ingredients = make(map[string]int)
	recipe.name = string(recipeInput[0:strings.Index(recipeInput, " bags")])

	ingredientsInput := strings.Split(recipeInput, " contain ")[1]
	if ingredientsInput == "no other bags." {
		//noop
	} else {
		for _, ingredientInput := range strings.Split(ingredientsInput, ", ") {
			name, amount := getIngredient(ingredientInput)
			recipe.ingredients[name] = amount
		}
	}
	return recipe
}

func getIngredient(ingredientInput string) (string, int) {
	amount, _error := strconv.Atoi(string([]rune(ingredientInput)[0]))
	if _error != nil {
		fmt.Println("Could not parse amount:", _error, "for:", ingredientInput)
	}

	name := ingredientInput[2:strings.Index(ingredientInput, " bag")]

	return name, amount
}
