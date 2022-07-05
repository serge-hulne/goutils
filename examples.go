package main

import (
	"fmt"
	"sort"
	"strings"

	fi "github.com/serge-hulne/goutils/files"
	it "github.com/serge-hulne/goutils/iter"
	. "github.com/serge-hulne/goutils/str"
)

func main() {

	_, f := fi.Open("../../../100-0.txt")

	lines := fi.ReadLines(f)

	/*
		TODO
			- Faire une variante de Map() (par exemple : Apply()) qui retourne un type composé (slice ou type générique) qui peut être itéré à son tour.
	*/

	wordsArraysStream := it.MapSpilt(lines, SplitAndTrim)

	m := make(map[string]int)

	// Counting frequency using a map
	for words := range wordsArraysStream {
		for _, w := range words {
			m[w]++
		}
	}

	// Map to array (for sorting)
	var wordCount []Pair
	for k, v := range m {
		wordCount = append(wordCount, Pair{k, v})
	}

	/*
		TODO: create a (sortable) frequency map (as in Nim)
	*/

	// Sorting array
	sort.SliceStable(wordCount, func(i, j int) bool {
		return wordCount[i].cpt > wordCount[j].cpt
	})

	// Print top 10
	for index, item := range wordCount {
		if index < 10 {
			fmt.Printf("%v\n", item)
		}
	}

	println("- - - ")

	// Testing Str:
	s := Str{"Hello Woild"}
	fmt.Printf("%s\n", s.ReplaceAll("i", "r").Capitalize())
	println(s.Str)
}

type Pair struct {
	word string
	cpt  int
}

func SplitAndTrim(x string) []string {
	x = strings.ToLower(x)
	words := strings.Split(x, " ")
	var temp []string
	for _, w := range words {
		w = strings.Trim(w, " ’',./()\";:!?")
		if len(w) > 0 {
			temp = append(temp, w)
		}
	}
	return temp
}
