package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
	i := 0
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}

		// we're in the first line of the header
		if i == 0 {
			f := strings.Fields(ln)	
			method := f[0]
			uri := f[1]
			fmt.Println("Method:", method)
			fmt.Println("URI:", uri)
		}

		i++
	}

	body := "I see you connected"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
