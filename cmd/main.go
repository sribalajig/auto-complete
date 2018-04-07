package main

import (
	"fmt"
	"trie"
)

func main() {
	t := trie.NewTrie()

	t.Add("Rambunctious")

	t.Add("Sublime")

	t.Add("Substitue")

	t.Add("Rampart")

	fmt.Println("Done creating the trie")
}
