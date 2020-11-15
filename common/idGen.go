package common

import "sync"

var mutex sync.Mutex
var id uint64 = 1000 // 1000以下为保留id

const UINT64_MAX uint64 = ^uint64(0)

func GetId() uint64 {
	mutex.Lock()
	defer mutex.Unlock()
	if id == UINT64_MAX {
		id = 1000
	}
	id += 1
	return id
}
