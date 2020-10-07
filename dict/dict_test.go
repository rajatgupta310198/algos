package dict

import (
	"strconv"
	"sync"
	"testing"
)

func TestDict_Add(t *testing.T) {
	myDict := NewDict()
	err := myDict.Add("rajat", "Rajat")

	if err != nil {
		t.Error(err)
	}
	err = myDict.Add("rajat", 2)
	if err == nil {
		t.Error("Duplicated key entered")
	}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			err := myDict.Add(i, strconv.Itoa(i))
			if err != nil {
				t.Error(err)
			}
			defer wg.Done()
		}(i)

	}
	wg.Wait()

}

func TestDict_Get(t *testing.T) {
	myDict := NewDict()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			err := myDict.Add(i, strconv.Itoa(i))
			if err != nil {
				t.Error(err)
			}
			defer wg.Done()
		}(i)

	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			item, err := myDict.Get(i)
			if err != nil {
				t.Error(err)
			}
			if item != strconv.Itoa(i) {
				t.Error("invalid data")
			}
			defer wg.Done()
		}(i)

	}
}

func TestDict_GetAndRemove(t *testing.T) {
	myDict := NewDict()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			err := myDict.Add(i, strconv.Itoa(i))
			if err != nil {
				t.Error(err)
			}
			defer wg.Done()
		}(i)

	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			item, err := myDict.GetAndRemove(i)
			if err != nil {
				t.Error(err)
			}
			if item != strconv.Itoa(i) {
				t.Error("invalid data corresponding to key")
			}
			defer wg.Done()
		}(i)

	}
	wg.Wait()

	if len(myDict.items) != 0 {
		t.Error("getandremove didn't poped elements")
	}
}

func TestDict_Iter(t *testing.T) {
	myDict := NewDict()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			err := myDict.Add(i, i)
			if err != nil {
				t.Error(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	for item := range myDict.Iter() {
		if item != myDict.items[item] {
			t.Error("Mismatch")
		}
	}

}
