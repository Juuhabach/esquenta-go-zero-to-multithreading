package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	requestID := make(chan int)
	concurrency := 3
	for i := 0; i < concurrency; i++ {
		go worker(requestID, i)
	}
	for i := 0; i < 20; i++ {
		requestID <- i
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {
		res, err := http.Get("http://localhost:8585/product")
		if err != nil {
			log.Fatal("Error")
		}
		defer res.Body.Close()
		content, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Worker %d. Request ID: %d. Content: %s\n", worker, r, string(content))
		i := rand.Intn(5)
		time.Sleep(time.Second * time.Duration(i))
	}
}
