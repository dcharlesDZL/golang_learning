package main

import (
	"fmt"
	"time"
)

type Service struct {
	RecChannel chan int
	ReqChannel chan int
}

func (s *Service) Create() {
	for i := 0; i < 20; i++ {
		s.ReqChannel <- i
		time.Sleep(time.Second)
	}
}

func (s *Service) Receive() {
	for {
		result := <-s.ReqChannel
		fmt.Printf("===============received req : %d\n", result)
		go s.handle(result)
	}
}

func (s *Service) handle(i int) {
	// handle data
	fmt.Printf("i is :%d\n", i)
	i *= i
	time.Sleep(5 * time.Second)
	s.RecChannel <- i
	//fmt.Println(i)
}
func (s *Service) HandleReceive() {
	for {
		i := <-s.RecChannel
		fmt.Printf("===handleReceive %d\n", i)
	}

}
func main() {
	s := &Service{
		RecChannel: make(chan int, 5),
		ReqChannel: make(chan int, 10),
	}
	go s.Create()
	go s.Receive()
	go s.HandleReceive()
	time.Sleep(2000 * time.Second)
}
