package main

import "fmt"

type System struct {
	workflows map[string][]Rule
}

func (this *System) process(part Part) []Part {
	var (
		workflow string = "in"
		rules    []Rule
		ok       bool
		accepted []Part
	)
	for {
		rules, ok = this.workflows[workflow]
		if !ok {
			panic(fmt.Sprintf("workflow '%s' does not exist", workflow))
		}
		for _, rule := range rules {
			if rule.Applies(part) {
				workflow = rule.Destination()
				break
			}
		}
		if !ok {
			panic("workflow did not succeed")
		}
		switch workflow {
		case "A":
			accepted = append(accepted, part)
			return accepted
		case "R":
			return accepted
		}
	}
}

func part1(workflows map[string][]Rule, parts []Part) {
	system := System{workflows}
	result := 0
	for _, part := range parts {
		accepted := system.process(part)
		for _, item := range accepted {
			result += item.Sum()
		}
	}
	fmt.Println(result)
}
