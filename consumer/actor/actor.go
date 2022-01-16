package actor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// -- Consumer below here!
type Consumer struct {
	IngestChan chan int
	JobsChan   chan int
}

// GetDataFromKafkaTopic is to emulate that get data from kafka
func (c Consumer) GetDataFromKafkaTopic() {
	myData := 0
	for {
		c.IngestChan <- myData
		log.Printf("Emulated data from kafka %d", myData)
		myData++
	}
}

// workerFunc starts a single worker function that will range on the jobsChan until that channel closes.
func (c Consumer) WorkerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	for myInteger := range c.JobsChan {

		fmt.Printf("Worker %d started job %d\n", index, myInteger)
		payload := `{"Mydata: "}` + strconv.FormatInt(int64(myInteger), 10)
		url := "DB service URL"
		//todo: error handle
		req, _ := http.NewRequest("POST", url, bytes.NewReader([]byte(payload)))
		client := &http.Client{}
		//todo: error handle and responce handler
		response, _ := client.Do(req)
		fmt.Printf("%v", response)
		time.Sleep(time.Millisecond * time.Duration(50))
		log.Printf("Worker %d finished processing job %d\n", index, myInteger)
	}
	fmt.Printf("Worker %d interrupted\n", index)
}

// startConsumer acts as the proxy between the ingestChan and jobsChan, with a select to support graceful shutdown.
func (c Consumer) StartConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.IngestChan:
			c.JobsChan <- job // get the integer from kafak and pass to worker
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.JobsChan)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}
