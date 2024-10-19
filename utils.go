package main

import "errors"

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
