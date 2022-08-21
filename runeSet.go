package main

import (
	"sort"
	"strings"
)

type RuneSet map[rune]struct{}

var exists = struct{}{}

func newRuneSet(word string) RuneSet {
	runeSet := make(RuneSet)
	for _, letter := range word {
		runeSet[letter] = exists
	}
	return runeSet
}

func (r RuneSet) String() string {
	output := []string{}
	for letter := range r {
		output = append(output, string(letter))
	}
	sort.Strings(output)
	return strings.Join(output, "")
}

func (r RuneSet) Union(d RuneSet) RuneSet {
	union := make(RuneSet)
	for key, value := range r {
		union[key] = value
	}
	for key, value := range d {
		union[key] = value
	}
	return union
}

func (r RuneSet) Intersection(d RuneSet) int {
	return (len(r) + len(d)) - len(r.Union(d))
}
