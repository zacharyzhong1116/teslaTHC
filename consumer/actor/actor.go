package consumer

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// -- Consumer below here!
type Consumer struct {
	IngestChan chan int
	JobsChan   chan int
}

// callbackFunc is invoked each time the external lib passes an event to us.
func (c Consumer) CallbackFunc(event int) {
	c.IngestChan <- event
}

// workerFunc starts a single worker function that will range on the jobsChan until that channel closes.
func (c Consumer) WorkerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", index)
	for eventIndex := range c.JobsChan {
		// simulate work  taking between 1-3 seconds
		fmt.Printf("Worker %d started job %d\n", index, eventIndex)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Worker %d finished processing job %d\n", index, eventIndex)
	}
	fmt.Printf("Worker %d interrupted\n", index)
}

// startConsumer acts as the proxy between the ingestChan and jobsChan, with a select to support graceful shutdown.
func (c Consumer) StartConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.IngestChan:
			c.JobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.JobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}
