package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, errors chan<- error) {
	for j := range jobs {
		fmt.Printf("worker #%d: init-d job №%d\n", id, j)
		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		fmt.Printf("Worker #%d: fin-d job №%v\n", id, j)
		results <- j
		wg.Done()
	}
}

func main() {
	nJobs := 10
	nWorkers := 3
	jobs := make(chan int, nJobs)
	results := make(chan int, nWorkers)
	errors := make(chan error, nJobs)

	var wg sync.WaitGroup
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i, &wg, jobs, results, errors)
		}()
	}
	for j := 1; j <= nJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
