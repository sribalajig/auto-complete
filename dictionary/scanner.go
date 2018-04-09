package dictionary

import (
	"bufio"
)

type Scanner interface {
	Scan(s *bufio.Scanner) <-chan string
}

type WordScanner struct {
}

func NewWordScanner() Scanner {
	return &WordScanner{}
}

func (w *WordScanner) Scan(s *bufio.Scanner) <-chan string {
	words := make(chan (string))

	go func() {
		for s.Scan() {
			w := s.Text()

			words <- w
		}

		close(words)
	}()

	return words
}
