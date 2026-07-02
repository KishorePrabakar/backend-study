package main

import (
	"fmt"
	// "time"
)

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	// testing basic concurrency with with 'go' keyword
	/*go greet("Nitish")
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Hello World")*/

	//creating channels
	ch := make(chan int, 3) // buffered channel of size - 3
	ch <- 42
	ch <- 420
	ch <- 421
	val := <- ch
	val2 := <- ch
	// val3 := <- ch

	fmt.Println(val, val2)
	fmt.Println(len(ch), cap(ch)) // len - 1, as last value is queued in channel, cap - 3, as channel is buffered with size 3

	//creating unbuffered channel
	uch := make(chan int) // unbuffered channel
	// go func() {
	// 	uch <- 100 // sending value to channel
	// }()

	// time.Sleep(1 * time.Millisecond)
	
	// go func(){
	// 	uch <- 200 // sending value to channel
	// }()
	uch <- 100 // sending value to channel
	uval := <- uch
	fmt.Println(uval)
	fmt.Println(len(uch), cap(uch)) // len - 0, as all values are dequeued from channel, cap - 0, as channel is unbuffered
}
