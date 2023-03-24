package dotignore

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type DotIgnore struct {
	patterns []pattern
}

func FromFile(path string) *DotIgnore {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matcher, _ := regexp.Compile(`\s+`)
	patterns := []pattern{}
	for scanner.Scan() {
		line := scanner.Text()

		match := matcher.MatchString(line)
		if match {
			continue
		}

		if line[0] == '#' {
			continue
		}

		var inverted bool = false
		if line[0] == '!' {
			inverted = true
			line = line[1:]
		}

		patterns = append(patterns, pattern{
			inverted: inverted,
			value:    line,
		})
	}

	return &DotIgnore{
		patterns: patterns,
	}
}

func (ignore *DotIgnore) Matches(path string) bool {
	if len(path) == 0 {
		return false
	}

	var hit bool = false
	for _, pattern := range ignore.patterns {
		if pattern.Matches(path) {
			hit = !pattern.inverted
		}
	}

	return hit
}

type pattern struct {
	inverted bool
	value    string
}

func (pattern *pattern) Matches(path string) bool {
	if len(path) == 0 {
		return false
	}
	return compare(strings.Split(path, "/"), strings.Split(pattern.value, "/"))
}

func compare(parts []string, patterns []string) bool {
	if len(patterns) == 0 {
		return true
	}

	i, j := 0, 0
	for i < len(parts) && j < len(patterns) {
		part := parts[i]
		px := patterns[j]

		if px == "*" || px == "**" {
			return compare(parts[i:], patterns[j+1:])
		}

		match, _ := regexp.MatchString(px, part)
		if match {
			j++
		} else if j > 0 {
			return false
		}
		i++
	}
	return i == len(parts) && j == len(patterns)
}
