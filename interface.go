package main

import (
	"fmt"
	"time"
)

func color(string string, ansi uint16) string {
	ansiString := fmt.Sprintf("\x1b[%dm", ansi)
	colored := fmt.Sprintf("%s%s%s", ansiString, string, "\x1b[0m")
	return colored
}

func binaryChoice(choice [2]string, arrows bool) string {
	firstChoice := choice[0]
	secondChoice := choice[1]

	formatArrow := func(direction Key) string {
		if direction == LeftArrow && arrows {
			return "(←) "
		} else if direction == RightArrow && arrows {
			return "(→)"
		}
		return ""
	}

	formatChoices := func(selected Key) string {
		if selected == LeftArrow {
			firstChoice = color(firstChoice, 36)
		} else if selected == RightArrow {
			secondChoice = color(secondChoice, 36)
		}
		return fmt.Sprintf("%s %s%s %s %s", firstChoice, formatArrow(LeftArrow), color("|", 2), secondChoice, formatArrow(RightArrow))
	}
	fmt.Print(formatChoices(OtherKey))

	arrow := getKey()
	if arrow == OtherKey {
		eraseLine()
		return binaryChoice(choice, arrows)
	}
	eraseLine()
	// Highlight the right choice
	fmt.Println(formatChoices(arrow))
	time.Sleep(200 * time.Millisecond)

	eraseLine()
	return choice[arrow] // This is a bit confusing since arrow is part of an enum 😅
}

func eraseLine() {
	fmt.Printf("\r\x1b[K")
}
