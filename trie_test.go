package trie

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"trie/dictionary"
	"github.com/stretchr/testify/assert"
)

var results []string

func createTrie() *Trie {
	tr := NewTrie()

	tr.Add("Raven")
	tr.Add("Raid")
	tr.Add("Rapid")
	tr.Add("Rambunctious")
	tr.Add("Rapacious")
	tr.Add("Razor")
	tr.Add("Respect")
	tr.Add("Rapport")
	tr.Add("Conspicuous")
	tr.Add("Beauty")

	return tr
}

var prefixTestCases = []struct{
	Prefix string
	ExpectedMatches []string
} {
	{ "Rap", []string{"Rapacious", "Rapport", "Rapid"}},
	{ "Ra", []string{
		"Raven",
		"Raid",
		"Rapid",
		"Rapacious",
		"Rapport",
		"Rambunctious",
		"Razor"},
	},
	{ "R",
		[]string{
			"Rapacious",
			"Rapport",
			"Raven",
			"Raid",
			"Rambunctious",
			"Respect",
			"Razor",
			"Rapid",
		},
	},
	{ "B", []string{"Beauty"}},
}


func TestMatchPrefix(t *testing.T) {
	tr := createTrie()

	for _, testCase := range prefixTestCases {
		results := tr.MatchPrefix(testCase.Prefix)

		assert.ElementsMatch(t, testCase.ExpectedMatches, results)
	}
}

var matchAnywhereTestCases = []struct{
	Substring string
	ExpectedMatches []string
} {
	{ "Rp", []string{"Rapacious", "Rapport", "Rapid", "Respect"}},
	{ "Rz", []string{"Razor"},},
	{ "au", []string{"Rapacious", "Rambunctious", "Beauty"}},
	{ "B", []string{"Beauty"}},
}

func TestMatchAnywhere(t *testing.T) {
	tr := createTrie()

	for _, testCase := range matchAnywhereTestCases {
		results := tr.MatchAnywhere(testCase.Substring)

		assert.ElementsMatch(t, testCase.ExpectedMatches, results)
	}
}

func BenchmarkMatchAnywhereSmall(b *testing.B) {
	var r []string

	for n := 0; n < b.N; n++ {
		tr := createTrie()

		r = tr.MatchAnywhere("Rap")
		results = r
	}
}

func BenchmarkMatchPrefix(b *testing.B) {
	var r []string

	for n := 0; n < b.N; n++ {
		t := NewTrie()

		for wd := range scanWords() {
			t.Add(wd)
		}

		r = t.MatchAnywhere("freshn")
		results = r
	}
}

func BenchmarkMatchAnywhere(b *testing.B) {
	var r []string

	for n := 0; n < b.N; n++ {
		t := NewTrie()

		for wd := range scanWords() {
			t.Add(wd)
		}

		r = t.MatchPrefix("freshn")
		results = r
	}
}

func BenchmarkFileScan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		scanWords()
	}
}

func scanWords() <-chan string {
	file, err := os.Open("./data/english-words/dictionary.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opeining file %v, msg %s", file, err.Error()))
		os.Exit(1)
	}

	s := bufio.NewScanner(file)
	ws := dictionary.NewWordScanner()

	return ws.Scan(s)
}
