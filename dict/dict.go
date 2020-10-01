package dict

import (
	"errors"
	"sync"
)

type Dict struct {
	items map[interface{}]interface{}
	mutex sync.Mutex
}

func (d *Dict) Add(key, value interface{}) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.items[key]
	if ok {
		return errors.New("key already exists")
	}
	d.items[key] = value
	return nil
}

func (d *Dict) Get(key interface{}) (interface{}, error)  {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	item, ok := d.items[key]
	if ok != true {
		return nil, errors.New("key not present")
	}
	return item, nil
}

func (d *Dict) GetAndRemove(key interface{}) (interface{}, error)  {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	item, ok := d.items[key]
	if ok != true {
		return nil, errors.New("key not present")
	}
	delete(d.items, key)
	return item, nil
}

func (d *Dict) Remove(key interface{}) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.items[key]
	if ok != true {
		return errors.New("key not present")
	}
	delete(d.items, key)
	return nil
}

func (d *Dict) Iter() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for _, i:= range d.items {
			ch <- i
		}
	}()
	return ch
}

func (d *Dict) Update(key, value interface{}) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, ok := d.items[key]
	if !ok {
		return errors.New("key not present")
	}
	d.items[key] = value
	return nil
}

func NewDict() *Dict {
	return &Dict{
		items: make(map[interface{}]interface{}),
	}
}
