package dict

import "sync"

type Dict struct {
	items map[interface{}]interface{}
	mutex sync.Mutex
}



