package wc

import (
	"context"
	"sync"
)

func Wait(ctx context.Context, chans ...chan interface{}) chan struct{} {
	wg := &sync.WaitGroup{}
	done := make(chan struct{})
	for i := range chans {
		wg.Add(1)
		done := ctx.Done()
		go func(c chan interface{}) {
			defer wg.Done()
			select {
			case <-done:
			case <-c:
			}
		}(chans[i])
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}
