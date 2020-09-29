package queue

import (
	"errors"
	"fmt"
)

type queueItems []interface{}

type Queue struct {
	items queueItems
	len int
	first int
}

func (q *Queue) Add(item interface{}) (bool, error) {

	q.items = append(q.items, item)
	q.len = q.len + 1
	return true, nil
}

func (q *Queue) Pop() (interface{}, error)  {
	var item interface{}
	if q.len <=0 {
		return nil, errors.New("no elements in queue")
	}
	item, q.items = q.items[q.first], q.items[1:]
	return item, nil
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}



func (q *Queue) Print()  {
	for item := range q.items {
		fmt.Println(q.items[item])
	}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.len = 0
	q.first = 0
	return q
}
