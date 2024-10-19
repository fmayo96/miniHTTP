package main

import (
	"fmt"
	"net"
	"strings"
)

type Server struct {
	Port    int
	Handler func(conn net.Conn)
}

func (s *Server) Start() {
	addr := net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: s.Port,
	}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Server listening on http://localhost:%d\n", s.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConn(conn)
	}
}

type Request struct {
	Method string
	Route  string
}

type Response struct {
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		bytesRecv, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}
		if bytesRecv == 0 {
			fmt.Println("Client disconnected")
			break
		}
		parseReq := strings.Split(string(buf), " ")[:3]
		r := Request{Method: parseReq[0], Route: parseReq[1]}

		if (r.Method == "GET") && (r.Route == "/") {
			conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n<!DOCTYPE html>\r\n<html lang=\"en\">\r\n<head>\r\n<meta charset=\"UTF-8\">\r\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\r\n<title>Hello World</title>\r\n</head>\r\n<body>\r\n<h1>Hello World</h1>\r\n</body>\r\n</html>"))
			return
		} else {
			conn.Write([]byte("HTTP/1.1 404\r\n"))
			return
		}
	}
}
