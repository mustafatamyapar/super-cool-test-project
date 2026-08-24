package orders

// Queue is the line of orders waiting on the barista.
type Queue struct {
	items []Order
}

func (q *Queue) Push(o Order) { q.items = append(q.items, o) }

func (q *Queue) Pop() (Order, bool) {
	if len(q.items) == 0 {
		return Order{}, false
	}
	next := q.items[0]
	q.items = q.items[1:]
	return next, true
}

func (q *Queue) Len() int { return len(q.items) }
