package main

import (
	"log"
)

func main() {
	c := make(chan int)
	go func(chan int) {
		log.Println("Start")
		c <- 1
		c <- 2
		log.Println("Stop")
	}(c)
}
