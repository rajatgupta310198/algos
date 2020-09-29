package queue

import (
	"testing"
)

func TestQueue_Add(t *testing.T) {
	myTestQueue := NewQueue()
	_, err := myTestQueue.Add(1)

	_, err = myTestQueue.Add(3)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(2)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(3)
	if err != nil {
		t.Error(err)
	}

}

func TestQueue_Pop(t *testing.T) {
	myTestQueue := NewQueue()
	_, err := myTestQueue.Add(1)

	_, err = myTestQueue.Add(3)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(2)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(3)
	if err != nil {
		t.Error(err)
	}
	item, err := myTestQueue.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != 1 {
		t.Error("No proper implementation of queue")
	}

	myTestQueue2 := NewQueue()
	item, err = myTestQueue2.Pop()
	if err == nil {
		t.Error("Invalid implementation of Pop")
	}
}

func TestQueue_Len(t *testing.T) {
	myTestQueue := NewQueue()
	_, err := myTestQueue.Add(1)

	_, err = myTestQueue.Add(2)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(3)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(4)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(5)
	if err != nil {
		t.Error(err)
	}

	_, err = myTestQueue.Add(6)
	if err != nil {
		t.Error(err)
	}

	l := myTestQueue.Len()
	if l != 6 {
		t.Error("Length not equal")
	}
}
