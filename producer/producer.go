package main

import (
	"teslaTHC/producer/collector"
)

var producer collector.Producer

const MYDATAPATH = "teslaTHC/producer/random_number.txt"

func main() {
	exitCahnnel := make(chan int)
	producer = collector.Producer{FilePath: MYDATAPATH, Data: make(chan int, 1)}
	producer.Start()
	<-exitCahnnel
}
