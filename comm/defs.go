package comm

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	Sig  = make(chan os.Signal, 1)
	Quit = make(chan struct{})
)

func Init() {
	signal.Notify(Sig, syscall.SIGINT, syscall.SIGTERM)
}
