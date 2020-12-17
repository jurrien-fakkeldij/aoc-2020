package main

import (
	"aoc-2020/internal/filereader"
	"aoc-2020/internal/transformer"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Day16()
}

type Ticket struct {
	ticketNumbers []int
}

type TicketRule struct {
	lower int
	upper int
}

type ticketRuleSlice []TicketRule

func Day16() {
	fmt.Println("==============  Day16 =============")
	_, otherTickets, ticketRules := parseTicketData(filereader.ReadFile("./input/day-16/test.txt"))

	fmt.Println("==============  TEST  =============")
	errorRate := determineErrorRate(otherTickets, ticketRules)
	if errorRate == 71 {
		fmt.Println("Test succeeded:", errorRate)
	} else {
		fmt.Println("Test failed:", errorRate)
	}

	yourTicket, otherTickets, ticketRules := parseTicketData(filereader.ReadFile("./input/day-16/test-2.txt"))
	determineTicketLayout(yourTicket, removedInvalidTickets(otherTickets, ticketRules), ticketRules)
	fmt.Println("============== OUTPUT =============")
	yourTicket, otherTickets, ticketRules = parseTicketData(filereader.ReadFile("./input/day-16/input.txt"))
	fmt.Println("Error rate:", determineErrorRate(otherTickets, ticketRules))
	validTickets := removedInvalidTickets(otherTickets, ticketRules)
	ticketLayout := determineTicketLayout(yourTicket, validTickets, ticketRules)
	output := 1
	for index, rule := range ticketLayout {
		if strings.HasPrefix(rule, "departure") {
			output *= yourTicket.ticketNumbers[index]
		}
	}
	fmt.Println("Ticket rate:", output)
}

func removedInvalidTickets(otherTickets []Ticket, ticketRules map[string][]TicketRule) []Ticket {
	validTickets := []Ticket{}
	for _, ticket := range otherTickets {
		if determineValidTicket(ticket, ticketRules) {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func determineTicketLayout(validTicket Ticket, validTickets []Ticket, ticketRules map[string][]TicketRule) []string {
	output := make([]string, len(ticketRules))
	rulePositions := make(map[string][]int)
	for key, value := range ticketRules {
		rulePositions[key] = determineLocationsForRule(validTicket.ticketNumbers, validTickets, value, key)
	}

	foundPositions := []int{}

	for fieldsAreEmpty(output) {
		for rule, positions := range rulePositions {
			if len(positions) == 1 {
				output[positions[0]] = rule
				foundPositions = append(foundPositions, positions[0])
			}
		}

		for rule, positions := range rulePositions {
			for _, foundPosition := range foundPositions {
				position := pos(positions, foundPosition)
				if position > -1 {
					rulePositions[rule] = append(positions[:position], positions[position+1:]...)
				}
			}
		}
	}

	//fmt.Println(testOutput2)
	fmt.Println(output)
	return output
}

func pos(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

func fieldsAreEmpty(fields []string) bool {
	for _, field := range fields {
		if field == "" {
			return true
		}
	}

	return false
}

func determineLocationsForRule(ticketNumbers []int, validTickets []Ticket, ticketRules []TicketRule, name string) []int {
	testOutput := make(map[string][]int)
	testOutput2 := []int{}
	testOutput[name] = []int{}
	for index := range ticketNumbers {
		validTicketsForRule := 0
		for _, ticket := range validTickets {
			number := ticket.ticketNumbers[index]
			validTicketIndexForRule := false
			for _, rule := range ticketRules {
				if number >= rule.lower && number <= rule.upper {
					validTicketIndexForRule = true
				}
			}

			if validTicketIndexForRule {
				validTicketsForRule++
			}
		}
		//fmt.Println("rule", name, "index:", index, "validTickets:", validTicketsForRule, "len:", len(validTickets))
		if validTicketsForRule == len(validTickets) {
			testOutput[name] = append(testOutput[name], index)
			testOutput2 = append(testOutput2, index)
		}
	}

	fmt.Println("rule", name, ":", testOutput2)
	//fmt.Println(testOutput2)

	return testOutput2
}

func determineValidTicket(ticket Ticket, ticketRules map[string][]TicketRule) bool {
	invalidRules := 0
	for _, number := range ticket.ticketNumbers {
		validRule := false
		for _, value := range ticketRules {
			for _, rule := range value {
				if number >= rule.lower && number <= rule.upper {
					validRule = true
				}
			}
		}
		if !validRule {
			//fmt.Println("not valid:", number)
			invalidRules++
		} else {
			//fmt.Println("valid:", number)
		}
	}

	return invalidRules == 0
}

func determineErrorRate(otherTickets []Ticket, ticketRules map[string][]TicketRule) int {
	errorRate := 0
	for _, ticket := range otherTickets {
		for _, number := range ticket.ticketNumbers {
			validRule := false
			for _, value := range ticketRules {
				for _, rule := range value {
					//fmt.Println("rule:", key, ":", rule)
					if number >= rule.lower && number <= rule.upper {
						validRule = true
					}
				}
			}
			if !validRule {
				//fmt.Println("not valid:", number)
				errorRate += number
			} else {
				//fmt.Println("valid:", number)
			}
		}
	}
	return errorRate
}

func parseTicketData(ticketData []string) (Ticket, []Ticket, map[string][]TicketRule) {
	//first lines parse ticket rules
	//until empty line
	//your ticket:
	//next line is your ticket
	//until empty line
	//nearby tickets:
	//next are nearby tickets
	yourTicket := Ticket{}
	nearbyTickets := []Ticket{}
	ticketRules := make(map[string][]TicketRule)

	currentScheme := "TICKETRULES"

	for _, line := range ticketData {
		if line == "" {
			//noop
		} else if line == "your ticket:" {
			currentScheme = "YOURTICKET"
		} else if line == "nearby tickets:" {
			currentScheme = "NEARBYTICKET"
		} else {
			if currentScheme == "TICKETRULES" {
				name := strings.Split(line, ": ")[0]
				rules := strings.Split(line, ": ")[1]

				ticketRules[name] = parseTicketRules(rules)
			} else if currentScheme == "YOURTICKET" {
				ticketNumbers, _error := transformer.SliceAtoi(strings.Split(line, ","))
				if _error != nil {
					panic("Can't parse your ticket numbers")
				}
				yourTicket.ticketNumbers = ticketNumbers
			} else if currentScheme == "NEARBYTICKET" {
				ticket := Ticket{}
				ticketNumbers, _error := transformer.SliceAtoi(strings.Split(line, ","))
				if _error != nil {
					panic("Can't parse your ticket numbers")
				}
				ticket.ticketNumbers = ticketNumbers
				nearbyTickets = append(nearbyTickets, ticket)
			}
		}
	}
	/*fmt.Println("yourTicket:", yourTicket)
	fmt.Println("nearbyTickets:", nearbyTickets)
	fmt.Println("ticketRules:", ticketRules)*/
	return yourTicket, nearbyTickets, ticketRules
}

func parseTicketRules(rules string) []TicketRule {
	returnSlice := []TicketRule{}
	for _, rule := range strings.Split(rules, " or ") {
		lower, error := strconv.Atoi(strings.Split(rule, "-")[0])

		if error != nil {
			panic("lower can't be transformed")
		}
		upper, error := strconv.Atoi(strings.Split(rule, "-")[1])

		if error != nil {
			panic("upper can't be transformed")
		}
		returnSlice = append(returnSlice, TicketRule{lower, upper})
	}

	return returnSlice
}
