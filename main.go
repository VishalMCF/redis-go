package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("COuld not start the redis server")
		return
	}
	conn, err := l.Accept()
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		input := make([]byte, 1024)
		fmt.Println(conn.Read(input))
	}

}
