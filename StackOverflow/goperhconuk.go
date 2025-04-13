package __

import (
	"golang.org/x/exp/slices"
)

func addNames(names []string, name string) []string {
	const extraLines = 1000
	newNames := make([]string, len(names), len(names)+extraLines+1)
	copy(newNames, names)
	for i := 0; i < extraLines; i++ {
		newNames = append(newNames, "extra"+string(i))
	}
	newNames = append(newNames, name)
	return newNames
}

func addNames(names []string, name string) []string {
	const extraLines = 1000
	newNames := slices.Clone(names)
	slices.Grow(newNames, extraLines+1)
	for i := 0; i < extraLines; i++ {
		newNames = append(newNames, "extra"+string(i))
	}
	newNames = append(newNames, name)
	return newNames
}
