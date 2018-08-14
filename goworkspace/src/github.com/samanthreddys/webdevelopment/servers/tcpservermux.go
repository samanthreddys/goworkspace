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

	//respond(conn)
}

func request(conn net.Conn) {
	//defer conn.Close()
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++

	}
}

func mux(conn net.Conn, ln string) {
	//defer conn.Close()
	fmt.Println("Conn:", conn)
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	if m == "GET" && u == "/" {
		index(conn)

	}
	if m == "GET" && u == "/about" {
		about(conn)

	}
	if m == "GET" && u == "/contact" {
		contact(conn)

	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)

	}
	if m == "GET" && u == "/apply" {
		apply(conn)

	}

}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

	<strong>Index</strong><br>
	<a href="/" >Index</a><br>
	<a href="/about" >About</a><br>
	<a href="/apply" >Apply</a><br>
	<a href="/contact" >Contact</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)

}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

	<strong>ABOUT</strong><br>
	<a href="/" >Index</a><br>
	<a href="/about" >About</a><br>
	<a href="/apply" >Apply</a><br>
	<a href="/contact" >Contact</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)

}
func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

	<strong>APPLY</strong><br>
	<a href="/" >Index</a><br>
	<a href="/about" >About</a><br>
	<a href="/apply" >Apply</a><br>
	<a href="/contact" >Contact</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)

}
func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

	<strong>APPLY PROCESS</strong><br>
	<a href="/" >Index</a><br>
	<a href="/about" >About</a><br>
	<a href="/apply" >Apply</a><br>
	<a href="/contact" >Contact</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="applyProcess">
	</form>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)

}
func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang ="en"><head><meta charset="UTF-8"><title>GO World!</title></head><body>

	<strong>CONTACT</strong><br>
	<a href="/" >Index</a><br>
	<a href="/about" >About</a><br>
	<a href="/apply" >Apply</a><br>
	<a href="/contact" >Contact</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Connection-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Conent-Type:text/html \r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)

}
