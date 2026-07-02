package main

import (
	"fmt"
)

func worker(ping chan string, pong chan string) {
	val := <- ping
	fmt.Println("Ping Message Recieved: ", val)

	pong <- "This is a pong message"
}

func main(){
	ping := make(chan string)
	pong := make(chan string)

	go worker(ping, pong)

	ping <- "PING 1 Message"
	p2 := <- pong
	fmt.Println("Pong Message Recieved: ", p2)
}