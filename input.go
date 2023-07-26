package main

import (
	"os"

	"github.com/eiannone/keyboard"
)

type Key int

const (
	LeftArrow  Key = 0
	RightArrow     = 1
	Escape         = 2
	OtherKey       = 3
)

func getKey() Key {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	direction := determineArrowDirection(key)
	if direction == Escape {
		os.Exit(0)
	}
	return direction
}

func determineArrowDirection(key keyboard.Key) Key {
	if key == 65515 {
		return LeftArrow
	} else if key == 65514 {
		return RightArrow
	} else if key == 27 {
		return Escape
	} else {
		return OtherKey
	}
}
