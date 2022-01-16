package collector

import (
	"fmt"
	"io"
	"log"
	"os"
)

// -- Producer simulates an external library that invokes the
// registered callback when it has new data for us once per 100ms.
type Producer struct {
	FilePath string
	Data     chan int
}

func readFile(filePath string, dataC chan int) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var data int
	for {

		_, err := fmt.Fscanf(fd, "%d\n", &data)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				close(dataC)
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		dataC <- data
	}
	return
}
func (p Producer) Start() {
	// here we suppose to check if the kafka service is abailabe
	if _, err := os.Stat(p.FilePath); err != nil {
		log.Fatal("File not exist")
	}
	go readFile(p.FilePath, p.Data)
	for data := range p.Data {
		//send data to Kakfa
		log.Printf("Get data %d", data)
	}
	log.Println("Finish processing data")
}
