package main

import (
	"fmt"
	"steez/comm"
	"steez/srv"
	"time"
)

func main() {
	fmt.Println("hello, steez!")

	comm.Init()

	go srv.SrvStart()

	<-comm.Sig

	close(comm.Quit)

	fmt.Printf("steez exit\n")
	time.Sleep(time.Second)

}
