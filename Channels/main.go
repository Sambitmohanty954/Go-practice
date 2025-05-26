package main

import (
	"fmt"
	"time"
)

//Channels in Go are pipes that let goroutines communicate with each other and share data safely.
// Think of them as message queues or shared mailboxes between goroutines.

//  Why Use Channels?
//To send data from one goroutine to another.
//To avoid using shared memory or locks.
//To synchronize work between goroutines.

// sending the data to the channels
//
//	func processNum(numChan chan int) {
//		for num := range numChan {
//			fmt.Println("Processing number", num)
//			time.Sleep(time.Second * 2)
//		}
//	}

// receiving the data
func sum(result chan int, num1 int, num2 int) {
	result2 := num1 + num2
	result <- result2
}

// emailChan <-chan string :- it means we can receive only data, we can't send it,
// done chan <- bool :- this means we can only send the data, can't receive
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Email sent to ", email)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "sambitt"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num1 := <-chan1:
			fmt.Println(num1)
		case num2 := <-chan2:
			fmt.Println(num2)
		}
	}

	//Buffered channel

	//emailChan := make(chan string, 100)
	//done := make(chan bool)
	//
	//go emailSender(emailChan, done)
	//
	//// using Sprintf we can format the string
	//for i := 0; i < 10; i++ {
	//	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	//}
	//
	//fmt.Println("Done sending emails")
	//emailChan <- "1@example.com"
	//emailChan <- "2@example.com"

	//fmt.Println(<-emailChan)
	//fmt.Println(<-emailChan)

	// this is important to close the buffer channel
	//close(emailChan)
	//<-done

	//result := make(chan int)
	//go sum(result, 4, 5)
	//
	//res := <-result
	//fmt.Println(res)
	//ch := make(chan string) // create a channel of type int
	//go func() {
	//	ch <- "Hello from goroutine!" // send data into channel
	//}()
	//msg := <-ch      // receive data from channel
	//fmt.Println(msg) // prints: "Hello from goroutine!"

	//numberChan := make(chan int)
	//go processNum(numberChan)
	//
	//for {
	//	numberChan <- rand.Intn(100)
	//}

	//numberChan <- 2
	//time.Sleep(1 * time.Second)
	//Unbuffered channel
	//messageChan := make(chan string)
	//messageChan <- "Hello World" // deadlock!
	//msg := <-messageChan
	//fmt.Println(msg)
}
