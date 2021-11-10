package main

import (
	"context"
	"fmt"
	"time"
)

/*type Student struct {
	Name string
	Age int
	Address []int
}

type Student2 struct {
	Name string
	Age int
}*/
func main() {
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go monitor(nil)
	time.Sleep(time.Second * 10)
}

func monitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("I am done,err: %v", ctx.Err())
			return
		}
	}
}
