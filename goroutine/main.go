package main

import (
	"log"
	"time"
	"sync"
)

func main() {
	log.Print("started.")

	funcs := []func() {
		func() {
			log.Print("sleep1 started.")
			time.Sleep(1 * time.Second)
			log.Print("sleep1 finished.")
		},
		func() {
			log.Print("sleep2 started.")
			time.Sleep(2 * time.Second)
			log.Print("sleep2 finished.")
		},
		func() {
			log.Print("sleep3 started.")
			time.Sleep(3 * time.Second)
			log.Print("sleep3 finished.")
		},
	}

	var waitGroup sync.WaitGroup
	for _, sleep := range funcs {
		waitGroup.Add(1)
		go func(run func()) {
			defer waitGroup.Done()
			run()
		}(sleep)

	}
	waitGroup.Wait()

	log.Print("all finished.")
}