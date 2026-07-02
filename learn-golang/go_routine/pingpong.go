package main

import (
	"fmt"
	"time"
)

//Ping Pong channels - Demonstrating the simulataneous sending and waiting of messages, between two unbufferd channels

func worker(ping chan string, pong chan string) {
	ping <- "Message 1"
	val := <- pong 
	fmt.Println(val)
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go worker(ping, pong)

	msg1 := <- ping
	pong <- msg1
	time.Sleep(1 * time.Millisecond)
}