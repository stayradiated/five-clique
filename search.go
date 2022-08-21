package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type SearchWord struct {
	word         string
	neighbourSet NeighbourSet
}

func search() {
	fmt.Println("Loading graph")

	words := []SearchWord{}

	f, err := os.Open("./word_graph.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		word := row[0]
		neighbours := []int{}
		if err = json.Unmarshal([]byte(row[1]), &neighbours); err != nil {
			panic(err)
		}
		neighbourSet := newNeighbourSet(neighbours)
		words = append(words, SearchWord{word, neighbourSet})
	}

	fmt.Println("Searching for 5-word cliques")

	cliques := [][]string{}
	for _, wordA := range words {
		neighboursA := wordA.neighbourSet
		for b := range neighboursA {
			wordB := words[b]
			// the remaining candidates are only the words in the intersection
			// of the neighborhood sets of i and j
			neighboursAB := neighboursA.Intersection(wordB.neighbourSet)
			for c := range neighboursAB {
				wordC := words[c]
				neighboursABC := neighboursAB.Intersection(wordC.neighbourSet)
				for d := range neighboursABC {
					wordD := words[d]
					neighboursABCD := neighboursABC.Intersection(wordD.neighbourSet)
					// all remaining neighbours form a 5-clique with i, j, k, and l
					for e := range neighboursABCD {
						wordE := words[e]
						fmt.Println(wordA.word, wordB.word, wordC.word, wordD.word, wordE.word)
						cliques = append(cliques, []string{
							wordA.word, wordB.word, wordC.word, wordD.word, wordE.word,
						})
					}
				}
			}
		}
	}

	fmt.Println("completed! Found", len(cliques), "cliques")

	fmt.Println("Write to output")

	f, err = os.Create("./cliques.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	for _, cliq := range cliques {
		err = writer.Write(cliq)
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		panic(err)
	}
}
