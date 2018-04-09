package main

import (
	"bufio"
	"fmt"
	"os"
	"trie"
	"trie/dictionary"
)

func main() {
	t := trie.NewTrie()

	file, err := os.Open("../data/english-words/dictionary.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(fmt.Sprintf("Error opeining file %v, msg %s", file, err.Error()))
		os.Exit(1)
	}

	s := bufio.NewScanner(file)
	ws := dictionary.NewWordScanner()
	words := ws.Scan(s)

	for wd := range words {
		t.Add(wd)
	}

	fmt.Println("Done scanning the dictionary.")

	results := t.GetMatches("epi")

	for _, result := range results {
		fmt.Println(result)
	}
}
