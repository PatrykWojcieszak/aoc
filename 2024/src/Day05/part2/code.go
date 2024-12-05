package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

func findElement(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func removeValue(slice []string, value string) []string {
	result := []string{}

	for _, v := range slice {
		if v != value {
			result = append(result, v)
		}
	}

	return result
}

func isMatchingRule(protocols []string, pageOne string, pageTwo string) bool {
	for _, protocol := range protocols {
		splittedProtocol := strings.Split(protocol, "|")
		if splittedProtocol[0] == pageOne && splittedProtocol[1] == pageTwo {
			return true
		}
	}

	return false
}

func checkPageRules(protocols []string, rule []string) bool {
	isMatching := true

	if len(rule) <= 1 {
		return true
	}

	for ruleIndex := 1; ruleIndex < len(rule); ruleIndex++ {
		isMatch := isMatchingRule(protocols, rule[0], rule[ruleIndex])

		if !isMatch {
			isMatching = false
		}
	}

	if isMatching {
		return checkPageRules(protocols, rule[1:])
	}

	return isMatching
}

func checkSecondPage(protocols []string, firstPage string, secondPage []string) bool {
	for _, page := range secondPage {
		if !findElement(protocols, firstPage+"|"+page) {
			return false
		}
	}

	return true
}

func fixRule(protocols []string, rule []string) []string {
	var fixedRule []string

	for {
		if len(rule) <= 1 {
			fixedRule = append(fixedRule, rule[0])
			rule = removeValue(rule, rule[0])
		}

		for _, char := range rule {
			if checkSecondPage(protocols, char, removeValue(rule, char)) {
				fixedRule = append(fixedRule, char)
				rule = removeValue(rule, char)
			}
		}

		if len(rule) == 0 {
			break
		}
	}

	return fixedRule
}

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	var protocols, pageRules, invalidRules []string
	result := 0

	for _, line := range output {
		if strings.Contains(line, "|") {
			protocols = append(protocols, line)
		}

		if strings.Contains(line, ",") {
			pageRules = append(pageRules, line)
		}
	}

	for _, rules := range pageRules {
		splittedRules := strings.Split(rules, ",")

		if !checkPageRules(protocols, splittedRules) {
			invalidRules = append(invalidRules, rules)
		}
	}

	for _, rules := range invalidRules {
		splittedRules := strings.Split(rules, ",")

		fixedRule := fixRule(protocols, splittedRules)

		middlePage, _ := strconv.Atoi(string(fixedRule[(len(fixedRule)-1)/2]))
		result += middlePage
	}

	fmt.Println(result)
}
