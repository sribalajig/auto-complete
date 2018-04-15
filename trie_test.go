package trie

import (
	"testing"
)

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
