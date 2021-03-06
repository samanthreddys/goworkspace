package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	//defer conn.Close()
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			m1 := strings.Fields(ln)[1]
			fmt.Println("Methods:", m)
			fmt.Println("URI", m1)
		}
		if ln == "" {
			break
		}
		i++

	}
}

func respond(conn net.Conn) {
	//defer conn.Close()
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

 
   


	<p>My Name is Tom
	
	Go Nested Templates
	Tom
	
	</p>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
