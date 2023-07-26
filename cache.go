package main

import (
	"encoding/json"
)

type CachedChoice map[string]int

var cache = CachedChoice{}

func addCache(key []string, choice string) bool {
	serialized := serialize(key)
	index := indexOf(key, choice)
	if index == -1 {
		return false
	}
	cache[serialized] = index
	return true
}

func getCache(key []string) (found int, ok bool) {
	serialized := serialize(key)
	found, ok = cache[serialized]
	return found, ok
}

func serialize(key []string) string {
	serialized, err := json.Marshal(key)
	if err != nil {
		panic(err)
	}
	return string(serialized)
}

func indexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}
