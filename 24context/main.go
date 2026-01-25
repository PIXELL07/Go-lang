package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // cancles work
	defer cancel()

	go task(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}
