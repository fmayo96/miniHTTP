package minihttp

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
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

type HttpStatusCode int

const (
	Ok                  HttpStatusCode = 200
	Created             HttpStatusCode = 201
	NoContent           HttpStatusCode = 204
	BadRequest          HttpStatusCode = 400
	Forbidden           HttpStatusCode = 403
	NotFound            HttpStatusCode = 404
	InternalServerError HttpStatusCode = 500
)

type Route struct {
	method  HttpMethod
	path    string
	handler HandlerFunc
}

type Request struct {
	Method  HttpMethod
	Path    string
	Headers []string
	Body    []byte
}

type Response struct {
	conn    net.Conn
	Status  HttpStatusCode
	Headers []string
	Body    string
	buf     []byte
}

func (res *Response) SetHeader(key, value string) {
	res.Headers = append(res.Headers, key+": "+value+"\r\n")
}

func (res *Response) SetStatus(status HttpStatusCode) {
	res.Status = status
}

func (res *Response) Send(message ...string) {
	res.Write([]byte("HTTP/1.1 "))
	res.Write([]byte(strconv.Itoa(int(res.Status)) + "\r\n"))
	for _, m := range message {
		res.Write([]byte("\r\n" + m))
	}
	_, err := res.conn.Write(res.buf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (res *Response) Write(p []byte) (int, error) {
	res.buf = append(res.buf, p...)
	return len(p), nil
}

func (res *Response) Json(v interface{}) {
	res.Write([]byte("HTTP/1.1 "))
	res.Write([]byte(strconv.Itoa(int(res.Status)) + "\r\n"))
	res.SetHeader("Content-Type", "application/json")
	for _, header := range res.Headers {
		res.Write([]byte(header))
	}
	res.Write([]byte("\r\n"))
	json.NewEncoder(res).Encode(v)
	res.conn.Write(res.buf)
}
