package wc

import (
	"context"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	count, chans := 5, []chan interface{}{}
	for i := 0; i < count; i++ {
		chans = append(chans, make(chan interface{}))
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			close(chans[i])
		}(i)
	}

	done := Wait(context.Background(), chans...)
	if _, ok := <-done; ok {
		t.Errorf("expected channel to be closed")
	}
}

func TestWait_WithContext(t *testing.T) {
	count, chans := 5, []chan interface{}{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	for i := 0; i < count; i++ {
		chans = append(chans, make(chan interface{}))
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			time.Sleep(2 * time.Second)
			close(chans[i])
		}(i)
	}

	done := Wait(ctx, chans...)
	if _, ok := <-done; ok {
		t.Errorf("expected channel to be closed")
	}
}
