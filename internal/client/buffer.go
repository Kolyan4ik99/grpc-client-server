package client

import (
	"log"
	"sync"
	"time"
)

type buffer struct {
	len       int64
	cap       int64
	threshold time.Duration
	ticker    *time.Ticker

	buf []int64
	mu  sync.Mutex
}

// newBuffer return buffer instance and run in goroutine internal scheduler
func newBuffer(cap int64, threshold time.Duration) *buffer {
	bufTmp := &buffer{
		len:       0,
		cap:       cap,
		threshold: threshold,
		buf:       make([]int64, cap),
	}

	go bufTmp.scheduler()
	return bufTmp
}

func (b *buffer) scheduler() {
	b.ticker = time.NewTicker(b.threshold)
	for {
		<-b.ticker.C
		b.mu.Lock()
		b.flush()
		b.mu.Unlock()
	}
}

func (b *buffer) put(index int64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.len == b.cap {
		b.flush()
	}
	b.buf[b.len] = index
	b.len++
}

func (b *buffer) flush() {
	for i := int64(0); i < b.len; i++ {
		log.Println(b.buf[i])
		b.buf[i] = 0
	}
	b.len = 0
	b.ticker.Reset(b.threshold)
}
