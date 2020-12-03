package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	lower     int
	upper     int
	character string
}

func main() {
	Day2()
}

func Day2() {
	fmt.Println("============== Day2 ==============")
	fmt.Println("============== Part1 =============")
	testPasswordFile := filereader.ReadFile("./input/day-2/test.txt")
	count := countCorrectPasswords(testPasswordFile)
	if count == 2 {
		fmt.Println("Test succeeded:", count)
	}

	count = countCorrectPasswords(filereader.ReadFile("./input/day-2/input.txt"))
	fmt.Println("Number of correct passwords:", count)

	fmt.Println("============== Part2 =============")
	count = countCorrectOTCPPasswords(testPasswordFile)
	if count == 1 {
		fmt.Println("Test succeeded:", count)
	}

	count = countCorrectOTCPPasswords(filereader.ReadFile("./input/day-2/input.txt"))
	fmt.Println("Number of correct passwords:", count)
}

func countCorrectPasswords(passwords []string) int {
	count := 0
	for _, passwordLine := range passwords {
		policy, password := getPasswordAndPolicyFromString(passwordLine)
		if strings.Count(password, policy.character) >= policy.lower &&
			strings.Count(password, policy.character) <= policy.upper {
			count++
		}
	}
	return count
}

func countCorrectOTCPPasswords(passwords []string) int {
	count := 0
	for _, passwordLine := range passwords {
		policy, password := getPasswordAndPolicyFromString(passwordLine)
		fmt.Printf("characters %c %c -> ", password[policy.lower], password[policy.upper])
		if (string(password[policy.lower]) == policy.character || string(password[policy.upper]) == policy.character) &&
			!(string(password[policy.lower]) == policy.character && string(password[policy.upper]) == policy.character) {
			count++
			fmt.Println("Correct:", passwordLine)
		} else {
			fmt.Println("Incorrect:", passwordLine)
		}
	}
	return count
}

func (policy *PasswordPolicy) fromString(policyString string) {
	policySlice := strings.Fields(policyString)
	policy.character = policySlice[1]
	lower, err := strconv.Atoi(strings.Split(policySlice[0], "-")[0])
	upper, err := strconv.Atoi(strings.Split(policySlice[0], "-")[1])
	if err != nil {
		fmt.Println("Error converting to integer:", err)
		return
	}
	policy.lower = lower
	policy.upper = upper
}

func getPasswordAndPolicyFromString(passwordLine string) (*PasswordPolicy, string) {
	splitPasswordPolicy := strings.Split(passwordLine, ":")
	policy := new(PasswordPolicy)
	policy.fromString(splitPasswordPolicy[0])
	password := splitPasswordPolicy[1]

	return policy, password
}
