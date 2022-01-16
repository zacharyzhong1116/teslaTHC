package main

import (
	"context"
	"sync"
	"teslaTHC/consumer/actor"
)

var consumer actor.Consumer

const workerPoolSize = 4

func main() {
	exitCahnnel := make(chan int)

	// Set up cancellation context and waitgroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	consumer := actor.Consumer{
		IngestChan: make(chan int, 1),
		JobsChan:   make(chan int, workerPoolSize),
	}
	// GetDataFromKafkaTopic is to emulate that get data from kafka
	go consumer.GetDataFromKafkaTopic()
	wg := &sync.WaitGroup{}
	// Start consumer with cancellation context passed

	go consumer.StartConsumer(ctx)
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.WorkerFunc(wg, i)
	}
	cancelFunc()
	<-exitCahnnel
}
