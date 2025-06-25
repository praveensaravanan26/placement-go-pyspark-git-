package main

import (
	"fmt"
	"sync"
)

func main() {
	const t = 100000000000000000
	done := make(chan struct{})
	num := 0
	m := sync.Mutex{}
	go func() {
		for i := 1; i < t; i++ {
			m.Lock()
			num++
			m.Unlock()
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 1; i < t; i++ {
			m.Lock()
			num--
			m.Unlock()
		}
		done <- struct{}{}
	}()
	<-done
	<-done

	fmt.Println(num)
}
