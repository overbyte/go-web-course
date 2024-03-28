package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error creating tcp server", err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	s := bufio.NewScanner(conn)
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}

	io.WriteString(conn, "I see you connected\n")
}
