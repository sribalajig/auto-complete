package trie

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"trie/dictionary"
)

var results []string

func TestMatchPrefix(t *testing.T) {
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

	matches := tr.MatchPrefix("Rap")
	if len(matches) != 3 {
		t.Logf("Expected %d results, got %d", 2, len(matches))
		t.Log(matches)
		t.Fail()
	}

	matches = tr.MatchPrefix("R")
	if len(matches) != 8 {
		t.Log(matches)
		t.Logf("Expected %d matches, got %d", 8, len(matches))
		t.Fail()
	}

	matches = tr.MatchPrefix("Ra")
	if len(matches) != 7 {
		t.Log(matches)
		t.Logf("Expected %d matches, got %d", 8, len(matches))
		t.Fail()
	}

	matches = tr.MatchPrefix("B")
	if len(matches) != 1 {
		t.Log(matches)
		t.Logf("Expected %d matches, got %d", 1, len(matches))
		t.Fail()
	}
}

func TestMatchAnywhere(t *testing.T) {
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

	r := tr.MatchAnywhere("Rp")
	if len(r) != 4 {
		t.Logf("Expected %d result(s), got %d, %s", 4, len(r), r)
		t.Fail()
	}

	r = tr.MatchAnywhere("Rz")
	if len(r) != 1 {
		t.Logf("Expected %d result(s), got %d, %s", 1, len(r), r)
		t.Fail()
	}

	r = tr.MatchAnywhere("au");
	if len(r) != 3 {
		t.Logf("Expected %d result(s), got %d, %s", 3, len(r), r)
		t.Fail()
	}
}

func BenchmarkMatchAnywhereSmall(b *testing.B) {
	var r []string

	for n := 0; n < b.N; n++ {
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
