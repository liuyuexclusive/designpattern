package singleton

import (
	"sync"
	"testing"
)

func Test_SingleTon1(t *testing.T) {

	var wg sync.WaitGroup
	ch := make(chan *SingletonObj, 100)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				ch <- Single1()
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var pre *SingletonObj
	for item := range ch {
		if pre != nil {
			if pre != item {
				t.Errorf("pre: %p,item: %p\n", pre, item)
			}
		}
		pre = item
	}

}

func Test_SingleTon2(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan *SingletonObj, 100)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			ch <- Single2()
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var pre *SingletonObj
	for item := range ch {
		if pre != nil {
			if pre != item {
				t.Errorf("pre: %p,item: %p\n", pre, item)
			}
		}
		pre = item
	}
}
