package usegdb

import (
	"fmt"
	"time"
)

//go build -gcflags "-N -l" -ldflags "-s" gdbfile.go
//gdb gdbfile

func couting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}

func Use_counting() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go couting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
