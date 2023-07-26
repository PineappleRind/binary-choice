package main

import (
	"fmt"
)

type Choices map[string]uint

func beginLoop(data Data) {
	choices := choicesFromData(data)
	isDone := question(choices, true)
	if isDone == true {
		fmt.Print(choices)
	}
}

func question(choices Choices, firstTime bool) (isDone bool) {
	choice1, choice2, ok := findTie(choices)

	if ok == false {
		return true
	}
	var choice string
	choiceSlice := []string{choice1, choice2}
	if found, ok := getCache(choiceSlice); ok {
		choice = choiceSlice[found]
	} else {
		choice = binaryChoice([2]string{choice1, choice2}, firstTime)
		addCache(choiceSlice, choice)
	}
	choices[choice] += 1
	return question(choices, false)
}

func findTie(choices Choices) (choice1 string, choice2 string, ok bool) {
	seenScores := map[uint]string{}
	// due to map iterations being random,
	// `choices` is already randomized
	for name, score := range choices {
		if _, alreadySeen := seenScores[score]; !alreadySeen {
			seenScores[score] = name
			continue
		}
		// If we have seen this score already,
		// return this name along with seenScores[score]
		return seenScores[score], name, true
	}
	// no ties left: we have a winner!!
	return "", "", false
}

func choicesFromData(data Data) Choices {
	choices := Choices{}
	for _, name := range data {
		choices[name] = 0
	}
	return choices
}
