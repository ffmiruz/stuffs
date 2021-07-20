package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var bufchan = make(chan struct{}, 8)
	var wg sync.WaitGroup

	for i := 1; i <= 25; i++ {
		wg.Add(1)
		bufchan <- struct{}{}
		go func(port int) {
			d := net.Dialer{Timeout: 1 * time.Second}
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := d.Dial("tcp", address)
			if err != nil {
				// port is closed or filtered.
			} else {
				fmt.Printf("Open %d \n", port)
				conn.Close()
			}
			<-bufchan
			wg.Done()
		}(i)
	}
	wg.Wait()
}
