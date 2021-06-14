package singleton

import "sync"

type SingletonObj struct {
	Data interface{}
}

var single1 *SingletonObj
var single2 *SingletonObj
var mutex sync.Mutex
var once sync.Once

func Single1() *SingletonObj {
	if single1 == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if single1 == nil {
			single1 = &SingletonObj{}
		}
	}
	return single1
}

func Single2() *SingletonObj {
	if single2 == nil {
		once.Do(func() {
			single2 = &SingletonObj{}
		})
	}
	return single2
}
