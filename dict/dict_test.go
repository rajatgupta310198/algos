package dict

import (
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
			err := myDict.Add(i, "s")
			if err != nil {
				t.Error(err)
			}
			defer wg.Done()
		}(i)

	}
	wg.Wait()

}
