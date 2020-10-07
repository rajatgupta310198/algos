package dict

import (
	"errors"
	"sync"
)

// Dict struct for golang
type Dict struct {
	items map[interface{}]interface{}
	mutex sync.Mutex
}

// Add will add items to dict
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

// Get will obtain value of given key
func (d *Dict) Get(key interface{}) (interface{}, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	item, ok := d.items[key]
	if ok != true {
		return nil, errors.New("key not present")
	}
	return item, nil
}

// GetAndRemove will get the item corresponding to key and removes it from dict
func (d *Dict) GetAndRemove(key interface{}) (interface{}, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	item, ok := d.items[key]
	if ok != true {
		return nil, errors.New("key not present")
	}
	delete(d.items, key)
	return item, nil
}

// Remove will remove item from dict
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

// Iter will create iterator over values in dict
func (d *Dict) Iter() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for _, i := range d.items {
			ch <- i
		}
	}()
	return ch
}

// IterKeys will create iterator over keys in dict
func (d *Dict) IterKeys() chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for key, _ := range d.items {
			ch <- key
		}
	}()
	return ch
}

// Update specific value of given key
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

// NewDict will create a dict pointer
func NewDict() *Dict {
	return &Dict{
		items: make(map[interface{}]interface{}),
	}
}
