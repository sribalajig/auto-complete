package trie

import "fmt"

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			Value:    '*',
			Children: []*Node{},
		},
	}
}

func (t *Trie) Add(word string) {
	cur := t.Root

	for _, ch := range word {
		n := find(ch, cur)

		if n == nil {
			newNode := &Node{
				Value:    ch,
				Children: []*Node{},
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

type Node struct {
	Value    rune
	Children []*Node
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
}
