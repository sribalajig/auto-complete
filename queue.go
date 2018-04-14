package trie

import (
	"errors"
)

type queue struct {
	front *queueItem
	rear  *queueItem
}

type queueItem struct {
	val  *Node
	next *queueItem
}

func (q *queue) IsEmpty() bool {
	return q.front == nil && q.rear == nil
}

func (q *queue) Enqueue(n *Node) {
	qItem := new(queueItem)
	qItem.val = n

	if q.IsEmpty() {
		q.front = qItem
		q.rear = qItem

		return
	}

	q.rear.next = qItem
	q.rear = qItem
}

func (q *queue) Dequeue() (*Node, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	if q.front == q.rear {
		t := q.front

		q.front = nil
		q.rear = nil

		return t.val, nil
	}

	t := q.front
	q.front = t.next

	return t.val, nil
}
