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
	m := ""
	u := ""
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}

		// we're in the first line of the header
		if i == 0 {
			f := strings.Fields(ln)	
			m = f[0]
			u = f[1]
			fmt.Println("Method:", m)
			fmt.Println("URI:", u)
		}

		i++
	}

	b := `I see you connected
method: %s
uri: %s`
	body := fmt.Sprintf(b, m, u)
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
