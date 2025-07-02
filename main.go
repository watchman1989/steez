package main

import (
	"fmt"
	"time"

	"github.com/watchman1989/steez/comm"
	"github.com/watchman1989/steez/srv"
	"github.com/watchman1989/steez/jobs"
)

func main() {
	fmt.Println("hello, steez!")

	comm.Init()

	jm := jobs.InitJobs()
	defer jm.Stop()

	go srv.SrvStart()

	<-comm.Sig

	close(comm.Quit)

	fmt.Printf("steez exit\n")
	time.Sleep(time.Second)

}
