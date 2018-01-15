package main

import (
	"net"
	log "github.com/nohupped/glog"
	"io"
)

func main() {
	li, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for ; ; {
		conn, err := li.Accept()
		if err != nil {
			log.Errorln(err)
		}
			io.WriteString(conn, "\nHello from tcp server\n")
			conn.Close()
	}
}
