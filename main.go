package main

import (
	"net/http"
	"sync"
	"log"
	"time"
)


func main() {
	url := "http://localhost:8080"

	maxConnection := make(chan bool, 10)
	wg := &sync.WaitGroup{}

	count := 0
	start := time.Now()

	for maxRequest := 0; maxRequest < 10000; maxRequest ++ {
		wg.Add(1)
		maxConnection <- true
		go func() {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				return
			}

			defer resp.Body.Close()

			count ++

			<-maxConnection
		}()

	}

	wg.Wait()
	end := time.Now()
	log.Println("%d request succeded", count)
	log.Println("%f sec took ", (end.Sub(start)).Seconds())

}