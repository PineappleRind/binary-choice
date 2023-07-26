package main

import (
	"github.com/eiannone/keyboard"
)

type Key int

const (
	LeftArrow  Key = 0
	RightArrow     = 1
	OtherKey       = 2
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

	return determineArrowDirection(key)
}

func determineArrowDirection(key keyboard.Key) Key {
	if key == 65515 {
		return LeftArrow
	} else if key == 65514 {
		return RightArrow
	} else {
		return OtherKey
	}
}
