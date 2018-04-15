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

func (t *Trie) MatchPrefix(prefix string) []string {
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

func find(ch rune, n *Node) *Node {
	for _, c := range n.Children {
		if c.Value == ch {
			return c
		}
	}

	return nil
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

func (t *Trie) MatchAnywhere(s string) []string {
	var results []string

	curLevel := []*Node{t.Root}

	for _, c := range s {
		curLevel = getMatchingChildren(curLevel, c)

		if len(curLevel) == 0 {
			return []string{}
		}
	}

	for _, n := range curLevel {
		reversePrefix := enumPathToRoot(n)

		r := depthFirstEnum(n, reverse(reversePrefix))

		for _, item := range r {
			results = append(results, item)
		}
	}

	return results
}

func getMatchingChildren(nodes []*Node, c rune) []*Node {
	q := new(queue)

	for _, node := range nodes {
		children := findAnywhere(c, node)

		for _, child := range children {
			q.Enqueue(child)
		}
	}

	var result []*Node

	for !q.IsEmpty() {
		n, _ := q.Dequeue()
		result = append(result, n)
	}

	return result
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

func enumPathToRoot(n *Node) string {
	var path string
	temp := n.Parent

	for temp.Parent != nil {
		path = fmt.Sprintf("%s%c", path, temp.Value)
		temp = temp.Parent
	}

	return path
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
