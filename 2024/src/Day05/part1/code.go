package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"aoc/2024/pkg/file"
)

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

func main() {
	absPathName, _ := filepath.Abs("input.txt")
	output, _ := file.ReadInput(absPathName)

	var protocols, pageRules []string
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

		if checkPageRules(protocols, splittedRules) {
			middlePage, _ := strconv.Atoi(string(splittedRules[(len(splittedRules)-1)/2]))
			result += middlePage
		}
	}

	fmt.Println(result)
}
