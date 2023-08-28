package queue

type Queue struct {
	TotalCapacity int
	Items         []interface{}
}

func New() *Queue {
	return &Queue{
		TotalCapacity: 0,
		Items:         make([]interface{}, 0),
	}
}

func (q *Queue) Allocate() {

}

func (q *Queue) GetCapacity() int {
	return q.TotalCapacity

}

func (q *Queue) Clear() {
	q.Items = make([]interface{}, 0)
}

func (q *Queue) Copy() *Queue {
	return &Queue{
		TotalCapacity: q.TotalCapacity,
		Items:         q.Items,
	}
}

func (q *Queue) Count() int {
	return len(q.Items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Peek() interface{} {
	return q.Items[0]
}

func (q *Queue) Pop() interface{} {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item
}

func (q *Queue) Push(item interface{}) {
	q.Items = append(q.Items, item)
}

func (q *Queue) ToSlice() []interface{} {
	return q.Items
}
