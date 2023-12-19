package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parse(scanner *bufio.Scanner) (Workflows, []Part) {
	workflows := make(Workflows)

	workflow_re := regexp.MustCompile(`([a-z]+){([^}]+)}`)
	rule_re := regexp.MustCompile(`([a-z])([<>])(\d+):([A-Za-z]+)`)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		match := workflow_re.FindStringSubmatch(line)
		fields := strings.Split(match[2], ",")

		var rules []Rule
		for _, field := range fields {
			match := rule_re.FindStringSubmatch(field)
			if len(match) > 0 {
				val, err := strconv.Atoi(match[3])
				if err != nil {
					panic(err)
				}
				var check Check
				switch match[2] {
				case "<":
					check = Less{rune(match[1][0]), val}
				case ">":
					check = Greater{rune(match[1][0]), val}
				default:
					panic(fmt.Sprintf("unknown condition: %s", match[2]))
				}
				rules = append(rules, Conditional{match[4], check})
			} else {
				rules = append(rules, SendTo{field})
			}
		}

		workflows[match[1]] = rules
	}

	var parts []Part
	for scanner.Scan() {
		part := make(map[rune]int)
		line := scanner.Text()
		fields := strings.Split(line[1:len(line)-1], ",")
		for _, field := range fields {
			pair := strings.Split(field, "=")
			val, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}
			part[rune(pair[0][0])] = val
		}
		parts = append(parts, part)
	}

	return workflows, parts
}
