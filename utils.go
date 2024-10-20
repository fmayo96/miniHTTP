package minihttp

import (
	"errors"
	"strings"
)

func findRoute(req Request, routes []Route) (int, error) {
	for i, route := range routes {
		if req.Path == route.path {
			if req.Method == route.method {
				return i, nil
			}
		}
	}
	return -1, errors.New("not found")
}

func ParseRequest(buf []byte) Request {
	parseReq := strings.Split(string(buf), "\r\n\r\n")
	headers := parseReq[0]
	body := trimBytes([]byte(parseReq[1]))
	parseHeaders := strings.Split(headers, "\r\n")
	method := strings.Split(parseHeaders[0], " ")[0]
	path := strings.Split(parseHeaders[0], " ")[1]
	req := Request{Method: HttpMethod(method), Path: path, Headers: parseHeaders[1:], Body: body}
	return req
}

func trimBytes(buf []byte) []byte {
	lastIdx := len(buf) - 1
	for i := len(buf) - 1; i > 0; i-- {
		if buf[i] != 0 {
			lastIdx = i + 1
			break
		}
	}
	return buf[:lastIdx]
}
