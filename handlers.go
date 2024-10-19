package main

import (
	"log"
	"net"
)

type HandlerFunc func(req Request, res *Response)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

type Route struct {
	method  HttpMethod
	path    string
	handler HandlerFunc
}

type Request struct {
	Method HttpMethod
	Path   string
}

type Response struct {
	conn net.Conn
}

func (res *Response) Write(message string) {
	log.Println("Writing started")
	responseString := "HTTP/1.1 200 Ok\r\nContent-Type: text/html\r\n\r\n" + message
	bytesWritten, err := res.conn.Write([]byte(responseString))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(bytesWritten)
}
