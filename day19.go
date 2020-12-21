package main

import (
	"aoc-2020/internal/filereader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Day19()
}

type Rule struct {
	isLiteral  bool
	contentRaw string
	ruleNum    string
}

func Day19() {
	fmt.Println("==============  Day18 =============")
	fmt.Println("==============  TEST  =============")
	testMessages := filereader.ReadFile("./input/day-19/test.txt")
	output := matchesRule(testMessages, "0", false)
	if output == 2 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	testMessages2 := filereader.ReadFile("./input/day-19/test2.txt")
	output = matchesRule(testMessages2, "0", false)
	if output == 3 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	output = matchesRule(testMessages2, "0", true)
	if output == 12 {
		fmt.Println("Test succeeded:", output)
	} else {
		fmt.Println("Test failed:", output)
	}

	fmt.Println("============== OUTPUT =============")
	messages := filereader.ReadFile("./input/day-19/input.txt")
	fmt.Println("Matches rule 0: ", matchesRule(messages, "0", false))
	fmt.Println("Matches rule 0: ", matchesRule(messages, "0", true))
	fmt.Println("Matches rule 0 with updated rules: ", Part2Regex(messages))
}

func Part2Regex(lines []string) int {
	rules, messages := parseInput(lines)

	rules[8] = "(42 | 42 8)"
	rules[11] = "(42 31 | 42 11 31)"

	count8 := map[int]int{}
	count11 := map[int]int{}
	maxRecurse := 10

	re := regexp.MustCompile(`\d+`)
	for {
		done := true

		for k := range rules {
			rules[k] = re.ReplaceAllStringFunc(rules[k], func(s string) string {
				done = false
				i, _ := strconv.Atoi(s)
				if i == 8 {
					if _, exists := count8[i]; !exists {
						count8[i] = 0
					}
					count8[i]++
					if count8[i] > maxRecurse {
						return "(42)"
					}
				}
				if i == 11 {
					if _, exists := count11[i]; !exists {
						count11[i] = 0
					}
					count11[i]++
					if count11[i] > maxRecurse {
						return "(42 31)"
					}
				}
				return rules[i]
			})
		}

		if done {
			break
		}
	}

	regexRules := map[int]*regexp.Regexp{}

	for k := range rules {
		replacer := strings.NewReplacer("\"", "", " ", "")
		regexRules[k] = regexp.MustCompile("^" + replacer.Replace(rules[k]) + "$")
	}

	rule0 := regexRules[0]

	count := 0
	for _, message := range messages {
		if rule0.MatchString(message) {
			count++
		}
	}
	return count
}

func parseInput(lines []string) (rules map[int]string, messages []string) {
	rules = map[int]string{}
	parts := strings.Split(strings.Join(lines, "\n"), "\n\n")
	for _, line := range strings.Split(parts[0], "\n") {
		k, _ := strconv.Atoi(strings.Split(line, ":")[0])
		v := "(" + strings.Split(line, ":")[1][1:] + ")"
		rules[k] = v
	}
	messages = strings.Split(parts[1], "\n")
	return
}

func matchesRule(data []string, ruleNum string, updatedRules bool) int {
	breakIndex := 0
	for index := range data {
		if data[index] == "" {
			breakIndex = index
		}
	}
	//break for rules and data
	rules := createRules(data[0:breakIndex], updatedRules)
	regexString := rules[ruleNum].resolveToRegex(rules, updatedRules)
	fmt.Println(regexString)
	regex := regexp.MustCompile("^" + regexString + "$")
	matchCount := 0
	for _, line := range data[breakIndex+1:] {
		matches := regex.MatchString(line)
		if matches {
			matchCount++
		}
	}
	return matchCount
}
func NewRule(str string) *Rule {
	strSplit := strings.SplitN(str, ": ", 2)
	str = strSplit[1]
	if str[0] == '"' {
		return &Rule{
			ruleNum:    strSplit[0],
			contentRaw: str[1 : len(str)-1],
			isLiteral:  true,
		}
	}
	return &Rule{
		ruleNum:    strSplit[0],
		contentRaw: str,
	}
}

func createRules(rulesList []string, updatedRules bool) map[string]*Rule {
	rules := make(map[string]*Rule)
	for _, rule := range rulesList {
		split := strings.Split(rule, ": ")
		rules[split[0]] = NewRule(rule)
	}

	if updatedRules {
		rules["8"].contentRaw = "42 | 42 8"
		rules["11"].contentRaw = "42 31 | 42 11 31"
	}

	return rules
}

var count11Time int = 0
var max11 int = 10

func (rule *Rule) resolveToRegex(rules map[string]*Rule, changedRules bool) string {
	if rule == nil {
		fmt.Println("WUT? ")
		return ""
	}

	// we're a literal, just return our content
	if rule.isLiteral {
		return rule.contentRaw
	}
	out := strings.Builder{}

	// We're more complex than a straight literal
	// Appears that we only ever get one OR, so split on that, then resolve down
	split := strings.Split(rule.contentRaw, " | ")
	orRules := []string{}
	hasSelf := false
	if changedRules {
		if rule.ruleNum == "8" {
			out.WriteString("(?:")
			out.WriteString("(")
			out.WriteString(rules["42"].resolveToRegex(rules, changedRules))
			out.WriteString(")+)")
			return out.String()
		} else if rule.ruleNum == "11" {
			if count11Time >= max11 {
				out.WriteRune('(')
				out.WriteString(rules["42"].resolveToRegex(rules, changedRules))
				out.WriteString(rules["31"].resolveToRegex(rules, changedRules))
				out.WriteRune(')')
				return out.String()
			}
			count11Time++
		}
	}

	out.WriteString("(?:")

	for _, rStr := range split {
		// resolve each rule in here, concat them together
		split := strings.Split(rStr, " ")
		concattedRules := []string{}
		for _, ruleNum := range split {
			toAdd := rules[ruleNum].resolveToRegex(rules, changedRules)
			concattedRules = append(concattedRules, toAdd)
		}
		orRules = append(orRules, strings.Join(concattedRules, ""))
	}

	out.WriteString(strings.Join(orRules, "|"))
	if hasSelf {
		out.WriteString("+")
	}
	out.WriteRune(')')
	return out.String()
}
