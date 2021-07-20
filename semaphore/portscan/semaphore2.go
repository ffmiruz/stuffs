package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var sem = semaphore.NewWeighted(int64(8))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)

	for i := 1; i <= 25; i++ {
		sem.Acquire(ctx, 1)
		go func(port int) error {
			g.Go(func() error {
				defer sem.Release(1)
				d := net.Dialer{Timeout: 1 * time.Second}
				address := fmt.Sprintf("scanme.nmap.org:%d", port)
				conn, err := d.DialContext(ctx, "tcp", address)
				if err != nil {
					return err
				} else {
					fmt.Printf("Open %d \n", port)
					conn.Close()
				}
				return nil
			})
			return nil
		}(i)
	}
	if err := g.Wait(); err == nil || err == context.Canceled {
		fmt.Println("finished clean")
	} else {
		fmt.Printf("received error: %v", err)
	}
}
