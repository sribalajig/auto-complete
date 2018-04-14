package trie

import (
	"testing"
)

func Test_Enqueue(t *testing.T) {
	q := new(queue)

	enqueue('a', q)

	if q.IsEmpty() {
		t.Log("Expected queue to not be empty")
		t.Fail()
	}

	enqueue('b', q)

	enqueue('c', q)

	n, _ := q.Dequeue()
	if n.Value != 'a' {
		t.Logf("Expected '%c' got '%c'", 'a', n.Value)
		t.Fail()
	}

	n, _ = q.Dequeue()
	if n.Value != 'b' {
		t.Logf("Expected '%c' got '%c'", 'b', n.Value)
		t.Fail()
	}

	n, _ = q.Dequeue()
	if n.Value != 'c' {
		t.Logf("Expected '%c' got '%c'", 'c', n.Value)
		t.Fail()
	}

	if !q.IsEmpty() {
		t.Log("Expected queue to be empty, but it is not")
		t.Fail()
	}
}

func enqueue(c rune, q *queue) {
	n := new(Node)
	n.Value = c

	q.Enqueue(n)
}
