package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	p.views++
	p.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for range 100 {
		wg.Add(1)
		go myPost.incrementViews(&wg)
	}
	wg.Wait()
	fmt.Println("Post views:", myPost.views)
}
