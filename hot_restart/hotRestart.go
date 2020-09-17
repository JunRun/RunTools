package hot_restart

import (
	"fmt"
	"net"
	"os"
	"sync"
	"testing"
)

const (
	envRestart  = "RESTART"
	envListenFD = "LISTENFD"
)

func TestRestart(t *testing.T) {
	v := os.Getenv(envRestart)
	if v != "1" {
		ln, err := net.Listen("tcp", "localhost:8090")
		if err != nil {
			panic(err)
		}
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				ln.Accept()
			}
		}()
		tcpLn := ln.(*net.TCPListener)
		f, err := tcpLn.File()
		if err != nil {
			panic(err)
		}
		os.Setenv(envRestart, "1")
		os.Setenv(envListenFD, fmt.Sprintf("%d", f.Fd()))

	}

}
