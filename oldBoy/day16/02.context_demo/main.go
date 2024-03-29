package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("hello sven")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	cancel()
	wg.Wait()

}
