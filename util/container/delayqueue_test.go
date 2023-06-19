package container

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func BenchmarkPushAndTake(b *testing.B) {
	q := NewDelayQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Push(i, time.Duration(i))
	}
	b.StopTimer()
	time.Sleep(time.Duration(b.N))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, ok := q.Take(context.Background())
		if !ok {
			b.Errorf("want %v, but %v", true, ok)
		}
	}
}

type timerEvent struct {
	id int
	cb func(id int)
}

func timerCallback(id int) {
	fmt.Println(id)
}

func TestDelayQueue(t *testing.T) {
	queue := NewDelayQueue[timerEvent]()

	go func() {
		i := 0
		for {
			queue.Push(timerEvent{id: i, cb: timerCallback}, time.Duration(i))
			fmt.Println(i)
			time.Sleep(time.Second)
			i++
		}
	}()

	go func() {
		e, ok := queue.Take(context.Background())
		if ok {
			e.cb(e.id)
		}
	}()
	
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
