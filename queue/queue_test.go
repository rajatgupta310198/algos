package queue

import (
	"testing"
)

func TestQueue_Add(t *testing.T) {
	myTestQueue := NewQueue()
	myTestQueue.Add(1)

	 myTestQueue.Add(3)
	myTestQueue.Add(2)

	myTestQueue.Add(3)
	item, err := myTestQueue.Pop()
	if err != nil {
		t.Error(err)
	}
	if item != 1 {
		t.Error("Wrong implementation")
	}

}

func TestQueue_Pop(t *testing.T) {
	myTestQueue := NewQueue()
	myTestQueue.Add(1)

	myTestQueue.Add(3)
	myTestQueue.Add(2)
	myTestQueue.Add(3)

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
	myTestQueue.Add(1)

	myTestQueue.Add(2)
	myTestQueue.Add(3)
	myTestQueue.Add(4)

	myTestQueue.Add(5)


	 myTestQueue.Add(6)


	l := myTestQueue.Len()
	if l != 6 {
		t.Error("Length not equal")
	}
}
