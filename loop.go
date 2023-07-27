package main

import (
	"fmt"
)

type Ranking map[string]uint

func beginLoop(data Data) {
	comparisons := determineComparisons(data)
	choices := rankingFromComparisons(comparisons)

	for i, name := range comparisons {
		choice := question(name, i)
		choices[choice] += 1
	}
	// isDone := question(choices, 1)

	fmt.Print(choices)
}

func question(choices string, number int) (choice string) {
	choice1, choice2 := deserialize(choices)

	fmt.Print(color(fmt.Sprintf("#%d — \n", number), 2))

	return binaryChoice([2]string{choice1, choice2}, number == 0)
}

func rankingFromComparisons(comparisons Comparisons) Ranking {
	choices := Ranking{}
	for _, name := range comparisons {
		choice1, choice2 := deserialize(name)
		choices[choice1] = 0
		choices[choice2] = 0
	}
	return choices
}
