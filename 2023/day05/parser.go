package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parse(file *os.File) Almanac {
	scanner := bufio.NewScanner(file)

	var seeds []int
	maps := make(map[string]Map)

	scanner.Scan()
	line := scanner.Text()
	for _, str := range strings.Fields(strings.Split(line, ":")[1]) {
		seed, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	scanner.Scan()

	for {
		m, ok := parseMap(scanner)
		if !ok {
			break
		}
		maps[m.source] = m
	}

	return Almanac{seeds, maps}
}

func parseMap(scanner *bufio.Scanner) (Map, bool) {
	var mappings []Mapping

	for {
		if scanner.Text() != "" {
			break
		}
		if !scanner.Scan() {
			return Map{"", "", mappings}, false
		}
	}

	re := regexp.MustCompile("([a-z]+)-to-([a-z]+) map:")
	line := scanner.Text()
	matched := re.FindStringSubmatch(line)
	if len(matched) != 3 {
		panic("failed to parse name")
	}

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 3 {
			panic("failed to parse range")
		}

		dest, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		source, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}

		mappings = append(mappings, Mapping{dest, source, length})
	}

	return Map{matched[1], matched[2], mappings}, true
}
