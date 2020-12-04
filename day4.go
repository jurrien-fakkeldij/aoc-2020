package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Day4()
}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func isValidBYR(byr string) bool {
	byrNumber, error := strconv.Atoi(byr)
	return len(byr) == 4 && error == nil && byrNumber >= 1920 && byrNumber <= 2002
}

func isValidIYR(iyr string) bool {
	iyrNumber, error := strconv.Atoi(iyr)
	return len(iyr) == 4 && error == nil && iyrNumber >= 2010 && iyrNumber <= 2020
}

func isValidEYR(eyr string) bool {
	eyrNumber, error := strconv.Atoi(eyr)
	return len(eyr) == 4 && error == nil && eyrNumber >= 2020 && eyrNumber <= 2030
}

func isValidHGT(hgt string) bool {
	if strings.Contains(hgt, "cm") {
		heightStr := hgt[:strings.Index(hgt, "cm")]
		height, _error := strconv.Atoi(heightStr)
		return _error == nil && height >= 150 && height <= 193
	} else if strings.Contains(hgt, "in") {
		heightStr := hgt[:strings.Index(hgt, "in")]
		height, _error := strconv.Atoi(heightStr)
		return _error == nil && height >= 59 && height <= 76
	} else {
		return false
	}
}

func isValidHCL(hcl string) bool {
	regex := regexp.MustCompile("^#([0-9a-f]{6})")
	return regex.MatchString(hcl)
}

func isValidECL(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}

func isValidPID(pid string) bool {
	return len(pid) == 9
}

func (passport *Passport) isValid(debug bool) bool {
	if debug {
		fmt.Print("valid byr: ", isValidBYR(passport.byr))
		fmt.Print(" valid iyr: ", isValidIYR(passport.iyr))
		fmt.Print(" valid eyr: ", isValidEYR(passport.eyr))
		fmt.Print(" valid hgt: ", isValidHGT(passport.hgt))
		fmt.Print(" valid hcl: ", isValidHCL(passport.hcl))
		fmt.Print(" valid ecl: ", isValidECL(passport.ecl))
		fmt.Print(" valid pid: ", isValidPID(passport.pid))
		fmt.Print("\n")
	}
	return isValidBYR(passport.byr) &&
		isValidIYR(passport.iyr) &&
		isValidEYR(passport.eyr) &&
		isValidHGT(passport.hgt) &&
		isValidHCL(passport.hcl) &&
		isValidECL(passport.ecl) &&
		isValidPID(passport.pid)
}

func fromStringArray(passportString []string) *Passport {
	passport := new(Passport)
	for _, line := range passportString {
		for _, property := range strings.Split(line, " ") {
			value := strings.Split(property, ":")[1]
			switch key := strings.Split(property, ":")[0]; key {
			case "byr":
				passport.byr = value
			case "iyr":
				passport.iyr = value
			case "eyr":
				passport.eyr = value
			case "hgt":
				passport.hgt = value
			case "hcl":
				passport.hcl = value
			case "ecl":
				passport.ecl = value
			case "pid":
				passport.pid = value
			case "cid":
				passport.cid = value
			default:
				fmt.Println("something went wrong key:", key, "value", value)
			}
		}
	}
	//fmt.Printf("passport: %+v, valid: %v\n", passport, passport.isValid(true))
	return passport
}

func Day4() {
	fmt.Println("============== Day4 ==============")
	fmt.Println("============== TEST ==============")
	testPassports := filereader.ReadFile("./input/day-4/test.txt")
	actualPassports := filereader.ReadFile("./input/day-4/input.txt")
	testValidPassports := filereader.ReadFile("./input/day-4/valid.txt")
	testInvalidPassports := filereader.ReadFile("./input/day-4/invalid.txt")

	count := countValidPassports(testPassports)
	if count == 2 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}

	count = countValidPassports(testValidPassports)
	if count == 4 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}
	count = countValidPassports(testInvalidPassports)
	if count == 0 {
		fmt.Println("Test succeeded:", count)
	} else {
		fmt.Println("Test failed:", count)
	}

	fmt.Println("============== OUTPUT ============")

	count = countValidPassports(actualPassports)
	fmt.Println("Number of valid passports or NPC:", count)
}

func countValidPassports(passports []string) int {
	count := 0
	startIndexPassport := 0
	for index, line := range passports {
		if len(line) == 0 {
			if fromStringArray(passports[startIndexPassport:index]).isValid(false) {
				count++
			}
			startIndexPassport = index + 1
		} else if len(passports)-1 == index {
			if fromStringArray(passports[startIndexPassport : index+1]).isValid(false) {
				count++
			}
		}
	}

	return count
}
