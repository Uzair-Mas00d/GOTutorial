package main

import (
	"fmt"
	"sync"
	"time"
)

func processTransaction(transctionNumber int, done chan<- bool) {
	fmt.Println("Processing transcation", transctionNumber)
	time.Sleep(2 * time.Second)
	fmt.Println("Processed Transcation", transctionNumber)
	done <- true
}

func processTransactionData(transctionNumber int, done chan<- int) {
	fmt.Println("Processing transcation", transctionNumber)
	time.Sleep(2 * time.Second)
	fmt.Println("Processed Transcation", transctionNumber)
	done <- transctionNumber
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(transctionNumber int) {
			defer wg.Done()

			fmt.Println("Processing transcation", transctionNumber)
			time.Sleep(2 * time.Second)
			fmt.Println("Processed Transcation", transctionNumber)
		}(i)
	}

	wg.Wait()
	fmt.Println("All transcation processed")

	complete := make(chan bool)
	for i := 1; i <= 5; i++ {
		go processTransaction(i, complete)
	}

	for i := 1; i <= 5; i++ {
		<-complete
	}

	fmt.Println("All transcation processed")

	totalTranscation := 5
	proceed := make(chan int, totalTranscation)
	for i := 1; i <= totalTranscation; i++ {
		go processTransactionData(i, proceed)
	}

	for transcationNumber := range proceed {
		fmt.Printf("Recived completion signal for transcation %d\n", transcationNumber)
		if transcationNumber == totalTranscation {
			close(proceed)
		}
	}

	fmt.Println("All transcation processed")
}
