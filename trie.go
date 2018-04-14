package trie

import (
	"fmt"
)

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			Value:    '*',
			Children: []*Node{},
			Parent:   nil,
		},
	}
}

type Node struct {
	Value    rune
	Children []*Node
	Parent   *Node
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}

func (t *Trie) Add(word string) {
	cur := t.Root

	for _, ch := range word {
		n := find(ch, cur)

		if n == nil {
			newNode := &Node{
				Value:    ch,
				Children: []*Node{},
				Parent:   cur,
			}

			cur.Add(newNode)

			cur = newNode
		} else {
			cur = n
		}
	}
}

func (t *Trie) GetMatches(prefix string) []string {
	if prefix == "" {
		return []string{}
	}

	cur := t.Root

	for _, c := range prefix {
		n := find(c, cur)

		if n == nil {
			return []string{}
		}

		cur = n
	}

	return depthFirstEnum(cur, prefix[:len(prefix)-1])
}

func (t *Trie) MatchAnywhere(s string) []string {
	if s == "" {
		return []string{}
	}

	q := new(queue)
	q.Enqueue(t.Root)

	var results = []string{}

	for _, c := range s {
		var nodes []*Node

		for !q.IsEmpty() {
			node, _ := q.Dequeue()
			nodes = append(nodes, node)
		}

		for _, node := range nodes {
			children := findAnywhere(c, node)

			for _, child := range children {
				q.Enqueue(child)
			}
		}

		if q.IsEmpty() {
			return results
		}
	}

	for !q.IsEmpty() {
		n, _ := q.Dequeue()

		var prefix string
		temp := n.Parent

		for temp.Parent != nil {
			prefix = fmt.Sprintf("%s%c", prefix, temp.Value)
			temp = temp.Parent
		}

		r := depthFirstEnum(n, reverse(prefix))

		for _, item := range r {
			results = append(results, item)
		}
	}

	return results
}

func depthFirstEnum(n *Node, s string) []string {
	var results []string

	if n.Value != '*' {
		s = fmt.Sprintf("%s%c", s, n.Value)
	}

	if len(n.Children) == 0 {
		return []string{s}
	}

	for _, c := range n.Children {
		r := depthFirstEnum(c, s)

		for _, result := range r {
			results = append(results, result)
		}
	}

	return results
}

func find(ch rune, n *Node) *Node {
	for _, c := range n.Children {
		if c.Value == ch {
			return c
		}
	}

	return nil
}

func findAnywhere(ch rune, n *Node) []*Node {
	var results []*Node

	for _, child := range n.Children {
		r := dfs(child, ch)

		for _, item := range r {
			results = append(results, item)
		}
	}

	return results
}

func dfs(root *Node, ch rune) []*Node {
	var result []*Node

	if root.Value == ch {
		return []*Node{root}
	}

	for _, c := range root.Children {
		r := dfs(c, ch)

		for _, item := range r {
			result = append(result, item)
		}
	}

	return result
}

func reverse(s string) string {
	t := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}

	return string(t)
}
