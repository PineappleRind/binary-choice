package main

import (
	"strings"
)

type Comparisons []string

func determineComparisons(choices Data) Comparisons {
	combinations := []string{}
	for i := 0; i < len(choices); i++ {
		for j := i + 1; j < len(choices); j++ {
			combinations = append(combinations, serialize([]string{choices[i], choices[j]}))
		}
	}
	return combinations
}

func serialize(options []string) string {
	return strings.Join(options, ",")
}

func deserialize(options string) (choice1 string, choice2 string) {
	split := strings.Split(options, ",")
	return split[0], split[1]
}
