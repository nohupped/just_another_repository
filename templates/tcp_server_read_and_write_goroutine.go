package main

import (
	"net"
	"github.com/nohupped/glog"
	"bufio"
	"fmt"
	"time"
)

func main() {

	li, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for ; ; {
		conn, err := li.Accept()
		if err != nil {
			glog.Errorln(err)
		}
			go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	glog.SetStandardLogLevel(glog.DebugLevel)
	for scanner.Scan() {
		ln := scanner.Bytes()
		glog.Infoln(string(ln))
		fmt.Fprint(conn, "How are you?")
		conn.SetDeadline(time.Now().Add(time.Second*5))
	}
	glog.Println("Code reached here")
}
