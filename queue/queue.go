package queue

import (
	"errors"
	"fmt"
	"sync"
)

type queueItems []interface{}

type Queue struct {
	items queueItems
	len   int
	mutex sync.Mutex
}

// Add will add item to the Queue
func (q *Queue) Add(item interface{}) {

	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
	q.len = q.len + 1
}

// Pop will pop out first element in queue
func (q *Queue) Pop() (interface{}, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	var item interface{}
	if q.len <= 0 {
		return nil, errors.New("no elements in queue")
	}
	item, q.items = q.items[0], q.items[1:]
	return item, nil
}

// Len will give length of Queue
func (q *Queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items)
}

// IsEmpty will return bool whether Queue is empty or not
func (q *Queue) IsEmpty() bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	return len(q.items) == 0
}

// Print will simply log items present in Queue
func (q *Queue) Print() {
	for item := range q.items {
		fmt.Println(q.items[item])
	}
}

// Free will remove all elements in Queue
func (q *Queue) Free() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.len = 0
	q.items = nil
}

// NewQueue is constructor to initialize our Queue
func NewQueue() *Queue {
	return &Queue{len: 0}
}
