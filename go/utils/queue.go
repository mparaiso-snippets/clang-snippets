package utils

type Queue struct {
	First  *Node
	Last   *Node
	length int
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.First == nil
}

func (q *Queue) Enqueue(item Any) {
	oldLast := q.Last
	q.Last = &Node{}
	q.Last.Item = item
	q.Last.Next = nil
	if q.IsEmpty() {
		q.First = q.Last
	} else {
		oldLast.Next = q.Last
	}
	q.length++
}

func (q *Queue) Dequeue() Any {
	item := q.First.Item
	q.First = q.First.Next
	if q.IsEmpty() {
		q.Last = nil
	}
	q.length--
	return item
}
