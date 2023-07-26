package main

import "fmt"

func color(string string, ansi uint16) string {
	ansiString := fmt.Sprintf("\x1b[%dm", ansi)
	colored := fmt.Sprintf("%s%s%s", ansiString, string, "\x1b[0m")
	return colored
}

func binaryPrompt(prompt string, choice [2]string, arrows bool) {
	if prompt != "" {
		fmt.Println(color(prompt, 36))
	}
	formatArrow := func(direction Key) string {
		if arrows != true {
			return ""
		}
		if direction == LeftArrow {
			return "(←) "
		} else if direction == RightArrow {
			return "(→)"
		}
		return ""
	}

	formatChoices := func(selected Key) string {
		firstChoice := choice[0]
		secondChoice := choice[1]
		if selected == LeftArrow {
			firstChoice = color(color(firstChoice, 4), 36)
			secondChoice = color(secondChoice, 2)
		} else if selected == RightArrow {
			firstChoice = color(firstChoice, 2)
			secondChoice = color(color(secondChoice, 4), 36)
		}
		return fmt.Sprintf("%s %s%s %s %s", firstChoice, formatArrow(LeftArrow), color("|", 2), secondChoice, formatArrow(RightArrow))
	}
	fmt.Print(formatChoices(OtherKey))

	arrow := getKey()
	// Move cursor back to first column
	fmt.Printf("\r\x1b[K")
	fmt.Println(formatChoices(arrow))
}
