// Defines the Server and all it's functionalities
// including the methods for creating new routes
// and starting the server
package main

import (
	"log"
	"net"
	"strings"
)

type Server struct {
	Port   int
	routes []Route
}

func (s *Server) Start() {
	addr := net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: s.Port,
	}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("Server listening on http://localhost:%d\n/", s.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		go s.handleConn(conn)
	}
}

func (s *Server) GET(path string, handler HandlerFunc) {
	newRoute := Route{GET, path, handler}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) POST(path string, handler HandlerFunc) {
	newRoute := Route{POST, path, handler}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) PUT(path string, handler HandlerFunc) {
	newRoute := Route{PUT, path, handler}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) PATCH(path string, handler HandlerFunc) {
	newRoute := Route{PATCH, path, handler}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) DELETE(path string, handler HandlerFunc) {
	newRoute := Route{DELETE, path, handler}
	s.routes = append(s.routes, newRoute)
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	bytesRecv, err := conn.Read(buf)
	if err != nil {
		log.Println(err.Error())
	}
	if bytesRecv == 0 {
		log.Println("Client disconnected")
		return
	}
	parseReq := strings.Split(string(buf), " ")[:2]
	r := Request{Method: HttpMethod(parseReq[0]), Path: parseReq[1]}
	log.Println("request method", r.Method)
	log.Println("request path", r.Path)
	idx, err := findRoute(r, s.routes)
	if err != nil {
		log.Println(err.Error())
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		return
	}
	res := &Response{conn: conn}
	s.routes[idx].handler(r, res)
}
