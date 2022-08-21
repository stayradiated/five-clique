package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Word struct {
	word       string
	runeSet    RuneSet
	neighbours []int
}

func generate() {
	f, err := os.Open("./words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("Reading words file")

	words := []*Word{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) != 5 {
			continue
		}
		runeSet := newRuneSet(word)
		if len(runeSet) != 5 {
			continue
		}
		words = append(words, &Word{word, runeSet, []int{}})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Building neighborhoods")

	lenWords := len(words)

	for i, wordA := range words {
		for j := i; j < lenWords; j += 1 {
			wordB := words[j]
			if wordA.runeSet.Intersection(wordB.runeSet) == 0 {
				wordA.neighbours = append(wordA.neighbours, j)
			}
		}
	}

	fmt.Println("Write to output")

	f, err = os.Create("./word_graph.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	for _, word := range words {
		neighboursJson, err := json.Marshal(word.neighbours)
		if err != nil {
			panic(err)
		}
		err = writer.Write([]string{word.word, string(neighboursJson)})
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		panic(err)
	}
}
