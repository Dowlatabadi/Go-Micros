package main

import (
	"fmt" 
	"time" 
	"sync"
)

func produce(inc chan int) {
	for i := 0; i < 10; i++ {
		inc <- i
	}
	close(inc)
}

func consume(inc chan int,fin chan bool, consumer_number int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data:= range inc{
		time.Sleep(1 * time.Second)
		//		time.Sleep(2 * time.Second)
		fmt.Printf("consumer %v received %v \n",consumer_number,data)
	}	
}

func main() {
	wg := &sync.WaitGroup{}
	var fin = make(chan bool)
	var stream = make(chan int)
	go produce(stream)
	for i:=0 ; i<3;i++{
		wg.Add(1)
		go consume(stream,fin,i,wg)
	}
	wg.Wait()
}

